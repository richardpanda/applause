package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

type Creator struct {
	ID   string
	Name string
}

type Post struct {
	Creator
	ID             string
	Title          string
	TotalClapCount int
	UniqueSlug     string
	URL            string
}

type ByTotalClapCount []Post

func (p ByTotalClapCount) Len() int           { return len(p) }
func (p ByTotalClapCount) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByTotalClapCount) Less(i, j int) bool { return p[i].TotalClapCount < p[j].TotalClapCount }

type StreamResponse struct {
	Payload struct {
		References struct {
			User map[string]struct {
				Name string `json:"name"`
			} `json:"user"`
			Post map[string]struct {
				CreatorID  string `json:"creatorId"`
				Title      string `json:"title"`
				UniqueSlug string `json:"uniqueSlug"`
				Virtuals   struct {
					TotalClapCount int `json:"totalClapCount"`
				} `json:"virtuals"`
			} `json:"post"`
		} `json:"references"`
	} `json:"payload"`
}

func sendGetRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.1 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func main() {
	resp, err := sendGetRequest("https://medium.com/_/api/topics/55f1c20aba7a/stream?limit=25")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var streamResponse StreamResponse
	b = b[16:]
	err = json.Unmarshal(b, &streamResponse)
	if err != nil {
		panic(err)
	}

	var posts []Post

	for postID, post := range streamResponse.Payload.References.Post {
		creatorName := strings.ToLower(strings.Replace(streamResponse.Payload.References.User[post.CreatorID].Name, " ", "", -1))
		p := Post{
			Creator: Creator{
				ID:   post.CreatorID,
				Name: creatorName,
			},
			ID:             postID,
			Title:          post.Title,
			TotalClapCount: post.Virtuals.TotalClapCount,
			UniqueSlug:     post.UniqueSlug,
			URL:            fmt.Sprintf("https://medium.com/@%s/%s", creatorName, post.UniqueSlug),
		}

		posts = append(posts, p)
	}

	sort.Sort(sort.Reverse((ByTotalClapCount(posts))))

	fmt.Println()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "Title\tTotal Clap Count\t URL")
	for _, p := range posts {
		fmt.Fprintf(w, "%s\t%d\t%s\n", p.Title, p.TotalClapCount, p.URL)
	}
	w.Flush()
}

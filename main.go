package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/richardpanda/applause/medium"
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

func printPosts(posts []Post) {
	fmt.Println()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "#\tTitle\tTotal Clap Count\t URL")
	for idx, p := range posts {
		fmt.Fprintf(w, "%d\t%s\t%d\t%s\n", idx+1, p.Title, p.TotalClapCount, p.URL)
	}
	w.Flush()
}

func main() {
	var (
		posts    []Post
		to       string
		baseURL  = "https://medium.com/_/api/topics/55f1c20aba7a/stream?limit=25"
		limit    = flag.Int("limit", 0, "max posts to display")
		numPages = flag.Int("pages", 1, "number of pages to process")
	)
	flag.Parse()

	for i := 0; i < *numPages; i++ {
		url := baseURL
		if to != "" {
			url = url + "&to=" + to
		}

		streamResponse, err := medium.FetchStreamResponse(url)
		if err != nil {
			panic(err)
		}

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

		if *limit != 0 && len(posts) > *limit {
			posts = posts[:*limit]
		}

		to = streamResponse.Payload.Paging.Next.To
	}

	printPosts(posts)
}

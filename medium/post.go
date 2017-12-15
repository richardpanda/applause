package medium

import (
	"fmt"
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

type Posts []Post

func (p Posts) Len() int {
	return len(p)
}

func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Posts) Less(i, j int) bool {
	return p[i].TotalClapCount < p[j].TotalClapCount
}

func (p Posts) SortByClapsDESC() {
	sort.Sort(sort.Reverse(p))
}

func (p *Posts) Append(nextPosts Posts) {
	*p = append(*p, nextPosts...)
}

func (p *Posts) Limit(n int) {
	if n == 0 || len(*p) <= n {
		return
	}

	var topPosts Posts
	for idx, post := range *p {
		if idx >= n {
			break
		}
		topPosts = append(topPosts, post)
	}
	*p = topPosts
}

func (p Posts) Print() {
	fmt.Println()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "#\tTitle\tTotal Clap Count\t URL")
	for idx, post := range p {
		fmt.Fprintf(w, "%d\t%s\t%d\t%s\n", idx+1, post.Title, post.TotalClapCount, post.URL)
	}
	w.Flush()
}

func PostsFromResponse(s *StreamResponse) Posts {
	var posts Posts
	for postID, post := range s.Payload.References.Post {
		creatorName := strings.ToLower(strings.Replace(s.Payload.References.User[post.CreatorID].Name, " ", "", -1))
		newPost := Post{
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
		posts = append(posts, newPost)
	}
	return posts
}

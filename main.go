package main

import (
	"flag"

	"github.com/richardpanda/applause/medium"
)

func main() {
	var (
		posts    medium.Posts
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

		posts.Append(medium.PostsFromResponse(streamResponse))
		to = streamResponse.Payload.Paging.Next.To
	}

	posts.SortByClapsDESC()
	posts.Limit(*limit)
	posts.Print()
}

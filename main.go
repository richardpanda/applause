package main

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/richardpanda/applause/medium"
)

func main() {
	var (
		posts     medium.Posts
		to        string
		limit     = flag.Int("limit", 0, "max posts to display")
		numPages  = flag.Int("pages", 1, "number of pages to process")
		topicName = flag.String("topic", "software-engineering", "topic to crawl")
	)
	flag.Parse()

	topicID, ok := medium.GetTopicID(*topicName)
	if !ok {
		panic(errors.New("unknown topic"))
	}

	baseURL := fmt.Sprintf("https://medium.com/_/api/topics/%s/stream?limit=25", topicID)
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

		time.Sleep(2 * time.Second)
	}

	posts.SortByClapsDESC()
	posts.Limit(*limit)
	posts.Print()
}

package medium

import (
	"testing"
	"time"
)

func TestSortByClapsDESC(t *testing.T) {
	posts := Posts{
		Post{
			Creator:        Creator{ID: "1", Name: "Test1"},
			ID:             "1",
			Title:          "Title 1",
			TotalClapCount: 1,
			UniqueSlug:     "unique-slug-1",
			URL:            "www.test.com",
		},
		Post{
			Creator:        Creator{ID: "2", Name: "Test2"},
			ID:             "2",
			Title:          "Title 2",
			TotalClapCount: 3,
			UniqueSlug:     "unique-slug-2",
			URL:            "www.test.com",
		},
		Post{
			Creator:        Creator{ID: "3", Name: "Test1"},
			ID:             "3",
			Title:          "Title 3",
			TotalClapCount: 2,
			UniqueSlug:     "unique-slug-3",
			URL:            "www.test.com",
		},
	}
	posts.SortByClapsDESC()

	want := []string{"2", "3", "1"}
	for idx, post := range posts {
		if post.ID != want[idx] {
			t.Fatalf("\nExpected: %s\nActual: %s\n", want, post.ID)
		}
	}
}

func TestLimit(t *testing.T) {
	p := make(Posts, 15)
	p.Limit(0)
	if len(p) != 15 {
		t.Fatalf("\nExpected: %d\nActual: %d\n", 15, len(p))
	}

	p.Limit(10)
	if len(p) != 10 {
		t.Fatalf("\nExpected: %d\nActual: %d\n", 10, len(p))
	}
}

func TestPostsFromResponse(t *testing.T) {
	url := "https://medium.com/_/api/topics/55f1c20aba7a/stream?limit=25"
	streamResponse, err := FetchStreamResponse(url)
	if err != nil {
		t.Fatal("Received non-nil error.")
	}

	posts := PostsFromResponse(streamResponse)
	if len(posts) != 25 {
		t.Fatalf("\nExpected: %d\nActual: %d\n", 25, len(posts))
	}

	p := posts[0]
	if p.Creator.ID == "" {
		t.Fatal("Creator ID is empty.")
	}
	if p.Creator.Name == "" {
		t.Fatal("Creator name is empty.")
	}
	if p.ID == "" {
		t.Fatal("Post ID is empty.")
	}
	if p.Title == "" {
		t.Fatal("Post title is empty.")
	}
	if p.UniqueSlug == "" {
		t.Fatal("Post unique slug is empty.")
	}
	if p.URL == "" {
		t.Fatal("Post URL is empty.")
	}

	time.Sleep(2 * time.Second)
}

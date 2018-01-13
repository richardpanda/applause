package medium

import (
	"testing"
	"time"
)

func TestFetchStreamResponse(t *testing.T) {
	url := "https://medium.com/_/api/topics/55f1c20aba7a/stream?limit=25"
	streamResponse, err := FetchStreamResponse(url)
	if err != nil {
		t.Fatal("Received non-nil error.")
	}

	want := 25
	got := len(streamResponse.Payload.References.Post)
	if got != want {
		t.Fatalf("\nExpected: %d\nActual: %d\n", want, got)
	}

	time.Sleep(2 * time.Second)
}

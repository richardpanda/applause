package medium

import (
	"log"
	"testing"
	"time"
)

func TestFetchTopicIDs(t *testing.T) {
	topicIDs, err := FetchTopicIDs()
	if err != nil {
		log.Fatal("Received non-nil error.")
	}

	numTopics := 52
	if len(topicIDs) != numTopics {
		t.Fatalf("\nExpected: %d\nActual: %d\n", numTopics, len(topicIDs))
	}

	got := topicIDs["software-engineering"]
	want := "55f1c20aba7a"
	if got != want {
		t.Fatalf("\nExpected: %s\nActual: %s\n", want, got)
	}

	time.Sleep(2 * time.Second)
}

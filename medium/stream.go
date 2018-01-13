package medium

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
		Paging struct {
			Next struct {
				To string `json:"to"`
			} `json:"next"`
		} `json:"paging"`
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

func FetchStreamResponse(url string) (*StreamResponse, error) {
	resp, err := sendGetRequest(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Remove "])}while(1);</x>" from beginning of slice
	b = b[16:]
	var streamResponse StreamResponse
	err = json.Unmarshal(b, &streamResponse)
	if err != nil {
		return nil, err
	}

	return &streamResponse, nil
}

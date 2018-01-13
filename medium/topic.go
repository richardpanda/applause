package medium

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	topicIDRegex = regexp.MustCompile(`-(\w+)-`)
	topicIDs     = map[string]string{
		"wellness":             "3d18b94f6858",
		"basic-income":         "dfcc5a84e698",
		"education":            "a205669c739e",
		"media":                "403f77178792",
		"humor":                "62b70885ee2b",
		"music":                "308a25bedcb5",
		"neuroscience":         "b315287d189e",
		"cybersecurity":        "d4e7f4144ac5",
		"digital-design":       "22d054b693da",
		"javascript":           "63c6f1f93ee",
		"programming":          "decb52b64abf",
		"technology":           "f862bfc84e38",
		"film":                 "a284e5a75d28",
		"business":             "7b2438b07d33",
		"productivity":         "8a146bc21b28",
		"creativity":           "24853457a119",
		"philosophy":           "40997f26fa15",
		"entrepreneurship":     "ba2d3ab15ed6",
		"cryptocurrency":       "9213b0063bcc",
		"self":                 "aef1078a3ef5",
		"lit":                  "76d64bb2132c",
		"photography":          "76d56a8194b9",
		"social-media":         "bceaf21c0fb7",
		"freelancing":          "29bb77781daf",
		"software-engineering": "55f1c20aba7a",
		"world":                "95e38a3034fb",
		"science":              "9ff4c9770e22",
		"family":               "e9a66523a5d6",
		"relationships":        "830cded25262",
		"environment":          "f721a1120833",
		"equality":             "ff18cfb862d2",
		"art":                  "2cdb28854f0c",
		"work":                 "af49579e220a",
		"artificial-intelligence": "1af65db9c2f8",
		"future":                  "bd856b86de98",
		"history":                 "aa97b7f5ff87",
		"politics":                "4d562ee63426",
		"marketing":               "4861fee224fd",
		"space":                   "9f59f758e8e0",
		"travel":                  "4ae32b074351",
		"economy":                 "40c8e34e04ce",
		"health":                  "d61cf867d93f",
		"comics":                  "b559a1d858b1",
		"culture":                 "1f79d9387f85",
		"food":                    "935e48590732",
		"mental-health":           "80e54e691fc9",
		"psychology":              "8c44fd843e59",
		"sexuality":               "e15e46793f8d",
		"spirituality":            "230848015f8c",
		"sports":                  "205a8f70c356",
		"data-science":            "ae5d4995e225",
		"math":                    "7808efc0cf94",
	}
)

func FetchTopicIDs() (map[string]string, error) {
	out := make(map[string]string)
	resp, err := sendGetRequest("https://medium.com/topics")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	doc.Find(".u-flexColumn").Each(func(i int, s *goquery.Selection) {
		name := strings.Replace(strings.ToLower(s.Find("a").Text()), " ", "-", -1)
		dataActionSource, _ := s.Find("button").Attr("data-action-source")
		id := topicIDRegex.FindStringSubmatch(dataActionSource)[1]
		out[name] = id
	})

	return out, nil
}

func GetTopicID(name string) (string, bool) {
	id, ok := topicIDs[name]
	return id, ok
}

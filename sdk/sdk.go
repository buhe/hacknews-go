package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const topURL string = "https://hacker-news.firebaseio.com/v0/topstories.json"
const itemURL string = "https://hacker-news.firebaseio.com/v0/item/%d.json"

type Story struct {
	Title string
	Time  int
	By    string
	Url   string
	Score int
}

func FetchTitles(max int) {
	resp, err := http.Get(topURL)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	var ids = make([]int, 50)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&ids)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	var stories = make([]Story, max)

	for index, id := range ids {
		resp, err := http.Get(fmt.Sprintf(itemURL, id))
		var story = new(Story)
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&story)
		if err != nil {
			fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		}
		stories[index] = *story
		if index >= max-1 {
			break
		}
	}
	for index, story := range stories {
		fmt.Println(index, ". ", story.Title, " > ", story.Url)
	}

}

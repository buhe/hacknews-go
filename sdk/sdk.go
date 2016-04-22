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

func FetchTitles(max int) []Story {
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

	//var stories = make([]Story, max)
	channel := make(chan Story, max)
	//fmt.Println(ids);
	for index, id := range ids {
		go func(tid int) {
			//直接传进去id 不行
			//fmt.Println(id);
			resp, err := http.Get(fmt.Sprintf(itemURL, tid))
			var story = new(Story)
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&story)
			if err != nil {
				fmt.Printf("%T\n%s\n%#v\n", err, err, err)
			}
			//stories[index] = *story
			channel <- *story;
		}(id)
		if index >= max - 1 {
			break
		}
	}
	result := make([] Story, max);
	for i := 0; i < max; i++ {
		//顺序有问题了啊.要搞一下
		result[i] = <-channel
		//result[i] = fmt.Sprintf("%d. %s" ,i + 1, story.Title);
		//fmt.Println(i, ". ", story.Title, " > ", story.Url)

	}
	return result;

}


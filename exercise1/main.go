package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Photo Photo struct
type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnaiUrl"`
}

// Doer Do interface
type Doer interface {
	Do(string) (*http.Response, error)
	Dont()
}

// DoGet a
type DoGet struct {
	Doer
}

// Do a
func (t DoGet) Do(url string) (*http.Response, error) {
	return http.Get(url)
}

// getTypicode a
type getTypicode struct {
	client Doer
	url    string
}

// GetPhoto get photo data
func (t getTypicode) GetPhoto(p *[]Photo) error {
	resp2, err := t.client.Do(t.url)
	if err != nil {
		return err
	}
	defer resp2.Body.Close()
	return json.NewDecoder(resp2.Body).Decode(p)
}

func main() {
	photo := []Photo{}
	t := getTypicode{
		client: DoGet{},
		url:    "https://jsonplaceholder.typicode.com/photos",
	}
	err := t.GetPhoto(&photo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(photo)
	return
}

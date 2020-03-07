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

// GetPhoto get photo data
func GetPhoto(p *[]Photo) error {
	url := "https://jsonplaceholder.typicode.com"
	resp2, err := http.Get(url + "/photos")
	if err != nil {
		return err
	}
	defer resp2.Body.Close()
	return json.NewDecoder(resp2.Body).Decode(p)
}

func main() {
	photo := []Photo{}
	err := GetPhoto(&photo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(photo)
	return
}

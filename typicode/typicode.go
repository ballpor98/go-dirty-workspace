package typicode

import (
	"encoding/json"
	"net/http"
)

type Album struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Doer interface {
	Do(url string) (*http.Response, error)
}

type DoGet struct {
	Doer
}

func (d DoGet) Do(url string) (*http.Response, error) {
	return http.Get(url)
}

type getTypicode struct {
	url    string
	client Doer
}

func (tc getTypicode) Get(p interface{}) error {
	resp, err := tc.client.Do(tc.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(p)
}

func New(path string) getTypicode {
	return getTypicode{
		url:    "https://jsonplaceholder.typicode.com" + path,
		client: DoGet{},
	}
}

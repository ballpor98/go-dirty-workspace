package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockDoGet struct {
	Doer
}

func (m mockDoGet) Do(url string) (*http.Response, error) {
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`[
			{
			  "albumId": 1,
			  "id": 1,
			  "title": "accusamus beatae ad facilis cum similique qui sunt",
			  "url": "https://via.placeholder.com/600/92c952",
			  "thumbnailUrl": "https://via.placeholder.com/150/92c952"
			}]`))),
	}, nil
}

func TestGetPhoto(t *testing.T) {
	photo := []Photo{}
	tc := getTypicode{
		client: mockDoGet{},
	}
	tc.GetPhoto(&photo)
	if photo[0].ID != 1 {
		t.Errorf("id: %d", photo[0].ID)
	}
}

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/photos", func(w http.ResponseWriter, r *http.Request) {
		pht := []byte(`{id: 1}`)
		w.WriteHeader(http.StatusOK)
		w.Write(pht)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

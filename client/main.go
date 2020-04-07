package main

import (
	"fmt"
	"go-dirty-workspace/typicode"
	"log"
)

func main() {
	p := []typicode.Photo{}
	t := typicode.NewTypicode("/photos")
	err := t.Get(&p)

	tc := typicode.NewTypicode("/albums")
	a := []typicode.Album{}
	err = tc.Get(&a)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}

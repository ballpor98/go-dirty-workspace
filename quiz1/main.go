package main

import "fmt"

func main() {
	// fmt.Println("print")
	fruits := []string{"a", "b", "c", "d", "e"}
	// temp := make([]string, len(fruits), cap(fruits))
	// copy(temp, fruits)
	temp := fruits[1:3:3]
	fmt.Printf("%d %d\n", len(temp), cap(temp))
	temp = append(temp, "f", "f", "f", "f", "f")
	fmt.Printf("%d %d\n", len(temp), cap(temp))
	fmt.Println(temp)
}

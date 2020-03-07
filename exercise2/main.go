package main

import (
	"fmt"
)

// BusinessError .
type BusinessError struct {
	err string
}

func (b BusinessError) Error() string {
	return "error: " + b.err
}
func main() {
	err := &BusinessError{err: "my error"}
	fmt.Println(err.Error())
}

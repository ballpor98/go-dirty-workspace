package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func hit(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(time.Second)
		table <- ball
	}
}
func main() {
	table := make(chan *Ball)
	go hit("P1", table)
	go hit("P2", table)
	table <- &Ball{}
	time.Sleep(100 * time.Second)
	<-table
}

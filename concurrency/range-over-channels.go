package main

import "fmt"

func rangeOver(){
	queue := make(chan string, 2)
	queue<- "hello"
	queue <- "world"
	close(queue)

	for q := range queue{
		fmt.Printf("got %s from channel\n", q)
	}
}


package main

import (
	"fmt"
)

func main(){
	ch1 := sq(gen(2,3,4,5,6,7))
	ch2 := sq(gen(23,24,25,26))
	done := make(chan struct{})
	for c := range merge(done, ch1,ch2){
		fmt.Printf("squares %v\n", c)
	}

	//waitGroup()
}


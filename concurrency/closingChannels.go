package main

import "fmt"

func jobs(){
	j := make(chan string, 5)
	done := make(chan bool)
	go func(){
		for{

			job, ok := <-j
			if ok{
				fmt.Printf("Got job %s \n", job)
			} else {
				fmt.Printf("finished job\n")
				done<- true
			}
		}
	}()

	for i:=0; i<5;i++{
		j<-fmt.Sprintf("got job %d done\n", i)
	}
	close(j)
	<-done
}

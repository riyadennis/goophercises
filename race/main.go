package main

import (
	"fmt"
	"time"
	"sync"
)

var balance = 100

func main(){
	var wt sync.WaitGroup
	wt.Add(3)
	go debit(10, &wt)
	go debit(20, &wt)
	go debit(10, &wt)
	wt.Wait()
	fmt.Printf("balance %v", balance)
}


func debit(amount int, wt *sync.WaitGroup){
	b := balance
	b -= amount
	balance = b
	time.Sleep(100)
	wt.Done()
}
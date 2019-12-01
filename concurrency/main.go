package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	msg  string
	wait chan bool
}

func main() {
	ch := fanIn(boring("joe"), boring("ann"))
	for i := 0; i < 10; i++ {
		msg := <-ch
		fmt.Printf("you say %s\n", msg.msg)
		msg.wait <- true
	}
	fmt.Println("just leave")
}

func fanIn(in1 <-chan message, in2 <-chan message) <-chan message {
	ch := make(chan message)
	go func() {
		for {
			select {
			case s := <-in1:
				ch <- s
			case s := <-in2:
				ch <- s
			case <-time.After(1 * time.Second):
				fmt.Println("timed out")
				return
			}
		}
	}()
	return ch
}

func boring(msg string) <-chan message {
	c := make(chan message)
	wait := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			msg := message{
				msg:  fmt.Sprintf("%s : %d", msg, i),
				wait: wait,
			}
			c <- msg
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			<-wait
		}
	}()
	return c
}

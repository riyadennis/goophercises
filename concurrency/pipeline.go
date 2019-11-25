package main

import "sync"

func merge(done <-chan struct{}, ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}
	wg.Add(len(ch))
	for _, c := range ch {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func gen(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}
func sq(ch <-chan int) <-chan int {
	s := make(chan int)
	go func() {
		for c := range ch {
			s <- c * c
		}
		close(s)
	}()
	return s
}

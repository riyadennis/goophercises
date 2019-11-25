package main

import (
	"fmt"
	"net/http"
	"sync"
)

func waitGroup() {
	var wg sync.WaitGroup
	url := []string{
		"https://blog.golang.org/context",
		"https://google.co.uk",
	}
	wg.Add(len(url))
	for _, u := range url {
		go func(u string) {
			defer wg.Done()
			res, _ := http.Get(u)
			if res.StatusCode == http.StatusOK {
				fmt.Printf("%s is working\n", u)
			}
		}(u)
	}
	wg.Wait()
}

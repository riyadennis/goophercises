package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)
type data struct{
	cnt int
	name string
}
func fetchData(filenames []string){
	ch := make(chan map[string]*data)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		for _, n := range filenames{
			checkDuplicate(n, ch)
		}
		wg.Done()
		close(ch)
	}()

	for l, d := range <-ch{
		if d.cnt > 1{
			fmt.Printf("duplicated line: %q, %d times in file %s\n", l, d.cnt, d.name)
		}
	}
	wg.Wait()
}

func checkDuplicate(name string, ch chan map[string]*data) <-chan map[string]*data{
	m := make(map[string]*data)
	content, err := ioutil.ReadFile(name)
	if err != nil{
		fmt.Fprintf(os.Stdout,
			"error opening %v file %s ",  err, name)
	}
	for _, l := range strings.Split(string(content), "\n"){
		if m[l] != nil{
			m[l].cnt = m[l].cnt+1
		} else{
			m[l] = &data{
				cnt:1,
				name:name,
			}
		}
	}
	ch<-m
	return ch
}

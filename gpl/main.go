package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
type data struct{
	cnt int
	name string
}
func main(){
	m := make(map[string]*data)
	for _, n := range os.Args[1:]{
		cntnt, err := ioutil.ReadFile(n)
		if err != nil{
			fmt.Fprintf(os.Stdout,
				"error opening %v file %s ",  err, n)
		}
		for _, l := range strings.Split(string(cntnt), "\n"){
			if m[l] != nil{
				m[l].cnt = m[l].cnt+1
			} else{
				m[l] = &data{
					cnt:1,
					name:n,
				}
			}
		}
	}

	for l, d := range m{
		if d.cnt > 1{
			fmt.Printf("duplicated line: %q, %d times in file %s\n", l, d.cnt, d.name)
		}
	}

}

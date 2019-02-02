package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fp, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	r, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	fileString := string(r)
	lines := strings.Split(string(fileString), "\n")
	for i := 0; i < len(lines); i++ {
		qans := strings.Split(lines[i], ",")
		fmt.Printf("question: %v, answer: %v \n", qans[0], qans[1])
	}
}

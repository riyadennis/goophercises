package main

import (
	"fmt"

	"github.com/gorilla/schema"
)

type Person struct {
	fname  string
	lname  string
	email  string
	mobile string
}

func main() {
	var decoder = schema.NewDecoder()
	persons := make(map[string][]string)
	var p = make(map[string][]*Person)
	persons["employees"] = []string{
		"22",
		"eee",
	}
	decoder.Decode(p, persons)
	fmt.Printf("%#v", p)
}

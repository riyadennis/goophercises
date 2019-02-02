package main

import (
	"log"

	"github.com/riyadennis/goophercises/lib"
)

func main() {
	fileContent, err := lib.Read("data/problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	qans, err := lib.CSVSplitter(fileContent)
	if err != nil {
		log.Fatal(err)
	}
	lib.WriteQuestionAnswer(qans)
}

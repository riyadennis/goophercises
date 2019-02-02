package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/riyadennis/goophercises/lib"
)

var fileName = flag.String("fileName",
	"data/problems.csv",
	"name of the file that has your quiz contents",
)

func main() {
	flag.Parse()
	fileContent, err := lib.Read(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	qans, err := lib.CSVSplitter(fileContent)
	if err != nil {
		log.Fatal(err)
	}
	points, err := lib.CheckQuestionAnswer(os.Stdin, qans)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("you got %v points", points)
}

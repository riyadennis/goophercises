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
	fileContentArray, err := lib.ReadCsvFile(*fileName)
	if err != nil {
		log.Fatalf("unable to open % v:: %v", *fileName, err)
	}
	qans, err := lib.NewQuestionAnswer(fileContentArray)
	if err != nil {
		log.Fatalf("invalid csv file :: %v", err)
	}
	points, err := lib.CheckQuestionAnswer(os.Stdin, qans)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("you got %v points", points)
}

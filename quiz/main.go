package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/riyadennis/goophercises/quiz/lib"
)

var (
	fileName = flag.String(
		"fileName",
		"/quiz/data/problems.csv",
		"name of the file that has your quiz contents")
	timeLimit = flag.Int(
		"timelimit",
		30,
		"time limit for the quiz")
)

func main() {
	flag.Parse()
	vpwd, err  := os.Getwd()
	if err != nil {
		log.Fatalf("unable to directory %v", err)
	}
	fileContentArray, err := lib.ReadCsvFile(vpwd+*fileName)
	if err != nil {
		log.Fatalf("unable to open % v:: %v", *fileName, err)
	}
	qans, err := lib.NewQuestionAnswer(fileContentArray)
	if err != nil {
		log.Fatalf("invalid csv file :: %v", err)
	}

	points, err := lib.CheckQuestionAnswer(os.Stdin, qans, *timeLimit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("you got %v points", points)
}

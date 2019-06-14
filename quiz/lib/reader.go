package lib

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

// QuestionAnswer holds question number, question and answer
type QuestionAnswer struct {
	Num              int
	Question, Answer string
}

// ReadCsvFile reads csv file and return its content
func ReadCsvFile(fileName string) ([][]string, error) {
	fp, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	return reader.ReadAll()
}

// NewQuestionAnswer gets csv string and creates an
// array of struct with questions and answers.
func NewQuestionAnswer(fileContent [][]string) ([]*QuestionAnswer, error) {
	if fileContent == nil {
		return nil, errors.New("invalid csv content")
	}
	var qans = make([]*QuestionAnswer, len(fileContent))
	for lineNum, line := range fileContent {
		// we dont want to check empty lines
		if line != nil {
			qans[lineNum] = &QuestionAnswer{
				Num:      lineNum + 1,
				Question: line[0],
				Answer:   strings.TrimSpace(line[1]),
			}
		}
	}
	return qans, nil
}

package lib

import (
	"encoding/csv"
	"errors"
	"os"
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
	var qans = []*QuestionAnswer{}
	for lineNum, line := range fileContent {
		// we dont want to check empty lines
		if line != nil {
			qans = append(qans, &QuestionAnswer{
				Num:      lineNum + 1,
				Question: line[0],
				Answer:   line[1],
			})
		}
	}
	return qans, nil
}

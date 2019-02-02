package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type QuestionAnswer struct {
	Num              int
	Question, Answer string
}

// Read reads and file and return its content
func Read(fileName string) (string, error) {
	fp, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	r, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", err
	}
	return string(r), nil
}

// CSVSplitter gets csv string and creates an
// array of struct with questions and answers.
func CSVSplitter(csv string) ([]*QuestionAnswer, error) {
	lines := strings.Split(csv, "\n")
	var qans = []*QuestionAnswer{}
	for lineNum, line := range lines {
		// we dont want to check empty lines
		if line != "" {
			qansArray := strings.Split(line, ",")
			if len(qansArray) != 2 {
				return nil, errors.New("invalid line in csv")
			}
			qans = append(qans, &QuestionAnswer{
				Num:      lineNum,
				Question: qansArray[0],
				Answer:   qansArray[1],
			})
		}
	}
	return qans, nil
}

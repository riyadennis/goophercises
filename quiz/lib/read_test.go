package lib

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRead(t *testing.T) {
	scenarios := []struct {
		name        string
		fileName    string
		expectedErr string
	}{
		{
			"invalidFile",
			"invalid.csv",
			"open invalid.csv: no such file or directory",
		},
		{
			"validFile",
			"../data/problems.csv",
			"",
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			_, err := ReadCsvFile(sc.fileName)
			if err != nil {
				checkErr(t, err, sc.expectedErr)
			}
		})
	}
}

func TestCSVSplitter(t *testing.T) {
	scenarios := []struct {
		name         string
		csvString    [][]string
		expectedErr  string
		expectedQans []*QuestionAnswer
	}{
		{
			"validCSV",
			[][]string{
				{"q1", "a1"},
				{"q2", "a2"},
			},
			"",
			[]*QuestionAnswer{
				{
					Num:      1,
					Question: "q1",
					Answer:   "a1",
				},
				{
					Num:      2,
					Question: "q2",
					Answer:   "a2",
				},
			},
		},
		{
			"inValidCSV",
			nil,
			"invalid csv content",
			nil,
		},
	}
	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			actualQns, err := NewQuestionAnswer(sc.csvString)
			if err != nil {
				if err.Error() != sc.expectedErr {
					t.Errorf("error wanted %v, got %v", err.Error(), sc.expectedErr)
				}
			}
			if !cmp.Equal(actualQns, sc.expectedQans) {
				t.Errorf(" mismatch %v", cmp.Diff(actualQns, sc.expectedQans))
			}
		})
	}
}

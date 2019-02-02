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
			"problems.csv",
			"",
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			cnt, err := Read(sc.fileName)
			if err == nil && cnt == "" {
				t.Error("empty file")
			}
			if err != nil && err.Error() != sc.expectedErr {
				t.Errorf("error got %v, want %v", err.Error(), sc.expectedErr)
			}
		})
	}
}

func TestCSVSplitter(t *testing.T) {
	scenarios := []struct {
		name         string
		csvString    string
		expectedErr  string
		expectedQans []*QuestionAnswer
	}{
		{
			"validCSV",
			"q1,a1\nq2,a2\n",
			"",
			[]*QuestionAnswer{
				{
					Num:      0,
					Question: "q1",
					Answer:   "a1",
				},
				{
					Num:      1,
					Question: "q2",
					Answer:   "a2",
				},
			},
		},
		{
			"inValidCSV",
			"q1\nq2,a2\n",
			"invalid line in csv",
			nil,
		},
	}
	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			actualQns, err := CSVSplitter(sc.csvString)
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

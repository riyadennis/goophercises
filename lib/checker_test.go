package lib

import (
	"bytes"
	"io"
	"testing"
)

func TestCheckQuestionAnswer(t *testing.T) {
	scenarios := []struct {
		name            string
		input           io.Reader
		questionAnswers []*QuestionAnswer
		expectedPoints  int
		expectedErr     string
	}{
		{
			"invalidAnswers",
			bytes.NewReader([]byte("invalid answers")),
			[]*QuestionAnswer{
				{
					Num:      1,
					Question: "q1",
					Answer:   "a1",
				},
			},
			0,
			"",
		},
		{
			"validAnswer",
			bytes.NewReader([]byte("a1")),
			[]*QuestionAnswer{
				{
					Num:      1,
					Question: "q1",
					Answer:   "a1",
				},
			},
			1,
			"",
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			points, err := CheckQuestionAnswer(
				scenario.input,
				scenario.questionAnswers,
			)
			if points != scenario.expectedPoints {
				t.Errorf("points wanted %v, got %v", points, scenario.expectedPoints)
			}
			checkErr(t, err, scenario.expectedErr)
		})
	}
}
func checkErr(t *testing.T, actualErr error, expectedErr string) {
	if actualErr == nil && expectedErr != "" {
		t.Errorf("error expected %v but got none", expectedErr)
	}
	if actualErr != nil {
		if expectedErr == "" {
			t.Errorf("unexpected error %v", actualErr.Error())
		}
		if actualErr.Error() != expectedErr {
			t.Errorf("error expected %v, got %v", actualErr.Error(), expectedErr)
		}
	}
}

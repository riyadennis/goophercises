package lib

import "fmt"

// WriteQuestionAnswer displays questions to the user to answer
func WriteQuestionAnswer(qans []*QuestionAnswer) {
	for _, q := range qans {
		fmt.Printf("%d Question %s\n", q.Num, q.Question)
	}
}

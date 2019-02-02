package lib

import (
	"bufio"
	"fmt"
	"io"
)

// CheckQuestionAnswer displays questions to the user to answer
func CheckQuestionAnswer(input io.Reader, qans []*QuestionAnswer) (int, error) {
	points := len(qans)
	reader := bufio.NewReader(input)
	for _, q := range qans {
		fmt.Printf("%d) %s \n", q.Num, q.Question)
		answer, _, err := reader.ReadLine()
		if err != nil {
			return 0, err
		}
		if string(answer) != q.Answer {
			points--
			fmt.Printf("correct answer:%v \n", q.Answer)
		}
	}
	return points, nil
}

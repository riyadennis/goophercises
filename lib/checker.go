package lib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

// CheckQuestionAnswer displays questions to the user to answer
func CheckQuestionAnswer(input io.Reader, qans []*QuestionAnswer, timeLimit int) (int, error) {
	points := 0
	reader := bufio.NewReader(input)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	for _, q := range qans {
		fmt.Printf("%d) %s \n", q.Num, q.Question)
		answerCh := make(chan string)
		go func() {
			answer, _, err := reader.ReadLine()
			if err == nil {
				answerCh <- string(answer)
			}
		}()
		select {
		case <-timer.C:
			return points, nil
		case answer := <-answerCh:
			if strings.TrimSpace(answer) != q.Answer {
				fmt.Printf("correct answer:%v \n", q.Answer)
			}
			points++
		}
	}
	return points, nil
}

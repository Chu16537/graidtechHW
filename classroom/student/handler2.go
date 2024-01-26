package student

import (
	"fmt"
	"time"
)

func (s *student) enable() {
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()

	s.updateActionTime()

	nowAnsQuestionIdx := 0

	for {
		select {
		case <-s.ctx.Done():
			fmt.Println(fmt.Sprintf("student %v done", s.name))
			return
		case <-tick.C:
			_, isCorrect, _ := s.createAns(nowAnsQuestionIdx)
			// 代表有回答正確
			if isCorrect {
				nowAnsQuestionIdx++
				break
			}
			// 取得贏家
			q := s.question.GetQuestion(nowAnsQuestionIdx)
			if q == nil {
				break
			}

			if q.Winner.Name != "" {
				nowAnsQuestionIdx++
			}

			if q.Winner.Id != -1 {
				s.sayWin(q.Winner.Id, q.Winner.Name)
			}

		}
	}
}

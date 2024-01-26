package teacher

import (
	"fmt"
	"time"
)

func (t *teacherHandler) enable() {
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-t.ctx.Done():
			fmt.Println(fmt.Sprintf("Teacher enable done"))
			return
		case <-tick.C:
			// 執行動作

			t.enableSayAns()

			t.sayReady()

			time.Sleep(2 * time.Second)

			t.createQuestion()
		}
	}
}

// 公布答案
func (t *teacherHandler) enableSayAns() {
	t.lock.Lock()
	defer t.lock.Unlock()

	for i, v := range t.ansQuestionMap {
		if v == t.studentCount {
			t.sayAns(i)
		}
		q := t.question.GetQuestion(i)
		if q != nil && q.Winner.Name != "" {
			delete(t.ansQuestionMap, i)
		}
	}
}

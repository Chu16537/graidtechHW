package question

import "fmt"

// 創建題目
func (q *questionHandler) CreateQuestion(x, y, ans int, symbo string) *Question {
	qq := &Question{
		Num1:        x,
		Num2:        y,
		Ans:         ans,
		Sybo:        symbo,
		QuestionStr: fmt.Sprintf("%v %v %v = ?", x, symbo, y),
		Winner: &Winner{
			Id: -1,
		},
	}

	q.questions = append(q.questions, qq)

	return qq
}

// 確認題目答案
func (q *questionHandler) CheckAns(studentId int, studentName string, idx int, ans int) bool {
	// 沒有題目
	if idx >= len(q.questions) || idx < 0 {
		return false
	}

	// 回答正確 記錄學生姓名
	if ans == q.questions[idx].Ans {
		q.questions[idx].Winner.Id = studentId
		q.questions[idx].Winner.Name = studentName
		return true
	}

	return false
}

// 取得指定題目
func (q *questionHandler) GetQuestion(idx int) *Question {
	if len(q.questions) <= idx {
		return nil
	}
	return q.questions[idx]
}

// 是否可以回答題目
func (q *questionHandler) IsAnswer(idx int) bool {
	// 沒有題目 || 或有人回答正確了
	if len(q.questions) <= idx || q.questions[idx] == nil || q.questions[idx].Winner.Name != "" {
		return false
	}

	return true
}

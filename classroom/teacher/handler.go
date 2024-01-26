package teacher

import (
	"fmt"
	"math/rand"
)

// 確認準備
func (t *teacherHandler) sayReady() {
	fmt.Println("Teacher:Guys, are you ready?")
}

// 創建題目
func (t *teacherHandler) createQuestion() {
	n1 := rand.Intn(1001)
	n2 := rand.Intn(1001)
	symbo := getSymbo()
	ans := getAns(n1, n2, symbo)
	q := t.question.CreateQuestion(n1, n2, ans, symbo)

	fmt.Println("Teacher:", q.QuestionStr)
}

// 取得符號
func getSymbo() string {
	symboInt := rand.Intn(1001) % 4

	switch symboInt {
	case 0:
		return "+"
	case 1:
		return "-"
	case 2:
		return "*"
	case 3:
		return "/"
	default:
		return "+"
	}
}

// 取得答案
func getAns(n1, n2 int, symbo string) int {
	switch symbo {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	default:
		return 0
	}
}

// 學生向老師提供答案
func (t *teacherHandler) sandAns(studentId int, studentName string, queIdx int, ans int) bool {
	// 上鎖
	t.lock.Lock()
	defer t.lock.Unlock()

	if !t.question.IsAnswer(queIdx) {
		return false
	}

	q := t.question.GetQuestion(queIdx)
	fmt.Println(fmt.Sprintf("Student %v %v: %v %v %v = %v!", studentName, queIdx, q.Num1, q.Sybo, q.Num2, ans))
	t.ansQuestionMap[queIdx]++

	if t.question.CheckAns(studentId, studentName, queIdx, ans) {
		fmt.Println("Teacher:", studentName, ", you are right!")
		return true
	}

	fmt.Println("Teacher:", studentName, ", you are wrong.")
	return false
}

// 公布答案
func (t *teacherHandler) sayAns(idx int) {
	q := t.question.GetQuestion(idx)
	// 沒有題目
	if q == nil {
		fmt.Println("Teacher: not question ")
		return
	}

	q.Winner.Name = "Teacher"
	fmt.Println("Teacher: Boooo~ Answer is ", q.Ans)
}

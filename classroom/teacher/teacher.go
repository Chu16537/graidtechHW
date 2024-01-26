package teacher

import (
	"context"
	"graidtechHW/classroom/question"
	"sync"
)

type ITeacher interface {
	// 確認準備
	SayReady()
	// 創建題目
	CreateQuestion()
	// 公布答案
	SayAns(idx int)
}

type IQuestion interface {
	// 創建題目
	CreateQuestion(x, y, ans int, symbo string) *question.Question
	// 確認答案
	CheckAns(studentId int, studentName string, queIdx int, ans int) bool
	// 取得題目
	GetQuestion(idx int) *question.Question
	// 是否可以回答題目
	IsAnswer(idx int) bool
}

type teacherHandler struct {
	ctx    context.Context
	cancel context.CancelFunc
	lock   sync.Mutex

	question       IQuestion
	studentCount   int
	ansQuestionMap map[int]int
}

func New1(q IQuestion) *teacherHandler {
	t := &teacherHandler{
		question:       q,
		ansQuestionMap: make(map[int]int),
	}

	return t
}

func New2(c context.Context, q IQuestion, sc int) *teacherHandler {
	ctx, cancel := context.WithCancel(c)
	t := &teacherHandler{
		ctx:            ctx,
		cancel:         cancel,
		question:       q,
		studentCount:   sc,
		ansQuestionMap: make(map[int]int),
	}

	go t.enable()

	return t
}

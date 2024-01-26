package student

import (
	"context"
	"graidtechHW/classroom/question"
	"time"
)

type student struct {
	ctx      context.Context
	cancel   context.CancelFunc
	question IQuestion
	teacher  ITeacher
	winMode  int

	id         int       // 學號
	name       string    // 學生名字
	actionTime time.Time // 執行動作時間
	isWin      bool

	ansMap map[int]int // 回答過的題目
}

type IStudent1 interface {
	GetName() string
	// 更新學生狀態
	UpdateStatus()
	// 是否可以作答
	IsCanCreateAns(idx int) bool
	// 作答
	CreateAns(idx int) (bool, bool, int)
	// 說 哪位學生正確
	SayWin(winId int, winName string)
}

type IQuestion interface {
	// 取得題目
	GetQuestion(int) *question.Question
	// 是否可以回答題目
	IsAnswer(int) bool
}

// 學生可以對老師做的事情
type ITeacher interface {
	// 提供答案
	SandAns(id int, name string, queIdx int, ans int) bool
}

func New1(id int, name string, q IQuestion, t ITeacher, winMode int) IStudent1 {
	s := new(student)

	s.question = q
	s.teacher = t
	s.winMode = winMode
	s.ansMap = make(map[int]int)
	s.id = id
	s.name = name

	return s
}
func New2(id int, name string, c context.Context, q IQuestion, t ITeacher, winMode int) {
	ctx, cancel := context.WithCancel(c)
	s := new(student)

	s.ctx = ctx
	s.cancel = cancel
	s.question = q
	s.teacher = t
	s.winMode = winMode
	s.ansMap = make(map[int]int)
	s.id = id
	s.name = name

	go s.enable()
}

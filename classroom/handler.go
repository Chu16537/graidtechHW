package classroom

import (
	"context"
	"graidtechHW/classroom/question"
	"graidtechHW/classroom/student"
	"graidtechHW/classroom/teacher"
	"graidtechHW/util"
)

func Model1() {
	// 選擇學生答題模式
	winMode := util.WIN_MODE_ALWAYS_WIN
	studentNames := []string{"A", "B", "C", "D", "E"}

	ctx, _ := context.WithCancel(context.Background())
	// defer cancel()

	q := question.New()
	t := teacher.New1(q)

	// 創建學生
	studens := make([]student.IStudent1, len(studentNames))
	for i, name := range studentNames {
		studens[i] = student.New1(i, name, q, t, winMode)
	}

	c := NewClassroomHandler(ctx, t, studens)

	// 第一種
	go c.Model1()
}

func Model2() {

	// 選擇學生答題模式
	winMode := util.WIN_MODE_ALWAYS_LOSE
	studentNames := []string{"A", "B", "C", "D", "E"}

	ctx, _ := context.WithCancel(context.Background())
	// defer cancel()

	q := question.New()
	t := teacher.New2(ctx, q, len(studentNames))

	// 創建學生

	for i, name := range studentNames {
		student.New2(i, name, ctx, q, t, winMode)
	}

}

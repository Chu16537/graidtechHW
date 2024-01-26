package classroom

import (
	"context"
	"fmt"
	"graidtechHW/classroom/student"
	"graidtechHW/classroom/teacher"
	"math/rand"
	"time"
)

type classroomHandler struct {
	ctx    context.Context
	cancel context.CancelFunc

	teacher  teacher.ITeacher
	students []student.IStudent1
}

func NewClassroomHandler(c context.Context, t teacher.ITeacher, s []student.IStudent1) *classroomHandler {
	ctx, cancel := context.WithCancel(c)

	return &classroomHandler{
		ctx:      ctx,
		cancel:   cancel,
		teacher:  t,
		students: s,
	}
}

// 等待所有學生答對 或是 所有人回答錯誤
func (c *classroomHandler) Model1() {

	// 執行次數
	// count := 30
	// 等待時間
	waitTime := 2 * time.Second

	// 作答人數
	ansMap := make(map[int]struct{})
	// 是否有學生正確
	isCorrect := false

	tick := time.NewTicker(1000 * time.Millisecond)
	defer tick.Stop()

	nowQuestionIdx := 0
	for {
		ansMap = make(map[int]struct{})
		isCorrect = false

		// 老師喊準備
		c.teacher.SayReady()
		// 等待
		time.Sleep(waitTime)
		// 創建題目
		c.teacher.CreateQuestion()

		// 更新學生狀態
		for _, v := range c.students {
			v.UpdateStatus()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Nanosecond)
		}

		for {
			// 所有學生回答完畢
			if len(ansMap) == len(c.students) {
				break
			}
			// 有人回答正確
			if isCorrect {
				break
			}

			select {
			case <-c.ctx.Done():
				fmt.Println("Model1 done")
				return

			case <-tick.C:
				// 學生回答
				for j, v := range c.students {
					//已經有答題就跳過
					if _, ok := ansMap[j]; ok {
						continue
					}
					// 不能作答跳過
					isCanCreateAns := v.IsCanCreateAns(nowQuestionIdx)
					if !isCanCreateAns {
						ansMap[j] = struct{}{}
						continue
					}

					_, isCorrect, _ = v.CreateAns(nowQuestionIdx)

					// 假如有同學答對
					if isCorrect {
						c.studentSayWin(j, v.GetName())
						break
					}

				}
			}
		}

		//  老師講答案
		if !isCorrect {
			c.teacher.SayAns(nowQuestionIdx)
		}

		nowQuestionIdx++
	}
}

func (c *classroomHandler) studentSayWin(winId int, winName string) {
	for _, v := range c.students {
		v.SayWin(winId, winName)
	}
}

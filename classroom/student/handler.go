package student

import (
	"fmt"
	"graidtechHW/util"
	"math/rand"
	"time"
)

// 更新動作時間
func (s *student) updateActionTime() {
	t := util.GetStudentThinkTime()
	d := time.Duration(t * int(time.Millisecond))
	s.actionTime = time.Now().Add(d)
}

func (s *student) updateWinMode() {
	if s.winMode == util.WIN_MODE_ALWAYS_WIN {
		s.isWin = true
	} else if s.winMode == util.WIN_MODE_ALWAYS_LOSE {
		s.isWin = false
	} else {
		// 隨機是否回答正確
		s.isWin = false
		if rand.Uint32()%2 == 0 {
			s.isWin = true
		}
	}
}

// 是否可以作答
func (s *student) isCanCreateAns(idx int) bool {
	// 有人答對了
	if !s.question.IsAnswer(idx) {
		return false
	}

	// 判斷是否有回答過題目
	if _, ok := s.ansMap[idx]; ok {
		return false
	}

	return true
}

// 作答
func (s *student) createAns(idx int) (bool, bool, int) {
	// 假如 學生要動作時間還沒超過現在時間 不做事
	if s.actionTime.After(time.Now()) {
		return false, false, 0
	}

	// 判斷是否有回答過題目
	if _, ok := s.ansMap[idx]; ok {
		return false, false, 0
	}

	// 確認題目是否可以作答
	if !s.question.IsAnswer(idx) {
		return false, false, 0
	}

	// 更新回答模式
	s.updateWinMode()

	// 取得當前題目
	q := s.question.GetQuestion(idx)

	ans := q.Ans
	// 假如要回答錯誤
	if !s.isWin {
		ans = rand.Int()
	}

	// 紀錄回答
	s.ansMap[idx] = ans

	return true, s.teacher.SandAns(s.id, s.name, idx, ans), ans

}

// 說 哪位學生正確
func (s *student) sayWin(winId int, winName string) {
	if winName != "" && winId != s.id {
		fmt.Println(fmt.Sprintf("Student %v: %v, you win.", s.name, winName))
	}
}

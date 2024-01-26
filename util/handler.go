package util

import (
	"math/rand"
	"time"
)

// 取得學生思考時間
func GetStudentThinkTime() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(STUDENT_THINK_MAX_TIME) + STUDENT_THINK_MIN_TIME
}

// 隨機數
func RandRange(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

// 當前時間毫秒
func GetNowMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

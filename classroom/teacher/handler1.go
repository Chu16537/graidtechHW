package teacher

// 確認準備
func (t *teacherHandler) SayReady() {
	t.sayReady()
}

// 創建題目
func (t *teacherHandler) CreateQuestion() {
	t.createQuestion()
}

// 學生向老師提供答案
func (t *teacherHandler) SandAns(studentId int, studentName string, queIdx int, ans int) bool {
	return t.sandAns(studentId, studentName, queIdx, ans)
}

// 公布答案
func (t *teacherHandler) SayAns(idx int) {
	t.sayAns(idx)
}

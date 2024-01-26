package student

// 取得名字
func (s *student) GetName() string {
	return s.name
}

// 更新學生狀態
func (s *student) UpdateStatus() {
	s.updateActionTime()
}

// 是否可以作答
func (s *student) IsCanCreateAns(idx int) bool {
	return s.isCanCreateAns(idx)
}

// 作答
func (s *student) CreateAns(idx int) (bool, bool, int) {
	return s.createAns(idx)
}

// 說 哪位學生正確
func (s *student) SayWin(winId int, winName string) {
	s.sayWin(winId, winName)
}

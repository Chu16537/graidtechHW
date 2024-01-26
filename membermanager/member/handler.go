package member

// 設定其他會員資料
func (m *Member) SetOtherMember(om map[int]IMemberRPC) {
	m.otherMemberMap = make(map[int]*OtherMember, len(om))

	for i, v := range om {
		if i == m.ID {
			continue
		}
		m.otherMemberMap[i] = &OtherMember{
			IMemberRPC:    v,
			lastHeartTime: 0,
			isOnline:      true,
		}
	}

}

// 取得在線會員數量
func (m *Member) GetAllOnlineCount() int {
	count := 0
	for _, v := range m.otherMemberMap {
		if v.isOnline {
			count++
		}
	}

	return count
}

func (m *Member) GetAllOnlineHalfCount() int {
	return m.GetAllOnlineCount()/2 + 1
}

package member

import (
	"fmt"
	"graidtechHW/membermanager/message"
	"strconv"
)

// 廣播給其他人
func (m *Member) broadcast(f func(o *OtherMember)) {
	// 是否斷線
	if m.isLineOff() {
		return
	}

	for i, v := range m.otherMemberMap {
		// 自己不用
		if i == m.ID {
			continue
		}

		f(v)
	}
}

// 投票
func (m *Member) Vote(req message.MessageReq) message.MessageRes {
	res := message.MessageRes{
		Idx:      req.Idx,
		IsOnline: true,
	}

	// 是否斷線
	if m.isLineOff() {
		res.IsOnline = false
		res.Err = fmt.Errorf("failed Vote with Member %v", m.ID)
		return res
	}

	// 沒有投票 & 沒有 leader
	if m.votedForID == -1 && m.leaderID == -1 {
		id, err := strconv.Atoi(req.Data)
		if err == nil {
			m.setVoteFor(id)
			res.Data = strconv.Itoa(m.ID)
			return res
		}
	}

	res.Err = fmt.Errorf("not Vote with Member %v", m.ID)
	return res
}

// 更新leader
func (m *Member) UpdateLeader(req message.MessageReq) message.MessageRes {
	res := message.MessageRes{
		Idx:      req.Idx,
		IsOnline: true,
	}

	// 是否斷線
	if m.isLineOff() {
		res.IsOnline = false
		res.Err = fmt.Errorf("failed UpdateLeader with Member %v", m.ID)
		return res
	}

	id, err := strconv.Atoi(req.Data)
	if err == nil {
		m.setLeader(id)
		res.Data = strconv.Itoa(m.ID)
		return res
	}

	m.reDefault()
	res.Err = fmt.Errorf("not UpdateLeader with Member %v", m.ID)
	return res
}

// 心跳包
func (m *Member) HeartBeatResponse(req message.MessageReq) message.MessageRes {
	res := message.MessageRes{
		Idx:      req.Idx,
		IsOnline: true,
		Data:     strconv.Itoa(m.ID),
	}

	// 是否斷線
	if m.isLineOff() {
		res.IsOnline = false
		res.Err = fmt.Errorf("failed HeartBeatResponse is line off with Member %v", m.ID)
		return res
	}

	id, err := strconv.Atoi(req.Data)
	if err == nil {
		fmt.Printf("member %v 收到 other member %v 心跳包\n", m.ID, id)
		return res
	}

	res.Err = fmt.Errorf("not HeartBeatResponse with Member %v", m.ID)
	return res
}

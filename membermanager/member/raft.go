package member

import (
	"fmt"
	"graidtechHW/membermanager/message"
	"graidtechHW/util"
	"strconv"
	"time"
)

// 檢查心跳
func (m *Member) checkHeart() {
	for {
		time.Sleep(1 * time.Second)

		if m.isLineOff() {
			continue
		}

		// 心跳超时
		if util.GetNowMillisecond()-m.leaderlastHeartTime > util.TIMEOUT_HEART*1000 {
			fmt.Printf("member %v 心跳超時，重新開始選舉\n", m.ID)
			m.reDefault()
			m.leaderlastHeartTime = util.GetNowMillisecond()
			// 準備選舉
			go m.readyElection()
		}
	}
}

// 發送心跳包
func (m *Member) sendHeartPacket() {

	f := func(o *OtherMember) {
		req := message.MessageReq{
			Data: strconv.Itoa(m.ID),
		}

		res := o.HeartBeatResponse(req)
		o.isOnline = res.IsOnline

		id, err := strconv.Atoi(res.Data)
		if err != nil {
			return
		}

		if !res.IsOnline {
			fmt.Printf("member %v %v\n", m.ID, res.Err.Error())
			// 假如是 leader斷線
			if id == m.leaderID {
				m.setLeader(-1)
			}

			o.lastHeartTime = util.GetNowMillisecond()

			return
		}

		if id == m.leaderID {
			m.leaderlastHeartTime = util.GetNowMillisecond()
		}
	}

	for {
		if m.isLineOff() {
			continue
		}

		if m.leaderID == m.ID {
			m.leaderlastHeartTime = util.GetNowMillisecond()
		}

		// 心跳廣播
		go m.broadcast(f)

		time.Sleep(time.Second * time.Duration(util.TIMEOUT_HEART_RATE))
	}
}

// 準備選舉
func (m *Member) readyElection() {
	for {
		if m.isLineOff() {
			return
		}

		// 嘗試變候選人
		if m.becomeCandidate() {
			// 變為候選人 發起投票
			if m.election() {
				return
			} else {
				continue
			}
		} else {
			// 不是候選人 || 投給別人 || 不是follower || 有leader
			return
		}
	}
}

// 準備成為候選人
func (m *Member) becomeCandidate() bool {
	r := util.RandRange(util.TIMEOUT_ELECTION_MIN*1000, util.TIMEOUT_ELECTION_MAX*1000)
	time.Sleep(time.Duration(r) * time.Millisecond)

	if m.isLineOff() {
		return false
	}

	//是follower & 沒有leader & 沒有投票
	if m.status == util.STATUS_FOLLOWER && m.leaderID == -1 && m.votedForID == -1 {
		//變為候選人
		fmt.Printf("member %v 變為候選人 \n", m.ID)

		m.SetStatus(util.STATUS_CANDIDATE)
		//投給自己
		m.setVoteFor(m.ID)
		m.voteAdd()
		m.termAdd()

		return true
	}
	return false
}

// 選舉 candidate
func (m *Member) election() bool {
	fmt.Printf("member %v 給所有人發起投票 \n", m.ID)
	f := func(o *OtherMember) {
		// 發送投票請求
		req := message.MessageReq{
			Data: strconv.Itoa(m.ID),
		}

		res := o.Vote(req)

		o.isOnline = res.IsOnline

		if res.Err == nil && res.Data != "" {
			m.voteChan <- true
			return
		}

		m.voteChan <- false

	}
	go m.broadcast(f)

	for {
		select {
		case <-time.After(time.Second * time.Duration(util.TIMEOUT_ELECTION)):
			//超時
			fmt.Printf("member %v 選舉超時，改為follower \n", m.ID)

			m.reDefault()
			return false
		case ok := <-m.voteChan:
			if ok {
				m.voteAdd()
			}

			// 票數超過一半 && 沒有leader
			if m.vote >= m.GetAllOnlineHalfCount() && m.leaderID == -1 {
				m.toLeader()
				return true
			}
		}
	}
}

// 成為leader
func (m *Member) toLeader() {
	m.SetStatus(util.STATUS_LEARDER)
	m.setLeader(m.ID)

	fmt.Printf("member %v 成為leader 得票數 %v 廣播我是leader\n", m.ID, m.vote)

	// 通知所有人更改leader
	f := func(o *OtherMember) {

		req := message.MessageReq{
			Data: strconv.Itoa(m.ID),
		}

		res := o.UpdateLeader(req)
		o.isOnline = res.IsOnline
	}
	go m.broadcast(f)

	// 發送心跳包
	// go m.sendHeartPacket()
}

package member

import (
	"fmt"
	"graidtechHW/membermanager/message"
	"graidtechHW/util"
	"sync"
)

type Member struct {
	lock                sync.Mutex
	ID                  int
	status              int // 狀態
	leaderID            int
	term                int   // 任期
	votedForID          int   // 投票給誰
	vote                int   // 獲得投票數
	leaderlastHeartTime int64 // 最后一次心跳的时间
	voteChan            chan bool

	otherMemberMap map[int]*OtherMember // 紀錄其他人員
}

// 其他會員資料
// 假如是用網路 會有紀錄其他IP 這邊用這個而已
type OtherMember struct {
	IMemberRPC
	lastHeartTime int64
	isOnline      bool
}

// RPC 接口 負責對其他人做的操作
type IMemberRPC interface {
	// 投票
	Vote(message.MessageReq) message.MessageRes
	// 更新leader
	UpdateLeader(message.MessageReq) message.MessageRes
	// 心跳包
	HeartBeatResponse(message.MessageReq) message.MessageRes
}

func New(id int) *Member {
	m := &Member{
		ID:         id,
		status:     util.STATUS_FOLLOWER,
		leaderID:   -1,
		term:       0,
		votedForID: -1,
		voteChan:   make(chan bool),
	}

	fmt.Printf("member %v Hi \n", id)
	go m.checkHeart()
	go m.sendHeartPacket()

	return m
}

// 更改狀態
func (m *Member) SetStatus(status int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.status = status
	fmt.Printf("member %v status %v \n", m.ID, status)
}

func (m *Member) GetStatus() int {
	return m.status
}

// 設定 leader
func (m *Member) setLeader(id int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.leaderID = id
	fmt.Printf("member %v update leader %v \n", m.ID, id)
}

// 設定 任期
func (m *Member) setTerm(id int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.term = id
}

// 任期累加
func (m *Member) termAdd() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.term++
}

// 設定投給誰
func (m *Member) setVoteFor(id int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.votedForID = id
	fmt.Printf("member %v 投票給 %v \n", m.ID, id)
}

// 設定投票數量
func (m *Member) setVote(num int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.vote = num
}

// 投票累加
func (m *Member) voteAdd() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.vote++
	fmt.Printf("member %v 當前票數 %v \n", m.ID, m.vote)
}

// 預設
func (m *Member) reDefault() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.vote = 0
	m.votedForID = -1
	m.status = util.STATUS_FOLLOWER
}

// 是否斷線
func (m *Member) isLineOff() bool {
	return m.status == util.STATUS_DIE
}

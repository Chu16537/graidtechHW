package membermanager

import (
	"fmt"
	"graidtechHW/membermanager/member"
	"graidtechHW/util"
	"time"
)

type memberManager struct{}

var mm *memberManager
var memberMap map[int]member.IMemberRPC
var mMap map[int]*member.Member // 方便使用而已

func Enable() {
	memberIDs := []int{0, 1, 2, 3, 4}

	mm = new(memberManager)
	memberMap = make(map[int]member.IMemberRPC, len(memberIDs))
	mMap = make(map[int]*member.Member, len(memberIDs))

	// 創建會員
	for _, v := range memberIDs {
		m := member.New(v)

		memberMap[v] = m
		mMap[v] = m
	}

	for _, v := range mMap {
		v.SetOtherMember(memberMap)
	}

	// 模擬事件
	go sendEvent()

}

// 模擬事件
func sendEvent() {
	// 等待
	time.Sleep(20 * time.Second)

	// 刪除 follower
	for i, v := range mMap {
		if v.GetStatus() == util.STATUS_FOLLOWER {
			fmt.Printf("kill %v \n", i)
			v.SetStatus(util.STATUS_DIE)
			break
		}
	}

	// 等待
	time.Sleep(10 * time.Second)

	// 刪除 leader
	for i, v := range mMap {
		if v.GetStatus() == util.STATUS_LEARDER {
			v.SetStatus(util.STATUS_DIE)
			fmt.Printf("kill %v \n", i)
			break
		}
	}

}

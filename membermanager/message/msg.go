package message

type MessageReq struct {
	Idx  int    //  動作idx
	Data string // 資料
}

type MessageRes struct {
	Idx      int    //  動作idx
	Err      error  // 錯誤
	IsOnline bool   // 是否在線
	Data     string // 資料
}

package util

const (
	STUDENT_THINK_MIN_TIME = 1000 // 學生思考時間下限
	STUDENT_THINK_MAX_TIME = 3000 // 學生思考時間上限

	WIN_MODE_ALWAYS_WIN  = 0 // 贏
	WIN_MODE_ALWAYS_LOSE = 1 // 輸
	WIN_MODE_RANG        = 2 // 隨機

	STATUS_DIE       = -1 // 不存在
	STATUS_FOLLOWER  = 0  // 跟隨
	STATUS_CANDIDATE = 1  // 候選人
	STATUS_LEARDER   = 2  // 領導

	TIMEOUT_ELECTION     = 10 // 選舉超時時間  秒
	TIMEOUT_ELECTION_MAX = 15 // 等待選舉最大時間 毫秒
	TIMEOUT_ELECTION_MIN = 2  // 等待選舉最小時間 毫秒
	TIMEOUT_HEART        = 20 // 心跳超時 秒
	TIMEOUT_HEART_RATE   = 10 // 心跳頻率 秒

	DISK_LEN_MAX = 10
	DISK_LEN_MIN = 3
)

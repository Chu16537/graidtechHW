package question

type questionHandler struct {
	questions []*Question // 題庫
}

type Question struct {
	Num1        int
	Num2        int
	Ans         int     // 答案
	Sybo        string  // 符號
	QuestionStr string  // 題目字串
	Winner      *Winner // 回答正確的人
}

type Winner struct {
	Id   int
	Name string
}

func New() *questionHandler {
	return &questionHandler{
		questions: make([]*Question, 0, 100),
	}
}

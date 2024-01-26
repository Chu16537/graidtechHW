package raid

import "fmt"

type IRaid interface {
	// 寫入
	Write(position int, str string)
	// 讀取
	Read(position int, len int) string
	// 清除
	Clear(idx int)
}

func Enable() {
	// 寫入字串
	writeStr := "qawsedrftg"

	r0 := NewRaid0()
	r0.Write(0, writeStr)
	r0.Clear(0)
	fmt.Println("r0 結果 ", r0.Read(0, 9))

	r1 := NewRaid1()
	r1.Write(0, writeStr)
	r1.Clear(0)
	fmt.Println("r1 結果 ", r1.Read(0, 9))

}

package main

import (
	"fmt"
	"graidtechHW/classroom"
	"graidtechHW/membermanager"
	"graidtechHW/raid"

	"os"
	"os/signal"
	"syscall"
)

func main() {

	id := 3

	switch id {
	case 1:
		// 第一題 + Bonus 1
		classroom.Model1()
	case 2:
		// 第一題  Bonus 2
		classroom.Model2()
	case 3:
		// 第二題
		membermanager.Enable()
	case 4:
		// 第三題
		raid.Enable()
	default:
		return
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	fmt.Println("main done")
}

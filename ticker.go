package main

import (
	"fmt"
	"time"
)

func main() {
	// 每隔一秒触发一次
	ticker := time.NewTicker(time.Second)

	go func() {
		for range ticker.C {
			fmt.Println("Tick")
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop() // 手动停止定时器
	time.Sleep(2 * time.Second)
	fmt.Println("Program ended")
}

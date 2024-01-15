package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex
var wg sync.WaitGroup

func incrementCounter() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Millisecond)

		// 使用互斥锁确保对计数器的安全访问
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	startTime := time.Now()

	// 设置并发goroutine的数量
	numGoroutines := 1000000

	// 启动goroutine并发执行增加计数器的操作
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementCounter()
	}

	// 等待所有goroutine完成
	wg.Wait()

	// 输出运行时间和计数器最终数值
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("程序运行时间：%v\n", elapsedTime)
	fmt.Printf("计数器最终数值：%d\n", counter)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) { // chan<- int 单向通道，只能向通道发送int数据
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Producer: Sending", i)
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) { // <-chan int 单向通道，只能从通道接收int数据
	defer wg.Done()
	for num := range ch {
		fmt.Println("Consumer: Received", num)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}

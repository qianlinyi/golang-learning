package main

import (
	"fmt"
	"sync"
)

func Print(n int) {
	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	c1 <- true
	go func(n int) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if ok := <-c1; ok {
				fmt.Println("A")
				c2 <- true
			}
		}
	}(n)
	go func(n int) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if ok := <-c2; ok {
				fmt.Println("B")
				c1 <- true
			}
		}
	}(n)
	wg.Wait()
}

func main() {
	Print(10)
}

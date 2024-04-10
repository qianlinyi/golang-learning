package main

import "fmt"

func h() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	// do something with s ...

	// Reset pointer values.
	s[0], s[len(s)-1] = nil, nil
	return s[1:3:3]
}

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(s[1:3:3]) // 起始索引:结束索引:容量
	fmt.Println(h())
}

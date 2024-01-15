package main

import (
	"fmt"
	"sort"
)

type a struct {
	c   string
	cnt int
}

// 全小写的字符串，统计每个字母出现的次数，按出现次数由大到小排序，输出排序结果
func Sort(s string) []a {
	m := make([]a, 0)
	vis := make(map[string]int)
	for _, c := range s {
		_, ok := vis[string(c)]
		if ok {
			vis[string(c)] += 1
		} else {
			vis[string(c)] = 1
		}
	}
	for k, v := range vis {
		m = append(m, a{k, v})
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i].cnt > m[j].cnt
	})
	return m
}

func main() {
	fmt.Println(Sort("abcd122333444455555"))
}

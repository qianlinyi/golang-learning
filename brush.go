package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fuck(url string, wg *sync.WaitGroup, agent string, cnt *int) {
	defer wg.Done()
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	request.Header.Set("User-Agent", agent)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer response.Body.Close()
	*cnt++
	fmt.Println(*cnt)
}

func main() {
	agents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",
		"Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:40.0) Gecko/20100101 Firefox/40.0",
	}
	var wg sync.WaitGroup
	var cnt int = 0
	for true {
		wg.Add(20)
		for i := 0; i < 10; i++ {
			go fuck("https://www.nowcoder.com/discuss/610165644507500544", &wg, agents[i], &cnt)
			go fuck("https://www.nowcoder.com/discuss/612965923007299584", &wg, agents[i], &cnt)
		}
		wg.Wait()
		time.Sleep(7 * time.Second)
	}
}

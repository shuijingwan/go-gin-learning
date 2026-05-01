package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// visited 用于记录已经访问过的 URL，避免重复爬取
	visited := make(map[string]bool)
	var mu sync.Mutex     // 互斥锁，保护 visited 的并发访问
	var wg sync.WaitGroup // WaitGroup 用于等待所有爬取任务完成

	// crawl 是内部递归函数，实际执行爬取逻辑
	var crawl func(url string, depth int)
	crawl = func(url string, depth int) {
		defer wg.Done() // 确保每个 crawl 调用结束时调用 Done，通知 WaitGroup

		if depth <= 0 {
			return
		}

		mu.Lock() // 加锁，保护 visited 的访问
		if visited[url] {
			mu.Unlock() // 已访问过，解锁并返回
			return
		}
		visited[url] = true // 标记当前 URL 已访问
		mu.Unlock()         // 解锁

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			wg.Add(1)            // 为每个新的爬取任务增加计数
			go crawl(u, depth-1) // 启动新的 goroutine 爬取子 URL，深度减 1
		}
	}

	wg.Add(1)            // 为初始 URL 增加计数
	go crawl(url, depth) // 启动爬取初始 URL 的 goroutine
	wg.Wait()            // 等待所有爬取任务完成
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

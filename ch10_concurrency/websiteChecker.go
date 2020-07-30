package ch10_concurrency

// 检查网站相应情况，成功相应的值为true，错误相应的值为false
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	//for _, url := range urls {
	//	results[url] = wc(url)
	//}
	// 协程实现
	for _, url := range urls {
		go func(u string) {
			// 这里需要注意：变量url被重复用于for循环的；每次迭代，；每次都会从urls获取新知。
			// 但是我们每个goroutine都是url变量的引用，它们没有自己的独立副本。所以它们都会
			// 它们都会写入在迭代结束时的url，也就是最后一个url。
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	//time.Sleep(2 * time.Second)
	return results
}

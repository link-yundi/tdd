package ch10_concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}
	got := CheckWebsite(mockWebsiteChecker, websites)
	//if !reflect.DeepEqual(want, got) {
	//	t.Fatalf("Wanted %v, got %v", want, got)
	//}
	assert.Equal(t, want, got)
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		// 模拟网站耗时，每个网站耗时20毫秒
		CheckWebsite(slowStubWebsiteChecker, urls)
	}
}

package ch03_iteration

// character: 单字符
// repeatCount: 重复次数
// 单字符重复组成字符串
func Repeat(character string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return
}


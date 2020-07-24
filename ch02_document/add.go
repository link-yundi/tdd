package ch02_document

// 对n个int数求和
func Add(ops ...int) (sum int) {
	for _, op := range ops {
		sum += op
	}
	return
}
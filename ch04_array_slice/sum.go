package ch04_array_slice

/*
args:
	numbers: array or slice
returns:
	sum: sum of numbers
*/
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

/*
args:
    ops: 多个切片
returns:
    sumSlice: 各个切片的和组成的新的切片
*/
func SumAll(ops ...[]int) (sumSlice []int) {
	sumSlice = make([]int, len(ops))
	for i, numbers := range ops {
		for _, num := range numbers {
			sumSlice[i] += num
		}
	}

	return
}
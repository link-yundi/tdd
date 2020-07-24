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

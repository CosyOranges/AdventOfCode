package maths

func MaxIntSlice(array ...int) int {
	maxNum := array[0]
	for _, x := range array {
		if x > maxNum {
			maxNum = x
		}
	}

	return maxNum
}

/*
Return the sum of all elements in an array of integers
*/
func SumIntSlice(array []int) int {
	sum := 0

	for _, x := range array {
		sum += x
	}

	return sum
}

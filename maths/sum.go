package maths

func SumInt32(nums ...int32) int64 {
	var sum int64
	for _, num := range nums {
		sum += int64(num)
	}
	return sum
}

func SumInt64(nums ...int64) int64 {
	var sum int64
	for _, num := range nums {
		sum += num
	}
	return sum
}

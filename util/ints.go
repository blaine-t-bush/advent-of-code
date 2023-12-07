package util

func IntsSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func IntsProduct(nums []int) int {
	product := 1
	for _, num := range nums {
		product *= num
	}

	return product
}

func IntsPow(base int, exponent int) int {
	if exponent == 0 {
		return 1
	}
	result := base
	for i := 2; i <= exponent; i++ {
		result *= base
	}
	return result
}
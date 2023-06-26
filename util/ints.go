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

package sol

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	res := sum
	for pos := 0; pos < n; pos++ {
		res -= nums[pos]
	}
	return res
}

package sol

func missingNumberXOR(nums []int) int {
	n := len(nums)
	res := 0
	for pos := 0; pos < n; pos++ {
		res ^= ((pos + 1) ^ nums[pos])
	}
	return res
}

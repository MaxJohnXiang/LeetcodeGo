
func rotate(nums []int, k int) {
	// 我已经假定 k >= 0
	n := len(nums)
	if k > n {
		k = k % n
	}
	reverseArray(nums, 0, n-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, n-1)
}

func reverseArray(nums []int, i, j int) []int {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

// package problem0283

func moveZeroes(nums []int) {
	l := 0
	// lddb
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l] = nums[i]
			l++
		}
	}
	//把 l 后面的变成0

	for l < len(nums) {
		nums[l] = 0
		l++
	}
}

// package problem0016
//
// import (
// 	"sort"
// )

func threeSumClosest(nums []int, target int) int {
	// 排序后，可以按规律查找
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[len(nums)-1]
	for i := 0; i < len(nums); i++ {
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[lo] + nums[hi] + nums[i]

			if sum < target {
				lo++
			} else {
				hi--
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

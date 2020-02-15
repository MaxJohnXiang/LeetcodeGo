// package problem0016
//
// import (
// 	"sort"
// )

func threeSumClosest(nums []int, target int) int {
	// 排序后，可以按规律查找
	sort.Ints(nums)

	max := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]

			if sum < target {
				lo++
			} else {
				hi--
			}

			if abs(sum-target) < abs(max-target) {
				max = sum
			}
		}
	}

	return max
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

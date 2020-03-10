// package problem0018
//
// import (
// 	"sort"
// )

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}

	for k := 0; k < len(nums)-3; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		for i := k + 1; i < len(nums)-2; i++ {

			if i > k+1 && nums[i] == nums[i-1] {
				continue
			}

			lo, hi := i+1, len(nums)-1
			for lo < hi {
				sum := nums[i] + nums[lo] + nums[hi] + nums[k]
				if sum == target {
					res = append(res, []int{nums[k], nums[i], nums[lo], nums[hi]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if sum > target {
					hi--
				} else {
					lo++
				}
			}
		}
	}
	return res
}

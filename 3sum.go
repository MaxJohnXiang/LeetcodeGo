// package problem0015

// import "sort"

func threeSum(nums []int) [][]int {
	//双指针法
	//sort array first

	sort.Ints(nums)

	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		lo := i + 1
		hi := len(nums) - 1
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {

			for lo < hi {
				if nums[lo]+nums[hi]+nums[i] == 0 {
					res = append(res, []int{nums[lo], nums[hi], nums[i]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if nums[lo]+nums[hi]+nums[i] > 0 {
					hi--
				} else {
					lo++
				}
			}
		}
	}
	return res
}

// package problem0015
//
// import "sort"

func threeSum(nums []int) [][]int {
	// 排序后，可以按规律查找
	sort.Ints(nums)
	res := [][]int{}
	//需要保证最少有三个数字， 如果少于三个数， 则没有结果
	for i := 0; i < len(nums)-2; i++ {
		lo := i + 1
		hi := len(nums) - 1
		sum := -nums[i]
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			for lo < hi {
				//这个时候lo , hi 都可能越界
				if nums[lo]+nums[hi] == sum {
					res = append(res, []int{nums[i], nums[lo], nums[hi]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if nums[lo]+nums[hi] < sum {
					lo++
				} else {
					hi--
				}
			}
		}
	}
	return res
}

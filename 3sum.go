// package problem0015

// import "sort"

//
// //

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	//排序是前提
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		//处理重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			//大了 k
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				//重点要处理重复的情况.
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			}
		}
	}
	return res
}

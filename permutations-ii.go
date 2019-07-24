import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	var dfs func([]int)
	sort.Ints(nums)
	used := make([]bool, len(nums))
	dfs = func(set []int) {
		if len(set) == len(nums) {
			temp := make([]int, len(set))
			copy(temp, set)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			set = append(set, nums[i])
			used[i] = true
			dfs(set)
			set = set[:len(set)-1]
			used[i] = false
		}
	}
	dfs([]int{})
	return res
}

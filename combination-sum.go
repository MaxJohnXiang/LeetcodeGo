import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	nums := candidates

	res := [][]int{}
	var dfs func([]int, int, int)
	dfs = func(set []int, remain int, start int) {
		if remain < 0 {
			return
		}
		if remain == 0 {
			temp := make([]int, len(set))
			copy(temp, set)
			res = append(res, temp)
			return
		}

		for i := start; i < len(nums); i++ {
			set = append(set, nums[i])
			dfs(set, remain-nums[i], i)
			set = set[:len(set)-1]
		}
	}

	dfs([]int{}, target, 0)
	return res
}

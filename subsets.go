
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	dfs(&res, nums, []int{}, 0)
	return res
}

func dfs(res *[][]int, nums []int, set []int, start int) {
	temp := make([]int, len(set))
	copy(temp, set)
	*res = append(*res, temp)
	for i := start; i < len(nums); i++ {
		set = append(set, nums[i])
		dfs(res, nums, set, i+1)
		set = set[:len(set)-1]
	}
}

func subsets(nums []int) [][]int {

	res := make([][]int, 0)
	var dfs func([]int, int)
	dfs = func(set []int, start int) {
		temp := make([]int, len(set))
		copy(temp, set)
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			temp = append(temp, nums[i])
			dfs(temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}

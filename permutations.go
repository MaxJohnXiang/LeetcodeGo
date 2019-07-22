
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	dfs(&res, []int{}, nums)
	return res
}

func dfs(res *[][]int, set []int, nums []int) {
	if len(set) == len(nums) {
		temp := make([]int, len(set))
		copy(temp, set)
		*res = append(*res, temp)
		return
	} else {
		for i := 0; i < len(nums); i++ {
			if contain(set, nums[i]) {
				continue
			}
			set = append(set, nums[i])
			dfs(res, set, nums)
			set = set[:len(set)-1]
		}
	}
}
func contain(set []int, a int) bool {
	for _, v := range set {
		if v == a {
			return true
		}
	}
	return false
}

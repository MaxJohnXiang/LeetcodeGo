// package problem0022

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(int, int, string)
	dfs = func(left, right int, temp string) {
		if right == n {
			res = append(res, temp)
		}

		if left < n {
			dfs(left+1, right, temp+"(")
		}
		if right < left {
			dfs(left, right+1, temp+")")
		}
	}
	dfs(0, 0, "")
	return res
}

// package problem0022

func generateParenthesis(n int) []string {
	var res []string

	var dfs func(string, int, int)

	dfs = func(str string, open, close int) {
		if len(str) == n*2 {
			res = append(res, str)
		}
		if open < n {
			dfs(str+"(", open+1, close)
		}

		if close < open {
			dfs(str+")", open, close+1)
		}
	}
	dfs("", 0, 0)

	return res
}

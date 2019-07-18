func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//求left node max path sum
	maxSum := root.Val

	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(dfs(root.Left), 0)
		right := max(dfs(root.Right), 0)
		sum := root.Val + left + right
		if sum > maxSum {
			maxSum = sum
		}
		return max(left, right) + root.Val
	}
	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

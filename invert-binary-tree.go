// package problem0226
//
// import (
// 	"github.com/aQuaYi/LeetCode-in-Go/kit"
// )
//
// type TreeNode = kit.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var Dfs func(*TreeNode) *TreeNode
	Dfs = func(t *TreeNode) *TreeNode {
		if t == nil {
			return nil
		}
		return &TreeNode{
			Val:   t.Val,
			Left:  Dfs(t.Right),
			Right: Dfs(t.Left),
		}
	}

	return Dfs(root)
}

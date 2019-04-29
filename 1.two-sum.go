/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (45.82%)
 * Total Accepted:    347.3K
 * Total Submissions: 757.9K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * <p>给定一个整数数组 <code>nums</code>&nbsp;和一个目标值
 * <code>target</code>，请你在该数组中找出和为目标值的那&nbsp;<strong>两个</strong>&nbsp;整数，并返回他们的数组下标。</p>
 *
 * <p>你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。</p>
 *
 * <p><strong>示例:</strong></p>
 *
 * <pre>给定 nums = [2, 7, 11, 15], target = 9
 *
 * 因为 nums[<strong>0</strong>] + nums[<strong>1</strong>] = 2 + 7 = 9
 * 所以返回 [<strong>0, 1</strong>]
 * </pre>
 *
 */
func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, v := range nums {
		if in, ok := index[target-v]; ok {
			return []int{in, i}
		}
		index[v] = i
	}
	return []int{}
}

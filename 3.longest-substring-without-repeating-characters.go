/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (28.25%)
 * Total Accepted:    891.9K
 * Total Submissions: 3.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 *
 *
 */
func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	//首先假设所有的字符都没有看过
	for i := range location {
		location[i] = -1
	}
	//left 是没有重复的左边界
	//
	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		//location[s[i]] 如果没有出现应该都是-1
		//出现后则是下标
		if location[s[i]] >= left {
			//出现了重复， 计算长度
			left = location[s[i]] + 1
		} else if i-left+1 > maxLen {
			maxLen = i - left + 1
		}
		//记录每一个字符出现的位置， 如果有重复的字母， 则会覆盖，记录最后一个字母的位置
		location[s[i]] = i
	}
	return maxLen
}

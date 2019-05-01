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
	for i := range location {
		location[i] = -1
	}
	//left 代表不重复的左边界
	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		//重复出现了
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i-left+1 > maxLen {
			maxLen = i - left + 1
		}
		location[s[i]] = i
	}
	return maxLen
}

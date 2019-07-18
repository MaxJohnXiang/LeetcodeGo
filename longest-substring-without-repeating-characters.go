
func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	left := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		//如何判断s[left:i]是否重复出现过。
		//两种情况一种是从未出现过 location[s[i]] = -1 如果大于0 ， 说明之前出现过
		//但是还是要考虑， 窗口是滑动的， 这个字符串的开始不是0，
		//而是最后一个不重复的开始
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i-left+1 > maxLen {
			maxLen = i - left + 1
		}
		//记录每个字符出现的位置
		location[s[i]] = i
	}
	return maxLen
}

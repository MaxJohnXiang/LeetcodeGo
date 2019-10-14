// package problem0003

func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	start := 0
	maxLen := 0

	location := make([]int, 128)
	for i, _ := range location {
		location[i] = -1
	}
	for i, v := range s {
		//start 记录上次不重复的地方。
		if location[v] >= start {
			start = location[v]
		} else if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		location[v] = i + 1
	}
	return maxLen
}

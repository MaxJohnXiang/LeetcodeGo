// package problem0003

func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := make([]int, 128)

	//init location
	for i := range location {
		location[i] = -1
	}

	start := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		//start 是重复的地方
		if location[s[i]] >= start {
			start = location[s[i]] + 1
		}
		if i-start+1 >= maxLen {
			maxLen = i - start + 1
		}
		location[s[i]] = i
	}

	return maxLen
}

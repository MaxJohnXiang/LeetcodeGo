// package problem0003

func lengthOfLongestSubstring(s string) int {
	location := make([]int, 128)
	for i := range location {
		location[i] = -1
	}

	start := 0
	maxLen := 0
	//abbc
	for i := 0; i < len(s); i++ {
		//滑动窗口
		//start 记录没有重复的的开始区间， 一旦出现重复， 后移动格子一
		if location[s[i]] >= start {
			start = location[s[i]] + 1
		} else if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		location[s[i]] = i
	}
	return maxLen
}

// package problem0003

func lengthOfLongestSubstring(s string) int {
	location := [128]int{}

	//init array
	for i := range location {
		location[i] = -1
	}

	start := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		//判断这个出现过了没, 如果出现了. start update
		//这个字符上次的出现是否大于 start
		if location[s[i]] >= start {
			start = location[s[i]] + 1
		} else if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		//previouse element location
		location[s[i]] = i
	}

	return maxLen
}

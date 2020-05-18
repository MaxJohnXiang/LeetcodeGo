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
		w := s[i]
		if location[w] >= start {
			start = location[w]  + 1
		} else if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		location[w] = i
	}
	return maxLen
}

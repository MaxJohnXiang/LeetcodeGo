func lengthOfLongestSubstring(s string) int {

	location := [256]int{}
	maxLen := 0
	for i := range location {
		location[i] = -1
	}
	for i, j := 0, 0; j < len(s); j++ {
		//怎么判断重复的。 location 数组的初始值都是0， 如果location[j] >0 ,说明这个重复了， 重新开始计数
		//如果重复了， 重新求值， 如果不重复求长度
		if location[s[j]] >= i {
			i = location[s[j]] + 1
		} else if j-i+1 > maxLen {
			maxLen = j - i + 1
		}
		location[s[j]] = j
	}
	return maxLen
}

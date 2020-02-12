// package problem0005

func longestPalindrome(s string) string {
	//动态规划
	if len(s) < 2 {
		//肯定是回文
		return s
	}
	var max string
	for i := 0; i < len(s); i++ {
		max = isPalindrome(s, i, i, max)
		max = isPalindrome(s, i, i+1, max)
	}

	return max
}

//找到最长的回文
func isPalindrome(s string, i int, j int, max string) string {
	var sub string
	for (i >= 0 && j < len(s)) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(sub) > len(max) {
		max = sub
	}

	return max
}

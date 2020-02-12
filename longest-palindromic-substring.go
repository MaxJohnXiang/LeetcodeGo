// package problem0005

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, 0)

	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}

	maxLen := 0
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 >= maxLen {
				maxLen = j - i + 1
				res = s[i : j+1]
			}
		}
	}
	return res
}

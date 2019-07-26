// package problem0005

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, 0)
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}
	res := ""
	maxLen := 0
	for i := n; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 >= maxLen {
				maxLen = j - i + 1
				res = s[i : j+1]
			}
		}
	}
	return res
}

// package problem0005

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, 0)
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}
	res := ""
	for i := n; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

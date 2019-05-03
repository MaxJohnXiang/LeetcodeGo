// package problem0005

// package main

// func main() {
// 	longestPalindrome("cbbd")
// }
func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	dp := make([][]bool, 0)
	for i := 0; i <= len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}
	left, right := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 2 || dp[i+1][j-1])
			if dp[i][j] && j-i > right-left {
				left = i
				right = j
			}
		}
	}
	return s[left : right+1]
}

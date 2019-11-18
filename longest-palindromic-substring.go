// package problem0005

func longestPalindrome(s string) string {
	//动态规划
	if len(s) < 2 {
		//肯定是回文
		return s
	}

	res := ""
	n := len(s)

	dp := make([][]bool, n)

	// for i := 0; i < n; i++ {
	// 	dp = append(dp, make([]bool, n))
	// }
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	//声明数组
	maxLen := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			//有两种条件是回文， 第一种字符串长度为1 是回文
			//2. 字符串长度为2 ， 两个字母都相等就为回文
			//3, 字符串长度大于2, 最外层字母相等，并且里面是回文
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}

			if dp[i][j] && j-i+1 >= maxLen {
				res = s[i : j+1]
				maxLen = j - i + 1
			}
		}
	}

	return res
}

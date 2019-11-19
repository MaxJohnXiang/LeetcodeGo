// package problem0010

func isMatch(s, p string) bool {
	sSize := len(s)
	pSize := len(p)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	/* dp[i][j] 代表了 s[:i] 能否与 p[:j] 匹配 */

	dp[0][0] = true
	/**
	 * 根据题目的设定， "" 可以与 "a*b*c*" 相匹配
	 * 所以，需要把相应的 dp 设置成 true
	 *  j第二个开始, 每次步长是2
	 */
	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 1; i < sSize+1; i++ {
		for j := 1; j < pSize+1; j++ {
			parten := p[j-1]
			str := s[i-1]

			if parten == '.' || parten == str {
				dp[i][j] = dp[i-1][j-1]
			} else if parten == '*' {
				prevParten := p[j-2]
				if prevParten != '.' && prevParten != str {
					//因为前一个字符和当前text不匹配，而
					//* 则是通过匹配前一个字符一次，或者多次来实现匹配， 所以，
					//只能是匹配0次
					dp[i][j] = dp[i][j-2]
				} else {
					//这种情况p[j-1] == s[i], 可以通过匹配一次，或者多次的情况解决问题
					//parten == * 要分好几种情况
					/**
					* 1. 匹配0次
					 */
					matchZero := dp[i][j-2]
					//匹配1次
					matchOnce := dp[i][j-1]
					//匹配多次
					matchMul := dp[i-1][j]
					dp[i][j] = matchZero || matchOnce || matchMul
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[sSize][pSize]
}

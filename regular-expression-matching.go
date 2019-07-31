// package problem0010

// 1, If p.charAt(j) == s.charAt(i) :  dp[i][j] = dp[i-1][j-1];
//    如果p[j]== s[i] , 这个状态依赖于前一个状态
// 2, If p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1];

// 3, If p.charAt(j) == '*':
//    here are two sub conditions:
//                1   if p.charAt(j-1) != s.charAt(i) : dp[i][j] = dp[i][j-2]  //in this case, a* only counts as empty
//                2   if p.charAt(i-1) == s.charAt(i) or p.charAt(i-1) == '.':
//                               dp[i][j] = dp[i-1][j]    //in this case, a* counts as multiple a
//                            or dp[i][j] = dp[i][j-1]   // in this case, a* counts as single a
//                            or dp[i][j] = dp[i][j-2]   // in this case, a* counts as empty

func isMatch(s, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	//初始化
	//都是空字符串的情况下
	dp[0][0] = true

	// s = ""  p = "b*a*"

	for i := 1; i < len(dp[0]); i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

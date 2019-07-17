
func minDistance(word1, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		//init array
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			//init dp
			//first column is zero
			if i == 0 || j == 0 {
				dp[i][j] = i | j
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				}
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

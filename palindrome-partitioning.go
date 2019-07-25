
func partition(s string) [][]string {
	res := [][]string{}
	var dfs func([]string, int)
	//拆解问题， 先找到所有的subset ，然后把符合条件的subset 加入到结果中
	dfs = func(set []string, start int) {
		if start >= len(s) {
			temp := make([]string, len(set))
			copy(temp, set)
			res = append(res, temp)
		}
		for i := start; i < len(s); i++ {
			if isPalindrome(s[start : i+1]) {
				set = append(set, s[start:i+1])
				dfs(set, i+1)
				set = set[:len(set)-1]
			}
		}
	}
	dfs([]string{}, 0)
	return res
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	l, h := 0, len(s)-1

	for l < h {
		if s[l] != s[h] {
			return false
		}
		l++
		h--
	}
	return true
}

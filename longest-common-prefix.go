// package problem0014

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	s := shortest(strs)
	i, j := 0, len(s)
	for _, v := range strs {
		//这个最小的字符串， 是否是longestCommonPrefix
		for v[i:j] != s {
			j--
			s = s[i:j]
		}
	}
	return s
}

func shortest(strs []string) string {
	s := strs[0]

	for _, v := range strs {
		if len(v) < len(s) {
			s = v
		}
	}
	return s
}

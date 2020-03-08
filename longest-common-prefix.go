// package problem0014

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := shortest(strs)
	i, j := 0, len(prefix)

	for _, v := range strs {
		for v[i:j] != prefix {
			j--
			prefix = prefix[i:j]
		}
	}

	return prefix
}

func shortest(strs []string) string {
	min := strs[0]

	for _, v := range strs {
		if len(v) < len(min) {
			min = v
		}
	}

	return min
}

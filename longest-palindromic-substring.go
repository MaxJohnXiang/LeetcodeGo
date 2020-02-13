// package problem0005

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var max string
	for i := 0; i < len(s); i++ {
		max = isPlidrome(s, i, i, max)
		max = isPlidrome(s, i, i+1, max)
	}
	return max
}

func isPlidrome(s string, i, j int, max string) string {

	var sub string
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}

	if len(sub) > len(max) {
		max = sub
	}

	return max
}

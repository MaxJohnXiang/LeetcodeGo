// package problem0005

func longestPalindrome(s string) string {
	// froi
	if len(s) < 2 {
		return s
	}
	max := ""
	for i := 0; i < len(s)-1; i++ {
		max = isPalindrome(s, i, i, max)
		max = isPalindrome(s, i, i+1, max)
	}

	return max
}

// judge that string is palindrome or not
func isPalindrome(s string, i, j int, max string) string {
	tmp := ""

	for i >= 0 && j < len(s) && s[i] == s[j] {
		tmp = s[i : j+1]
		if len(tmp) > len(max) {
			max = tmp
		}
		i--
		j++
	}

	return max
}

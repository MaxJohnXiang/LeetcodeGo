
func minWindow(s string, t string) string {
	have := make(map[rune]int)
	need := make(map[rune]int)

	for _, v := range t {
		need[v]++
	}

	min := len(s) + 1
	res := ""
	for l, r, count := 0, 0, 0; r < len(s); r++ {
		ch := rune(s[r])
		have[ch]++
		if _, ok := need[ch]; ok && have[ch] == need[ch] {
			count++
		}

		for count == len(need) && l <= r {
			if r-l+1 < min {
				min = r - l + 1
				res = s[l : r+1]
			}

			ch = rune(s[l])
			have[ch]--
			l++
			if have[ch] < need[ch] {
				count--
			}
		}
	}
	return res
}

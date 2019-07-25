// package problem0076

func minWindow(s string, t string) string {
	have := make(map[rune]int)
	need := make(map[rune]int)

	for _, v := range t {
		need[rune(v)]++
	}
	count := 0
	min := len(s) + 1
	res := ""

	for l, r := 0, 0; r < len(s); r++ {
		ch := rune(s[r])
		have[ch]++
		if have[ch] == need[ch] {
			count++
		}
		for count >= len(need) && l <= r {
			ch := rune(s[l])
			if r-l+1 < min {
				min = r - l + 1
				res = s[l : r+1]
			}
			have[ch]--
			l++
			if have[ch] < need[ch] {
				count--
			}
		}
	}
	return res
}

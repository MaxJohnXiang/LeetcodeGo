// package problem0076

func minWindow(s string, t string) string {

	have := [128]int{}
	need := [128]int{}

	//init need
	for _, v := range t {
		need[v]++
	}

	minLen := len(s) + 1

	matchCount := 0

	slow, fast := 0, 0

	res := ""
	for ; fast < len(s); fast++ {
		//这个字符是否在 need 中出现过了, 如果出现过了,  matchCount++
		//如果这个字符是
		if have[s[fast]] < need[s[fast]] {
			matchCount++
		}
		have[s[fast]]++

		//shorter the window
		for slow <= fast && have[s[slow]] > need[s[slow]] {
			have[s[slow]]--
			slow++
		}

		//.ji
		if matchCount == len(t) && fast-slow+1 <= minLen {
			minLen = fast - slow + 1
			res = s[slow : fast+1]
		}
	}

	return res
}
func minWindow1(s string, t string) string {
	have := [128]int{}
	need := [128]int{}
	for i := range t {
		need[t[i]]++
	}

	size, total := len(s), len(t)

	min := size + 1
	res := ""

	// s[i:j+1] 就是 window
	// count 用于统计已有的 t 中字母的数量。
	// count == total 表示已经收集完需要的全部字母
	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			// 出现了 window 中缺失的字母
			count++
		}
		have[s[j]]++

		// 保证 window 不丢失所需字母的前提下
		// 让 i 尽可能的大
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}

		width := j - i + 1
		if count == total && min > width {
			min = width
			res = s[i : j+1]
		}
	}

	return res
}

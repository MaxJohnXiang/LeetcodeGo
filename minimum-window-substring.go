// package problem0076

func minWindow(s string, t string) string {
	have := [128]int{}
	need := [128]int{}

	//3、此时，我们停止增加 right，转而不断增加 left 指针缩小窗口 [left, right)，直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。同时，每次增加 left，我们都要更新一轮结果。
	//4、重复第 2 和第 3 步，直到 right 到达字符串 S 的尽头。

	//init need
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	res := ""
	min := len(s) + 1

	matchCount := 0
	//1、我们在字符串 S 中使用双指针中的左右指针技巧，初始化 left = right = 0，把索引左闭右开区间 [left, right) 称为一个「窗口」。
	for ; right < len(s); right++ {
		//2、我们先不断地增加 right 指针扩大窗口 [left, right)，直到窗口中的字符串符合要求（包含了 T 中的所有字符）。
		//什么时字符会++
		have[s[right]]++

		if have[s[right]] <= need[s[right]] {
			matchCount++
		}

		for left <= right && have[s[left]] > need[s[left]] {
			have[s[left]]--
			left++
		}

		if matchCount == len(t) && right-left+1 <= min {
			res = s[left : right+1]
			min = right - left + 1
		}
	}

	return res
}

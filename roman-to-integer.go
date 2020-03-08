// package problem0013

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	prev := m[s[len(s)-1]]
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := m[s[i]]

		if cur < prev {
			res -= cur
		} else {
			res += cur
		}

		//如果比上一个小
		prev = cur
	}

	return res
}

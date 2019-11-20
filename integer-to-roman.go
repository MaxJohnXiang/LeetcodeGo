// package problem0012

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""

	for i, v := range values {
		for num >= v {
			num -= v
			res += strs[i]
		}
	}

	return res
}

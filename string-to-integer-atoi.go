

// import (
// 	"math"
// 	"strings"
// )

// package problem0008

func myAtoi(s string) int {
	return convert(clean(s))
}

func clean(s string) (sign int, abs string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	default:
		abs = ""
		return
	}
	for i, b := range abs {
		if b > '9' || b < '0' {
			abs = abs[:i]
			break
		}
	}
	return
}

func convert(sign int, s string) int {
	ans := 0
	for _, b := range s {
		ans = ans*10 + int(b-'0')
		if sign == -1 && ans*sign < math.MinInt32 {
			return math.MinInt32
		}
		if sign == 1 && ans*sign > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return ans * sign
}

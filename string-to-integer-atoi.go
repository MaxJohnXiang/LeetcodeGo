// package problem0008
//
// import "strings"
// import "math"

//
// Example 1:
//
//
// Input: "42"
// Output: 42
//
//
// Example 2:
//
//
// Input: "   -42"
// Output: -42
// Explanation: The first non-whitespace character is '-', which is the minus sign.
//              Then take as many numerical digits as possible, which gets 42.
//
//
// Example 3:
//
//
// Input: "4193 with words"
// Output: 4193
// Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
//
func myAtoi(s string) int {

	//If no valid conversion could be performed, a zero value is returned.

	//If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

	//Only the space character ' ' is considered as whitespace character.
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}
	first := s[0]
	sign := 1

	if first == '+' {
		sign = 1
		s = s[1:]
	} else if first == '-' {
		sign = -1
		s = s[1:]
	} else if first < '0' || first > '9' {
		return 0
	}

	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		temp := int(s[i] - '0')
		if res > math.MaxInt32/10 || res == math.MaxInt32/10 && temp > math.MaxInt32%10 {
			if sign == 1 {
				return math.MaxInt32
			}
			if sign == -1 {
				return math.MinInt32
			}
		}
		res = res*10 + temp
	}
	return res * sign
}

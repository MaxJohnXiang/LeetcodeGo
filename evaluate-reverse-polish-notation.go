// package problem0150
//
// import (
// 	"strconv"
// )

func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens))

	for _, v := range tokens {
		if v == "+" ||
			v == "-" ||
			v == "/" ||
			v == "*" {
			b, a := nums[len(nums)-1], nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, compute(v, a, b))
		} else {
			temp, _ := strconv.Atoi(v)
			nums = append(nums, temp)
		}
	}
	return nums[0]
}

func compute(token string, a, b int) int {
	switch token {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	default:
		return a * b
	}
}

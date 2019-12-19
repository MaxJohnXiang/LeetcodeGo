// package problem0029
//
// import (
// 	"math"
// )
//
// //dividend : 被除数
// //divisor : 除数
// //dividend / divisor
func divide(dividend, divisor int) int {
	// 防止有人把0当做除数

	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}

	sign := 1

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	res := div(dividend, divisor)

	if sign > 0 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		return res
	}
	return -res
}

func div(dividend, divisor int) int {
	if dividend < divisor {
		return 0
	}

	tb := divisor
	count := 1

	for (tb + tb) < dividend {
		count += count
		tb += tb
	}

	return count + div(dividend-tb, divisor)
}

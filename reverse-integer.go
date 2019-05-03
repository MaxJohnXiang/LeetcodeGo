import "math"

//import "math"

// import "math"

// package problem0007

// import "math"

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = sign * -1
		x = x * sign
	}
	ans := 0
	for x > 0 {
		temp := x % 10
		ans = ans*10 + temp
		x /= 10
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}
	return ans * sign
}

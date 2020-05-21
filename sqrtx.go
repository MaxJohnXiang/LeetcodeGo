// package problem0069

func mySqrt(x int) int {
	// res := x
	// // 牛顿法求平方根
	// for res*res > x {
	// 	res = (res + x/res) / 2
	// }
	// return res

	l := 0
	r := x
	ans := 0
	for l <= r {
		mid := l + (r-l)/2

		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

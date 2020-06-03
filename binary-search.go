// package problem0704

func search(a []int, target int) int {
	if len(a) < 1 {
		return -1
	}

	l, r := 0, len(a)-1

	for l <= r {
		m := l + (r-l)/2
		if a[m] == target {
			return m
		} else if a[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

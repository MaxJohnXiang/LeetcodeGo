// package problem0026

func removeDuplicates(a []int) int {
	if len(a) < 1 {
		return 0
	}
	left, right, size := 0, 1, len(a)
	for ; right < size; right++ {
		if a[right] == a[left] {
			continue
		}
		left++
		a[left], a[right] = a[right], a[left]
	}

	return left + 1
}

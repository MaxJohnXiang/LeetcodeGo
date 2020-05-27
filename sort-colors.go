// package problem0075

// 借鉴三路快排中的划分思路
func sortColors(a []int) {

	i, j, k := 0, 0, len(a)-1
	for j <= k {
		if a[j] == 0 {
			a[i], a[j] = a[j], a[i]
			i++
			j++
		} else if a[j] == 1 {
			j++
		} else {
			a[j], a[k] = a[k], a[j]
			k--
		}

	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

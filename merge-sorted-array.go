// package problem0088

// 本题的要求是，把nums1的前m项和nums2的前n项合并，放入nums1中。
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 深度复制 nums1
	temp := make([]int, m)
	copy(temp, nums1)

	i, j := 0, 0
	for k := 0; k < len(nums1); k++ {
		if i >= m {
			nums1[k] = nums2[j]
			j++
			continue
		}

		if j >= n {
			nums1[k] = temp[i]
			i++
			continue
		}

		if temp[i] < nums2[j] {
			nums1[k] = temp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}
}

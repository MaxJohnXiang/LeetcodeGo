// package problem0088

// 本题的要求是，把nums1的前m项和nums2的前n项合并，放入nums1中。
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 深度复制 nums1

	k := m + n - 1
	i := m - 1
	j := n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}

	for j >= 0 && len(nums1) > 0 {

		nums1[k] = nums2[j]
		j--
		k--
	}
}

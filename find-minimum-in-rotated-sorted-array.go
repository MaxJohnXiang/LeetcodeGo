// package problem0153

func findMin(a []int) int {
	nums := a
	left, right, mid := 0, len(a)-1, 0
	for left <= right {
		if nums[left] <= nums[right] {
			return nums[left]
		}
		mid = left + (right-left)/2
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}

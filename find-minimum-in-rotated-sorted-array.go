// package problem0153

func findMin(a []int) int {
	nums := a
	left, right, mid := 0, len(a)-1, 0

	for left <= right {
		mid = left + (right-left)/2
		//[mid, right] rotated
		if nums[left] <= nums[right] {
			return nums[left]
		}

		if nums[left] <= nums[mid] {
			//[left, mid ] 连续
			//肯定不在这里
			left = mid + 1
		} else {
			//在[left,mid]有最小值, 找到这个最小的递增开始.
			right = mid
		}
	}

	return -1
}

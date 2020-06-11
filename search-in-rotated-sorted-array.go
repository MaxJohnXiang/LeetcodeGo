// package problem0033

func search(nums []int, target int) int {
	rotated := indexOfMin(nums) /* 数组旋转了的距离 */
	size := len(nums)
	left, right := 0, size-1

	for left <= right {
		mid := left + (right-left)/2
		/* nums 是 rotated，所以需要使用 rotatedmid 来获取 mid 的值 */
		rotatedmid := (rotated + mid) % size
		switch {
		case nums[rotatedmid] < target:
			left = mid + 1
		case target < nums[rotatedmid]:
			right = mid - 1
		default:
			return rotatedmid
		}
	}

	return -1
}

/* nums 是被旋转了的递增数组 */
func indexOfMin(nums []int) int {
	size := len(nums)
	left, right := 0, size-1
	for left < right {
		mid := (left + right) / 2
		//说明 mid 左边是 roated
		if nums[right] < nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

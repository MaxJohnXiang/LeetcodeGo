// package problem0027

func removeElement(nums []int, val int) int {
	// j指向最后一个不为val的位置
	// i指向第一个值为val的位置

	i := 0
	j := len(nums) - 1
	for {
		for i < len(nums) && nums[i] != val {
			i++
		}

		for j >= i && nums[j] == val {
			j--
		}

		if i > j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i
}

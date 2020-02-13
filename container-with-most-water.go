// package problem0011

func maxArea(height []int) int {
	//这是一个双指针， 如何移动， 能够遍历到最大面积的区域
	//从中心向两边扩张
	//从两边向中心靠
	i := 0
	j := len(height) - 1

	max := 0
	for i < j {
		h := min(height[i], height[j])
		area := h * (j - i)
		if area > max {
			max = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

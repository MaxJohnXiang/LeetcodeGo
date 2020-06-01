// package problem0011

func maxArea(height []int) int {
	//这是一个双指针， 如何移动， 能够遍历到最大面积的区域
	//从中心向两边扩张
	//从两边向中心靠

	max := 0
	i, j := 0, len(height)-1

	for i < j {
		area := getArea(i, j, height)
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

func getArea(i, j int, height []int) int {
	return (j - i) * min(height[i], height[j])
}

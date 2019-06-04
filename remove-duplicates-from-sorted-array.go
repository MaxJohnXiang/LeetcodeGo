
func removeDuplicates(a []int) int {
	//这个数组是有序的
	//有两个指针， 一个记录正在遍历的指针， 一个记录不重复的指针
	left, right := 0, 1
	for right < len(a) {
		//重复了
		if a[right] != a[left] {
			left++
		}
		a[left] = a[right]
		right++
	}
	return left + 1
}

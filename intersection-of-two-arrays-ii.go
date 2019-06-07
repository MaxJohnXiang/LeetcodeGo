
func intersect(a1, a2 []int) []int {
	s1 := make(map[int]int)
	for _, v := range a1 {
		s1[v]++
	}
	res := make([]int, 0)
	for _, v := range a2 {
		if s1[v] > 0 {
			res = append(res, v)
			s1[v]--
		}
	}
	return res
}

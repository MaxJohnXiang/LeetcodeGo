// package problem0084

func largestRectangleArea(heights []int) int {
	res := 0
	s := NewStack()
	s.Push(-1)
	n := len(heights)

	for i := 0; i < len(heights); i++ {
		for s.Peek() != -1 && heights[i] <= heights[s.Peek()] {
			//计算最大值
			curr := s.Pop()
			res = max(res, (i-s.Peek()-1)*heights[curr])
		}
		s.Push(i)
	}

	for s.Len() > 1 {
		curr := s.Pop()
		res = max(res, heights[curr]*(n-s.Peek()-1))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack struct {
	nums []int
}

// NewStack 返回 *kit.Stack
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// Push 把 n 放入 栈
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}
func (s *Stack) Peek() int {
	res := s.nums[len(s.nums)-1]
	return res
}

// Pop 从 s 中取出最后放入 栈 的值
func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

// Len 返回 s 的长度
func (s *Stack) Len() int {
	return len(s.nums)
}

// IsEmpty 反馈 s 是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

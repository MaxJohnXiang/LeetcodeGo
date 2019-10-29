// package problem0150
//
// import (
// 	"strconv"
// )

type Stack struct {
	list []int
}

func (stack *Stack) Push(item int) {
	stack.list = append(stack.list, item)
}

func (stack *Stack) Pop() int {
	pop := stack.list[len(stack.list)-1]
	stack.list = stack.list[:len(stack.list)-1]
	return pop
}

func (stack *Stack) Peek() int {
	return stack.list[len(stack.list)-1]
}

func evalRPN(tokens []string) int {
	// 用于存放数字的栈
	stack := &Stack{}
	for _, s := range tokens {
		if s == "+" ||
			s == "-" ||
			s == "*" ||
			s == "/" {
			//如果是表达式//out stack
			pop1 := stack.Pop()
			pop2 := stack.Pop()
			stack.Push(compute(pop2, pop1, s))
		} else {
			temp, _ := strconv.Atoi(s)
			stack.Push(temp)
		}
	}

	return stack.Pop()
}

// 计算
func compute(a, b int, opt string) int {
	switch opt {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}

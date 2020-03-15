// package problem0023
//
// // ListNode 是链接节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeKLists(list []*ListNode) *ListNode {

	return partion(list, 0, len(list)-1)
}

func partion(list []*ListNode, s, e int) *ListNode {
	if s == e {
		return list[s]
	}

	if s < e {
		q := (s + e) / 2
		l1 := partion(list, s, q)
		l2 := partion(list, q+1, e)
		return merge(l1, l2)
	} else {
		return nil
	}
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	l3 := &ListNode{}
	newHead := l3
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l3.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			l3.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return newHead.Next
}

// package problem0021

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 是链接节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//corner case
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
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
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

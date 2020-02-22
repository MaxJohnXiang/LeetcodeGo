// package problem0002
//
// import (
// 	"github.com/aQuaYi/LeetCode-in-Go/kit"
// )

// ListNode defines for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// //You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
// //
// //You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// //
// //Example:
//
// //Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// //Output: 7 -> 0 -> 8
// //Explanation: 342 + 465 = 807.

// type ListNode = kit.ListNode
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	l3 := head

	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		l3.Next = &ListNode{Val: sum % 10}
		l3 = l3.Next
		carry = sum / 10
	}
	return head.Next
}

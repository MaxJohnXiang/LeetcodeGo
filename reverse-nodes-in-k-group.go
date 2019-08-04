// package problem0025
//
// import (
// 	"github.com/aQuaYi/LeetCode-in-Go/kit"
// )
//
// // ListNode defines for singly-linked list.
// type ListNode = kit.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {

	dummy := &ListNode{}
	start := dummy
	dummy.Next = head
	for {
		prev := start
		cur, next := prev, prev
		start = prev.Next
		for i := 0; i < k && next != nil; i++ {
			next = next.Next
		}
		if next == nil {
			break
		}
		for i := 0; i < k-1; i++ {
			cur = prev.Next
			prev.Next = cur.Next
			cur.Next = next.Next
			next.Next = cur
		}
	}
	return dummy.Next
}

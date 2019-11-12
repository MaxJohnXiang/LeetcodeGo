// package problem0019
//
// /*
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//
// // ListNode 是节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	start := &ListNode{}

	slow, fast := start, start

	start.Next = head

	//fast move n fast

	//n always valid
	for i := 1; i <= n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return start.Next
}

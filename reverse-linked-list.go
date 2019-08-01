// package problem0206

func reverseList(head *ListNode) *ListNode {
	// prev 是所有已经逆转的节点的head
	var prev *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

// // ListNode 是链接节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

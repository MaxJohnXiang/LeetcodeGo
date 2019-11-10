// package problem0206

func reverseList(head *ListNode) *ListNode {
	//  a->b->c->d
	//记录这个空节点
	var prevNode *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prevNode
		//游标转换
		prevNode = cur
		cur = temp
	}
	return prevNode
}

// // ListNode 是链接节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

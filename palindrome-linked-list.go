// package problem0234
//
// import (
// 	"github.com/aQuaYi/LeetCode-in-Go/kit"
// )
//
// type ListNode = kit.ListNode

func isPalindrome(head *ListNode) bool {
	middle := findMiddle(head)
	an := reverseList(middle)

	for an != nil {
		if an.Val != head.Val {
			return false
		}
		an = an.Next
		head = head.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
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

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

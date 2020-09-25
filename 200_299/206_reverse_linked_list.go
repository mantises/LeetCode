package _200_299

/*
Reverse a singly linked list.

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively.
Could you implement both?
*/

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// double pointers
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p, q := head, head.Next
	head.Next = nil
	for q != nil {
		tmp := q.Next
		q.Next = p
		p = q
		q = tmp
	}
	return p
}

// recursive
func reverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

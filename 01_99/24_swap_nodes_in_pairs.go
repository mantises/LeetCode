package _01_99

/*
Given a linked list, swap every two adjacent nodes and return its head.

Example 1:
Input: head = [1,2,3,4]
Output: [2,1,4,3]

Example 2:
Input: head = []
Output: []

Example 3:
Input: head = [1]
Output: [1]


Constraints:
The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{Next: head}
	p, q, r := newHead, head, head.Next
	for q != nil && r != nil {
		p.Next = r
		q.Next = r.Next
		r.Next = q
		p = q
		q = p.Next
		if q != nil && q.Next != nil {
			r = q.Next
		} else {
			r = nil
		}
	}
	return newHead.Next
}

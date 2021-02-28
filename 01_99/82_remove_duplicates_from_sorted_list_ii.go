package _01_99

/*
Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list. Return the linked list sorted
as well.


Example 1:
Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Example 2:
Input: head = [1,1,1,2,3]
Output: [2,3]


Constraints:
The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{Next: head}
	p, q := newHead, head
	for q != nil && q.Next != nil {
		r := q
		for r.Next != nil && q.Val == r.Next.Val {
			r = r.Next
		}
		if q != r {
			p.Next = r.Next
			q = p.Next
		} else if q.Next != nil {
			p, q = q, q.Next
		}
	}
	return newHead.Next
}

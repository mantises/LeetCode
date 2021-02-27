package _01_99

/*
Given the head of a sorted linked list, delete all duplicates such that each element
appears only once. Return the linked list sorted as well.


Example 1:
Input: head = [1,1,2]
Output: [1,2]

Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]


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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{Next: head}
	cur := newHead.Next
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			p := cur.Next
			cur.Next, p.Next = p.Next, nil
		} else {
			cur = cur.Next
		}
	}
	return newHead.Next
}

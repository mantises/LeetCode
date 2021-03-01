package _300_399

/*
Given a non-negative integer represented as a linked list of digits, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list.

Example 1:
Input: head = [1,2,3]
Output: [1,2,4]

Example 2:
Input: head = [0]
Output: [1]


Constraints:
The number of nodes in the linked list is in the range [1, 100].
0 <= Node.val <= 9
The number represented by the linked list does not contain leading zeros except for the zero
itself.

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
	slow := &ListNode{Next: head, Val: 0}
	fast := head

	for fast != nil {
		if fast.Val != 9 {
			slow = fast
		}
		fast = fast.Next
	}

	slow.Val += 1
	cur := slow.Next
	for cur != nil {
		cur.Val = 0
		cur = cur.Next
	}
	if slow.Next == head {
		head = slow
	}
	return head
}

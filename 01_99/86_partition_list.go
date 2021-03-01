package _01_99

/*
Given the head of a linked list and a value x, partition it such that all nodes less than x come
before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.


Example 1:
Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]

Example 2:
Input: head = [2,1], x = 2
Output: [1,2]


Constraints:
The number of nodes in the list is in the range [0, 200].
-100 <= Node.val <= 100
-200 <= x <= 200

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{Next: head}
	before := newHead
	var after *ListNode
	cur, head := head, nil
	for cur != nil {
		if cur.Val < x {
			before.Next = cur
			cur, cur.Next = cur.Next, nil
			before = before.Next
		} else {
			if after == nil {
				head = cur
				after = head
			} else {
				after.Next = cur
				after = after.Next
			}
			cur, cur.Next = cur.Next, nil
		}
	}
	before.Next = head
	return newHead.Next
}

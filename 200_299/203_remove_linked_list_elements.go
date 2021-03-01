package _200_299

/*
Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Next: head}
	pre := newHead
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next, pre.Next.Next = pre.Next.Next, nil
		} else {
			pre = pre.Next
		}
	}
	return newHead.Next
}

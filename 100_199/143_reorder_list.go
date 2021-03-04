package _100_199

/*
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:
Given 1->2->3->4, reorder it to 1->4->2->3.

Example 2:
Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	mid := getMid(head)
	t := mid.Next
	mid.Next = nil
	tail := reverseList(t)
	newHead := &ListNode{}
	cur := newHead
	for head != nil || tail != nil {
		if head != nil {
			cur.Next = head
			head = head.Next
			cur = cur.Next
		}
		if tail != nil {
			cur.Next = tail
			tail = tail.Next
			cur = cur.Next
		}
	}
	head = newHead.Next
}

func reverseList(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	p, q := l, l.Next
	l.Next = nil
	for q != nil {
		t := q.Next
		q.Next = p
		p = q
		q = t
	}
	return p
}

func getMid(head *ListNode) *ListNode {
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

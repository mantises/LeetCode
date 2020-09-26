package _100_199

/*
Given the head of a linked list, return the list after sorting it in ascending
order.

Follow up: Can you sort the linked list in O(n logn) time and O(1) memory
(i.e. constant space)?
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMidNode(head)
	left := sortList(head)
	right := sortList(mid)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
		cur.Next = nil
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

func getMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	if fast.Next == nil {
		tmp := slow.Next
		slow.Next = nil
		return tmp
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow.Next
	slow.Next = nil
	return tmp
}

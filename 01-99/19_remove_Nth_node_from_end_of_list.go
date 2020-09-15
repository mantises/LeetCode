package _01_99

func removeNthFromEndByFor(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	//target := newHead
	//cur := target

	return newHead.Next
}

func removeNthFromEndByDoublePointer(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	target := newHead
	cur := target
	for i := 0; i < n; i++ {
		target = target.Next
	}
	for target.Next != nil {
		cur = cur.Next
		target = target.Next
	}
	cur.Next = cur.Next.Next
	target.Next = nil
	return newHead.Next
}

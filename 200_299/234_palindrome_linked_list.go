package _200_299

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	mid := getMidNode(head)
	mid = reverseList(mid)
	p := head
	for mid != nil {
		if mid.Val != p.Val {
			return false
		}
		mid = mid.Next
		p = p.Next
	}
	return true
}
func getMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}
	return slow
}

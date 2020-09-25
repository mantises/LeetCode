package _200_299

import (
	"testing"
)

func TestFindTheDuplicate(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	t.Log(findDuplicateByBinarySearch(nums))
}

func TestProduct(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	t.Log(productExceptSelf(nums))
}

func TestReverseLinkedList(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	reverseList(l1)
}

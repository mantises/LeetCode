package _1600_1699

import (
	"fmt"
	"testing"
)

func TestMergeInBetween(t *testing.T) {
	l1 := &ListNode{Val: 00}
	l2 := &ListNode{Val: 11}
	l3 := &ListNode{Val: 22}
	l4 := &ListNode{Val: 33}
	l5 := &ListNode{Val: 44}
	l6 := &ListNode{Val: 55}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6

	l7 := &ListNode{Val: 77}
	l8 := &ListNode{Val: 88}
	l7.Next = l8

	fmt.Println(mergeInBetween(l1, 3, 5, l7).Val)
}

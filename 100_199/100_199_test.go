package _100_199

import (
	"fmt"
	"testing"
)

func TestMajorityElementMooreVote(t *testing.T) {
	nums := []int{2, 3, 4, 3, 2, 2, 5, 2, 2}
	t.Log(majorityElementMooreVote(nums))
}

func TestLRUCache(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(2, 6)
	lRUCache.Put(1, 5)
	t.Log(lRUCache.Get(1)) // return -1 (not found)

	lRUCache.Put(1, 2) // evicts key 1

	t.Log(lRUCache.Get(1)) // return -1 (not found)
	t.Log(lRUCache.Get(2)) // return 3
}

func TestGetIntersectionNode(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	// l1->l2->l3->l4
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	// l5->l3->l4
	l5.Next = l3
	getIntersectionNode(l1, l5)
}

func TestSortList(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 5}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 9}
	l5 := &ListNode{Val: 2}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	sortList(l1)
	fmt.Println(l1.Val)
}

func TestDetectCycle(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l3
	fmt.Println(detectCycle(l1).Val)
}

func TestConstructorMinStack(t *testing.T) {
	obj := ConstructorMinStack()
	obj.Push(-10)
	obj.Push(10)
	obj.Push(-2)
	obj.Push(-20)
	param3 := obj.Top()
	param4 := obj.GetMin()
	fmt.Println(param3, param4)
	obj.Pop()
	param3 = obj.Top()
	param4 = obj.GetMin()
	fmt.Println(param3, param4)
	obj.Pop()
	param3 = obj.Top()
	param4 = obj.GetMin()
	fmt.Println(param3, param4)
}

package _700_799

import (
	"fmt"
	"testing"
)

// ["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex",
// "addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
// [[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]
func TestMyLinkedList(t *testing.T) {
	myList := Constructor()
	myList.AddAtHead(7)
	myList.AddAtHead(2)
	myList.AddAtHead(1)
	myList.AddAtIndex(3, 0)
	myList.DeleteAtIndex(2)
	myList.AddAtHead(6)
	myList.AddAtTail(4)
	fmt.Println(myList.Get(4))
	myList.AddAtHead(4)
	myList.AddAtIndex(5, 0)
	myList.AddAtHead(6)

	fmt.Println(myList.Get(1))

}

func TestPartitionLabels(t *testing.T) {
	fmt.Println(partitionLabels("babacabcbdefegdehijhklij"))
}

func TestMyHashMap(t *testing.T) {
	m := ConstructorMyHashMap()
	m.Put(1, 1)
	m.Put(2, 2)
	fmt.Println(m.Get(1))
	fmt.Println(m.Get(3))
	m.Put(2, 1)
	fmt.Println(m.Get(2))
	m.Remove(2)
	fmt.Println(m.Get(2))
}

func TestRemoveElements(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 2}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	fmt.Println(splitListToParts(l1, 3))
}

func TestSumNumbers(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	fmt.Println(splitBST(root, 8))
}

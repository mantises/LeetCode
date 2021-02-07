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

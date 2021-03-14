package _200_299

import (
	"fmt"
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

func TestPalindromeLinkedList(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l1.Next = l2
	fmt.Println(isPalindrome(l1))
}

func TestMaximalSquare(t *testing.T) {
	mat := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalSquare(mat))
}

func TestNumSquares(t *testing.T) {
	fmt.Println(numSquares(13))
}

func TestCountPrimes(t *testing.T) {
	fmt.Println(countPrimes(10))
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
	fmt.Println(removeElements(l1, 2))
}

func TestBinaryTreePaths(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 78,
				Left: &TreeNode{
					Val:  1443546,
					Left: nil,
					Right: &TreeNode{
						Val: 10,
					},
				},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 9999,
			Left: &TreeNode{
				Val: 119,
			},
		},
	}
	fmt.Println(binaryTreePaths(root))
}

func TestSerializeAndDeserialize(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2222,
				Left: &TreeNode{
					Val: 5555,
				},
			},
			Right: &TreeNode{
				Val: 8888,
			},
		},
		Right: &TreeNode{
			Val: 9999,
			Left: &TreeNode{
				Val: 3333,
			},
			Right: &TreeNode{
				Val: 6666,
			},
		},
	}
	c := &Codec{}
	v := c.serialize(root)
	fmt.Println(v)
	s := c.deserialize(v)
	fmt.Println(s)
}

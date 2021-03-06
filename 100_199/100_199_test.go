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

func TestBestTimeToBuySellStock(t *testing.T) {
	prices := []int{1, 3, 6, 8, 2, 5, 10, 11, 0}
	fmt.Println(maxProfit(prices))
}

func TestHorseRobber(t *testing.T) {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
	fmt.Println(robV2(nums))
}

func TestWordBreak(t *testing.T) {
	s := "leetcode"
	words := []string{"leet", "code"}
	fmt.Println(wordBreak(s, words))
}

func TestMaxProfitIII(t *testing.T) {
	nums := []int{4, 3, 5, 0, 3, 1, 4}
	fmt.Println(maxProfitIII(nums))
}

func TestMaxProductSubarray(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

func TestCandy(t *testing.T) {
	fmt.Println(candy([]int{1, 2, 2}))
}

func TestTwoSumArrayIsSorted(t *testing.T) {
	fmt.Println(twoSum([]int{1, 2, 4, 5, 8}, 7))
}

func TestValidPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("0P"))
}

func TestFindMin(t *testing.T) {
	fmt.Println(findMin([]int{3, 4, 6, -5, -2, -1, 0, 2}))
}

func TestFindMinII(t *testing.T) {
	fmt.Println(findMinII([]int{-7, -7, -7, -7, -7, -7, -7, -1, 0, 2}))
}

func TestInsertionSortList(t *testing.T) {
	l1 := &ListNode{Val: 11}
	l2 := &ListNode{Val: 222}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 44}
	l5 := &ListNode{Val: 55}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	fmt.Println(insertionSortList(l1))
}

func TestReorderList(t *testing.T) {
	l1 := &ListNode{Val: 11}
	l2 := &ListNode{Val: 222}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 44}
	l5 := &ListNode{Val: 55}
	l6 := &ListNode{Val: 66}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	reorderList(l1)
}

func TestCopyRandomList(t *testing.T) {
	l1 := &Node{Val: 11}
	l2 := &Node{Val: 222}
	l3 := &Node{Val: 3}
	l4 := &Node{Val: 44}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l1.Random = l2
	l2.Random = l1
	l3.Random = l4
	l4.Random = l3

	copyRandomList(l1)
}

func TestConvertSortedArray(t *testing.T) {
	s := sortedArrayToBST([]int{-3, 0, 1, 4, 6})
	fmt.Println(s)
}

func TestLevelOrderTraversalII(t *testing.T) {
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
	fmt.Println(levelOrderBottom(root))
}

func TestRightSideView(t *testing.T) {
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
	fmt.Println(rightSideView(root))
}

func TestRightSideViewDFS(t *testing.T) {
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
	fmt.Println(rightSideViewDFS(root))
}

func TestFlatten(t *testing.T) {
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
	flatten(root)
}

func TestBuildTreeFromPreAndInOrder(t *testing.T) {
	v := buildTreeFromPreAndInOrder([]int{3, 9, 2, 6, 20, 15, 7}, []int{2, 9, 6, 3, 15, 20, 7})
	fmt.Println(v)
}

func TestBuildTreeFromInAndPostOrder(t *testing.T) {
	v := buildTreeFromInAndPostOrder([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(v)
}

func TestIsBalanced(t *testing.T) {
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
	fmt.Println(isBalanced(root))
}

func TestZigzagLevelOrder(t *testing.T) {
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
	fmt.Println(zigzagLevelOrder(root))
}

func TestConstructorBSTIterator(t *testing.T) {
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
	construct := ConstructorBSTIterator(root)
	for construct.HasNext() {
		fmt.Println(construct.Next())
	}
}

func TestPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  10,
				Left: nil,
				Right: &TreeNode{
					Val: 10089,
				},
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
				Val: 101,
			},
		},
	}
	fmt.Println(pathSum(root, 10100))
}

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val: 9,
		},
		// Right: &TreeNode{
		// 	Val: 20,
		// 	Left: &TreeNode{
		// 		Val: 15,
		// 	},
		// 	Right: &TreeNode{
		// 		Val: 7,
		// 	},
		// },
	}
	fmt.Println(maxPathSum(root))
}

func TestSumNumbers(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	fmt.Println(sumNumbers(root))
}

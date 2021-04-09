package _800_899

import (
	"fmt"
	"testing"
)

func TestParentPath(t *testing.T) {
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
	fmt.Println(distanceK(root, root.Left.Right.Left.Right, 3))
}

func TestConstructFromPrePost(t *testing.T) {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	post := []int{4, 5, 2, 6, 7, 3, 1}
	root := constructFromPrePost(pre, post)
	fmt.Printf("%v\n", root)
}

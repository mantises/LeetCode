package _600_699

import (
	"math"
	"strconv"
)

/*
Print a binary tree in an m*n 2D string array following these rules:

The row number m should be equal to the height of the given binary tree.
The column number n should always be an odd number.
The root node's value (in string format) should be put in the exactly middle of the first
row it can be put. The column and the row where the root node belongs will separate the
rest space into two parts (left-bottom part and right-bottom part). You should print the
left subtree in the left-bottom part and print the right subtree in the right-bottom part.
The left-bottom part and the right-bottom part should have the same size. Even if one
subtree is none while the other is not, you don't need to print anything for the none
subtree but still need to leave the space as large as that for the other subtree.
However, if two subtrees are none, then you don't need to leave space for both of them.
Each unused space should contain an empty string "".
Print the subtrees following the same rules.
Example 1:
Input:
     1
    /
   2
Output:
[["", "1", ""],
 ["2", "", ""]]
Example 2:
Input:
     1
    / \
   2   3
    \
     4
Output:
[["", "", "", "1", "", "", ""],
 ["", "2", "", "", "", "3", ""],
 ["", "", "4", "", "", "", ""]]
Example 3:
Input:
      1
     / \
    2   5
   /
  3
 /
4
Output:

[["",  "",  "", "",  "", "", "", "1", "",  "",  "",  "",  "", "", ""]
 ["",  "",  "", "2", "", "", "", "",  "",  "",  "",  "5", "", "", ""]
 ["",  "3", "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]
 ["4", "",  "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]]
Note: The height of binary tree is in the range of [1, 10].

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	// 求二叉树深度
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := dfs(root.Left), dfs(root.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
	m := dfs(root)
	n := int(math.Pow(2, float64(m)) - 1)
	// 构造结果切片
	var res [][]string
	for i := 0; i < m; i++ {
		v := []string{}
		for j := 0; j < n; j++ {
			v = append(v, "")
		}
		res = append(res, v)
	}

	// 在结果切片上填值
	var draw func(*TreeNode, int, int, int)
	draw = func(root *TreeNode, i int, l int, r int) {
		if root == nil {
			return
		}
		mid := (l + r) / 2
		res[i][mid] = strconv.Itoa(root.Val)
		draw(root.Left, i+1, l, mid)
		draw(root.Right, i+1, mid+1, r)
	}
	draw(root, 0, 0, len(res[0]))
	return res
}

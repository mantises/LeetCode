package _100_199

/*
Given the root of a binary tree, return the bottom-up level order traversal of its nodes'
values. (i.e., from left to right, level by level from leaf to root).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	var tmp []*TreeNode
	tmp = append(tmp, root)
	for len(tmp) > 0 {
		var nodes []*TreeNode
		var t []int
		for i := 0; i < len(tmp); i++ {
			if tmp[i].Left != nil {
				nodes = append(nodes, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				nodes = append(nodes, tmp[i].Right)
			}
			t = append(t, tmp[i].Val)
		}
		res = append(res, t)
		tmp = nodes
	}

	return reverse(res)
}

func reverse(s [][]int) [][]int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

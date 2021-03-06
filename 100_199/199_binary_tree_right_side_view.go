package _100_199

/*
Given the root of a binary tree, imagine yourself standing on the right side of it, return
the values of the nodes you can see ordered from top to bottom.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:
Input: root = [1,null,3]
Output: [1,3]

Example 3:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var t []int
	var tmp []*TreeNode
	tmp = append(tmp, root)
	for len(tmp) > 0 {
		var nodes []*TreeNode
		t = append(t, tmp[0].Val)
		for i := 0; i < len(tmp); i++ {
			if tmp[i].Right != nil {
				nodes = append(nodes, tmp[i].Right)
			}
			if tmp[i].Left != nil {
				nodes = append(nodes, tmp[i].Left)
			}
		}
		tmp = nodes
	}
	return t
}

func rightSideViewDFS(root *TreeNode) []int {
	var res []int
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, depth int, res *[]int) {
	if root == nil {
		return
	}
	if depth == len(*res) {
		*res = append(*res, root.Val)
	}
	depth++
	dfs(root.Right, depth, res)
	dfs(root.Left, depth, res)
}

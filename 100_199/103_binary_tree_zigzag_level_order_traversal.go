package _100_199

/*
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
(i.e., from left to right, then right to left for the next level and alternate between).


Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 2000].
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	res := make([][]int, 0)
	for level := 0; len(stack) != 0; level++ {
		tmp := make([]*TreeNode, 0)
		val := make([]int, 0)
		for i := 0; i < len(stack); i++ {
			val = append(val, stack[i].Val)
			if stack[i].Left != nil {
				tmp = append(tmp, stack[i].Left)
			}
			if stack[i].Right != nil {
				tmp = append(tmp, stack[i].Right)
			}
		}
		if level%2 != 0 {
			for j := 0; j < len(val)/2; j++ {
				val[j], val[len(val)-1-j] = val[len(val)-1-j], val[i]
			}
		}
		stack = tmp
		res = append(res, val)
	}
	return res
}

package _500_599

/*
Given the root of a binary tree, return an array of the largest value in each row of
the tree (0-indexed).


Example 1:
Input: root = [1,3,2,5,3,null,9]
Output: [1,3,9]

Example 2:
Input: root = [1,2,3]
Output: [1,3]

Example 3:
Input: root = [1]
Output: [1]

Example 4:
Input: root = [1,null,2]
Output: [1,2]

Example 5:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree will be in the range [0, 10^4].
-2^31 <= Node.val <= 2^31 - 1

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		var tmp []*TreeNode
		res = append(res, stack[0].Val)
		for i := 0; i < len(stack); i++ {
			if stack[i].Left != nil {
				tmp = append(tmp, stack[i].Left)
			}
			if stack[i].Right != nil {
				tmp = append(tmp, stack[i].Right)
			}
			if res[len(res)-1] < stack[i].Val {
				res[len(res)-1] = stack[i].Val
			}
		}
		stack = tmp
	}
	return res
}

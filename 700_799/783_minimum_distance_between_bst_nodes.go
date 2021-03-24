package _700_799

import "math"

/*
Given the root of a Binary Search Tree (BST), return the minimum difference between the
values of any two different nodes in the tree.

Note: This question is the same as 530:
https://leetcode.com/problems/minimum-absolute-difference-in-bst/

Example 1:
Input: root = [4,2,6,1,3]
Output: 1

Example 2:
Input: root = [1,0,48,null,null,12,49]
Output: 1


Constraints:
The number of nodes in the tree is in the range [2, 100].
0 <= Node.val <= 10^5

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minDiffInBST(root *TreeNode) int {
	res := math.MaxInt32
	var pre *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre == nil {
			pre = node
		} else {
			res = min(res, diff(pre, node))
			pre = node
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func minDiffInBSTNonRecursive(root *TreeNode) int {
	res := math.MaxInt32
	var pre *TreeNode
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		if len(stack) != 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = min(res, diff(pre, cur))
			cur, pre = cur.Right, cur
		}
	}
	return res
}

func diff(pre, cur *TreeNode) int {
	if pre == nil {
		return math.MaxInt32
	}
	if pre.Val >= cur.Val {
		return pre.Val - cur.Val
	}
	return cur.Val - pre.Val
}

package _500_599

import "math"

/*
Given a binary search tree with non-negative values, find the minimum absolute difference
between values of any two nodes.

Example:
Input:

   1
    \
     3
    /
   2

Output:
1
Explanation:
The minimum absolute difference is 1, which is the difference between 2 and 1 (or between
2 and 3).


Note:
There are at least two nodes in this BST.
This question is the same as 783:
https://leetcode.com/problems/minimum-distance-between-bst-nodes/

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
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

func getMinimumDifferenceNonRecursive(root *TreeNode) int {
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

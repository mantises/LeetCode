package _500_599

/*
Given the root of a tree, you are asked to find the most frequent subtree sum. The
subtree sum of a node is defined as the sum of all the node values formed by the
subtree rooted at that node (including the node itself). So what is the most frequent
subtree sum value? If there is a tie, return all the values with the highest frequency
in any order.

Examples 1
Input:

  5
 /  \
2   -3
return [2, -3, 4], since all the values happen only once, return all of them in any order.

Examples 2
Input:

  5
 /  \
2   -5
return [2], since 2 happens twice, however -5 only occur once.
Note: You may assume the sum of values in any subtree is in the range of 32-bit signed
integer.

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	hash := make(map[int]int)
	maxCount := 0
	var dfs func(node *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		cur := left + right + root.Val
		hash[cur]++
		if hash[cur] > maxCount {
			maxCount = hash[cur]
		}
		return cur
	}
	dfs(root)
	res := make([]int, 0)
	for k, v := range hash {
		if v == maxCount {
			res = append(res, k)
		}
	}
	return res
}

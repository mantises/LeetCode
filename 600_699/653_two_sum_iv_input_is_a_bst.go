package _600_699

/*
Given the root of a Binary Search Tree and a target number k, return true if there exist
two elements in the BST such that their sum is equal to the given target.


Example 1:
Input: root = [5,3,6,2,4,null,7], k = 9
Output: true

Example 2:
Input: root = [5,3,6,2,4,null,7], k = 28
Output: false

Example 3:
Input: root = [2,1,3], k = 4
Output: true

Example 4:
Input: root = [2,1,3], k = 1
Output: false

Example 5:
Input: root = [2,1,3], k = 3
Output: true


Constraints:
The number of nodes in the tree is in the range [1, 10^4].
-10^4 <= Node.val <= 10^4
root is guaranteed to be a valid binary search tree.
-10^5 <= k <= 10^5

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	set := make(map[int]bool, 0)
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, val int) bool {
		if node == nil {
			return false
		}
		left := dfs(node.Left, val)
		if _, ok := set[k-node.Val]; ok {
			return true
		} else {
			set[node.Val] = true
		}
		right := dfs(node.Right, val)
		return left || right
	}
	return dfs(root, k)
}

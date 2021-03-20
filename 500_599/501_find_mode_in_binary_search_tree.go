package _500_599

/*
Given the root of a binary search tree (BST) with duplicates, return all the mode(s)
(i.e., the most frequently occurred element) in it.

If the tree has more than one mode, return them in any order.

Assume a BST is defined as follows:
The left subtree of a node contains only nodes with keys less than or equal to the node's
key.
The right subtree of a node contains only nodes with keys greater than or equal to the
node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:
Input: root = [1,null,2,2]
Output: [2]

Example 2:
Input: root = [0]
Output: [0]


Constraints:
The number of nodes in the tree is in the range [1, 10^4].
-10^5 <= Node.val <= 10^5

Follow up: Could you do that without using any extra space? (Assume that the implicit
stack space incurred due to recursion does not count).

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			res = append(res, base)
		} else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}
	var pre *TreeNode
	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
		} else {
			pre = cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else {
				pre.Right = nil
				update(cur.Val)
				cur = cur.Right
			}
		}
	}
	return res
}

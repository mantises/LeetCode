package _600_699

/*
Given the root of a binary search tree and the lowest and highest boundaries as low and
high, trim the tree so that all its elements lies in [low, high]. Trimming the tree
should not change the relative structure of the elements that will remain in the tree
(i.e., any node's descendant should remain a descendant). It can be proven that there
is a unique answer.

Return the root of the trimmed binary search tree. Note that the root may change
depending on the given bounds.


Example 1:
Input: root = [1,0,2], low = 1, high = 2
Output: [1,null,2]

Example 2:
Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
Output: [3,2,null,1]

Example 3:
Input: root = [1], low = 1, high = 2
Output: [1]

Example 4:
Input: root = [1,null,2], low = 1, high = 3
Output: [1,null,2]

Example 5:
Input: root = [1,null,2], low = 2, high = 4
Output: [2]


Constraints:
The number of nodes in the tree in the range [1, 10^4].
0 <= Node.val <= 10^4
The value of each node in the tree is unique.
root is guaranteed to be a valid binary search tree.
0 <= low <= high <= 10^4

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func trimBSTNonRecursive(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	l, r := root, root
	for l != nil && l.Left != nil {
		if l.Left.Val < low {
			l.Left = l.Left.Right
		} else {
			l = l.Left
		}
	}
	for r != nil && r.Right != nil {
		if r.Right.Val > high {
			r.Right = r.Right.Left
		} else {
			r = r.Right
		}
	}
	return root
}

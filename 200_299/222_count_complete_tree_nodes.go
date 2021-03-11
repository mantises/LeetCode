package _200_299

/*
Given the root of a complete binary tree, return the number of the nodes in the tree.

According to Wikipedia, every level, except possibly the last, is completely filled in a
complete binary tree, and all nodes in the last level are as far left as possible. It can
have between 1 and 2h nodes inclusive at the last level h.


Example 1:
Input: root = [1,2,3,4,5,6]
Output: 6

Example 2:
Input: root = []
Output: 0

Example 3:
Input: root = [1]
Output: 1


Constraints:

The number of nodes in the tree is in the range [0, 5 * 10^4].
0 <= Node.val <= 5 * 10^4
The tree is guaranteed to be complete.


Follow up: Traversing the tree to count the number of nodes in the tree is an easy
solution but with O(n) complexity. Could you find a faster algorithm?

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodesII(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth, rDepth := getCompleteTreeLevel(root.Left), getCompleteTreeLevel(root.Right)
	if lDepth == rDepth {
		return countNodes(root.Right) + (1 << lDepth)
	} else {
		return countNodes(root.Left) + (1 << rDepth)
	}
}

func getCompleteTreeLevel(root *TreeNode) uint {
	level := uint(0)
	for root != nil {
		root = root.Left
		level += 1
	}
	return level
}

func countNodesBinarySearch(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 二叉树深度
	depth := getCompleteTreeLevel(root)
	depthPrev := depth - 1

	start, end := 1, 1<<depthPrev
	for start <= end {
		mid := start + ((end - start) >> 1)
		if isExist(root, mid, depthPrev) {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	// start - 1为最后一层节点数
	ret := (1 << depthPrev) - 1 + start - 1
	return ret
}

/*
 * 功能： 判断最后一层第index个索引是否存在
 * root： 二叉树根节点
 * index：判断最后一层索引为index的节点是否存在, 索引范围是[1, 2^depth]
 * depth：倒数第二层的深度, 这是因为满二叉树最后一层的节点数等于 2^depth
 */
func isExist(root *TreeNode, index int, depth uint) bool {
	node := root
	for depth > 0 {
		// 最后一层分界线
		mid := (1 << depth) >> 1
		if index > mid {
			// 如果在右子树，需要更新索引值
			index -= mid
			node = node.Right
		} else {
			node = node.Left
		}
		depth -= 1
	}
	return node != nil
}

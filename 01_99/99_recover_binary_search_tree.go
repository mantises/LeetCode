package _01_99

/*
You are given the root of a binary search tree (BST), where exactly two nodes of the tree
were swapped by mistake. Recover the tree without changing its structure.

Follow up: A solution using O(n) space is pretty straight forward. Could you devise a
constant space solution?


Example 1:
Input: root = [1,3,null,null,2]
Output: [3,1,null,null,2]
Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST
valid.

Example 2:
Input: root = [3,1,4,null,null,2]
Output: [2,1,4,null,null,3]
Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes
the BST valid.


Constraints:
The number of nodes in the tree is in the range [2, 1000].
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
func recoverTree(root *TreeNode) {
	// 需要交换的第一第二节点指针 / 上一个处理的指针 / 左子树最右节点的虚拟右节点
	var first, second, prev, predecessor *TreeNode
	for root != nil {
		if root.Left != nil {
			// 根据当前节点，找到其前序节点
			predecessor = root.Left
			// 一直寻找左节点的最右节点
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			// 如果前序节点的右孩子是空，那么把前序节点的右孩子指向当前节点，
			// 然后进入当前节点的左孩子
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
				continue
			}
			// 如果当前节点的前序节点其右孩子指向了它本身，那么把前序节点的右孩子设置为空
			// 打印当前节点，然后进入右孩子。
			predecessor.Right = nil
		}
		// 如果当前节点的左孩子为空，打印当前节点，然后进入右孩子
		if prev != nil && root.Val < prev.Val {
			// 第一个指针不可以重复设置
			if first == nil {
				first = prev
			}
			second = root
		}
		prev = root
		root = root.Right
	}
	first.Val, second.Val = second.Val, first.Val
}

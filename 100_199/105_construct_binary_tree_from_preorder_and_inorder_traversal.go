package _100_199

import "fmt"

/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal
of a binary tree and inorder is the inorder traversal of the same tree, construct and
return the binary tree.


Example 1:
Input: preorder = [3,2,9,6,20,15,7], inorder = [2,9,6,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]


Constraints:
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}
	fmt.Println(inorder[:index], inorder[index+1:])
	fmt.Println(preorder[1:1+index], preorder[index+1:])
	left := buildTree(preorder[1:1+index], inorder[:index])
	right := buildTree(preorder[index+1:], inorder[index+1:])
	return &TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}
}

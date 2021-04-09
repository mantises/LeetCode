package _800_899

/*
Return any binary tree that matches the given preorder and postorder traversals.

Values in the traversals pre and post are distinct positive integers.


Example 1:
Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
Output: [1,2,3,4,5,6,7]

Note:
1 <= pre.length == post.length <= 30
pre[] and post[] are both permutations of 1, 2, ..., pre.length.
It is guaranteed an answer exists. If there exists multiple answers, you can return any
of them.

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	if len(pre) == 1 {
		return &TreeNode{Val: pre[0]}
	}
	index := 0
	for ; index < len(post); index++ {
		if post[index] == pre[1] {
			break
		}
	}
	left := constructFromPrePost(pre[1:index+2], post[0:index+1])
	right := constructFromPrePost(pre[index+2:], post[index+1:len(pre)-1])
	return &TreeNode{
		Left:  left,
		Right: right,
		Val:   pre[0],
	}
}

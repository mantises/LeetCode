package _600_699

/*
Given a binary tree, write a function to get the maximum width of the given tree. The
maximum width of a tree is the maximum width among all levels.

The width of one level is defined as the length between the end-nodes (the leftmost and
right most non-null nodes in the level, where the null nodes between the end-nodes are
also counted into the length calculation.

It is guaranteed that the answer will in the range of 32-bit signed integer.

Example 1:
Input:

           1
         /   \
        3     2
       / \     \
      5   3     9

Output: 4
Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).

Example 2:
Input:

          1
         /
        3
       / \
      5   3

Output: 2
Explanation: The maximum width existing in the third level with the length 2 (5,3).

Example 3:
Input:

          1
         / \
        3   2
       /
      5
Output: 2
Explanation: The maximum width existing in the second level with the length 2 (3,2).

Example 4:
Input:

          1
         / \
        3   2
       /     \
      5       9
     /         \
    6           7
Output: 8
Explanation:The maximum width existing in the fourth level with the length 8
(6,null,null,null,null,null,null,7).


Constraints:
The given binary tree will have between 1 and 3000 nodes.

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type item struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*item, 0)
	queue = append(queue, &item{node: root, index: 1})
	width := 0
	for len(queue) != 0 {
		tmp := make([]*item, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i].node.Left != nil {
				tmp = append(tmp, &item{
					node:  queue[i].node.Left,
					index: 2 * queue[i].index,
				})
			}
			if queue[i].node.Right != nil {
				tmp = append(tmp, &item{
					node:  queue[i].node.Right,
					index: 2*queue[i].index + 1,
				})
			}
		}
		if queue[len(queue)-1].index-queue[0].index+1 > width {
			width = queue[len(queue)-1].index - queue[0].index + 1
		}
		queue = tmp
	}
	return width
}

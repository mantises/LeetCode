package _600_699

import (
	"strconv"
)

/*
You need to construct a string consists of parenthesis and integers from a binary tree with the
preorder traversing way.

The null node needs to be represented by empty parenthesis pair "()". And you need to omit all
the empty parenthesis pairs that don't affect the one-to-one mapping relationship between the
string and the original binary tree.

Example 1:
Input: Binary tree: [1,2,3,4]
       1
     /   \
    2     3
   /
  4
Output: "1(2(4))(3)"
Explanation: Originallay it needs to be "1(2(4)())(3()())",
but you need to omit all the unnecessary empty parenthesis pairs.
And it will be "1(2(4))(3)".

Example 2:
Input: Binary tree: [1,2,3,null,4]
       1
     /   \
    2     3
     \
      4
Output: "1(2()(4))(3)"
Explanation: Almost the same as the first example,
except we can't omit the first parenthesis pair to break the one-to-one mapping relationship
between the input and the output.

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}
	var left, right string
	if t.Left != nil {
		left = tree2str(t.Left)
	}
	if t.Right != nil {
		right = tree2str(t.Right)
	}
	str := strconv.Itoa(t.Val)
	if len(right) == 0 {
		str += "(" + left + ")"
	} else {
		str += "(" + left + ")" + "(" + right + ")"
	}
	return str
}

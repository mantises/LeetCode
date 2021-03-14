package _200_299

import (
	"strconv"
	"strings"
)

/*
Serialization is the process of converting a data structure or object into a sequence
of bits so that it can be stored in a file or memory buffer, or transmitted across a
network connection link to be reconstructed later in the same or another computer
environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction
on how your serialization/deserialization algorithm should work. You just need to ensure
that a binary tree can be serialized to a string and this string can be deserialized to
the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary
tree. You do not necessarily need to follow this format, so please be creative and come
up with different approaches yourself.


Example 1:
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Example 4:
Input: root = [1,2]
Output: [1,2]


Constraints:
The number of nodes in the tree is in the range [0, 10^4].
-1000 <= Node.val <= 1000

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := c.serialize(root.Left)
	right := c.serialize(root.Right)
	return strconv.Itoa(root.Val) + "," + left + "," + right
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	return c.buildTree(&list)
}

func (c *Codec) buildTree(list *[]string) *TreeNode {
	rootVal := (*list)[0]
	*list = (*list)[1:]
	if rootVal == "#" {
		return nil
	}
	Val, _ := strconv.Atoi(rootVal)
	root := &TreeNode{Val: Val}
	root.Left = c.buildTree(list)
	root.Right = c.buildTree(list)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

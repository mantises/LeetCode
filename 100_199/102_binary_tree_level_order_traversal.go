package _100_199

/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]

*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	var d1 []*TreeNode
	d1 = append(d1, root)
	for len(d1) != 0 {
		var t []int
		var d2 []*TreeNode
		for i, v := range d1 {
			t = append(t, v.Val)
			if v.Left != nil {
				d2 = append(d2, v.Left)
			}
			if v.Right != nil {
				d2 = append(d2, v.Right)
			}
			if i == len(d1)-1 {
				res = append(res, t)
				d1, d2 = d2, d1
			}
		}
	}
	return res
}

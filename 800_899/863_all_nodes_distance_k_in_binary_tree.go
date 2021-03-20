package _800_899

/*
We are given a binary tree (with root node root), a target node, and an integer value K.

Return a list of the values of all nodes that have a distance K from the target node.
The answer can be returned in any order.



Example 1:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
Output: [7,4,1]
Explanation:
The nodes that are a distance 2 from the target node (with value 5)
have values 7, 4, and 1.



Note that the inputs "root" and "target" are actually TreeNodes.
The descriptions of the inputs above are just serializations of these objects.


Note:
The given tree is non-empty.
Each node in the tree has unique values 0 <= node.val <= 500.
The target node is a node in the tree.
0 <= K <= 1000.

*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// parents 存了从root到target的父节点
	parents := make([]*TreeNode, 0, 500)
	// 用于存储父节点到parents
	var parentPath func(root, target *TreeNode) bool
	parentPath = func(root, target *TreeNode) bool {
		if root == nil {
			return false
		}
		if root == target {
			return true
		}
		parents = append(parents, root)
		if parentPath(root.Left, target) || parentPath(root.Right, target) {
			return true
		}
		parents = parents[:len(parents)-1]
		return false
	}
	parentPath(root, target)

	// 先从target子树下获取所有的节点
	res = append(res, find(target, K)...)

	// 从target的父节点处另一边的子树下获取所有节点
	depth := 1
	pre := target
	for len(parents) > 0 && depth <= K {
		last := parents[len(parents)-1]
		parents = parents[:len(parents)-1]
		if K == depth {
			res = append(res, last.Val)
		} else if pre == last.Left {
			res = append(res, find(last.Right, K-depth-1)...)
		} else {
			res = append(res, find(last.Left, K-depth-1)...)
		}
		depth++
		pre = last
	}
	return res
}

// 递归获取距离指定节点root的k个单位远的数值
func find(root *TreeNode, k int) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	if k == 0 {
		res = append(res, root.Val)
		return res
	}
	l := find(root.Left, k-1)
	r := find(root.Right, k-1)
	res = append(res, l...)
	res = append(res, r...)
	return res
}

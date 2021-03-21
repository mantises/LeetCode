package _400_499

/*
You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

Example:

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:
1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	var res int
	// 前缀和  + 递归
	var dfs func(*TreeNode, map[int]int, int)
	dfs = func(node *TreeNode, prefix map[int]int, curSum int) {
		if node == nil {
			return
		}
		// 更新前缀和
		curSum += node.Val
		// 当前路径中存在以当前节点为终点的和为sum的子路径
		if v, ok := prefix[curSum-sum]; ok {
			res += v
		}
		prefix[curSum]++                // 将当前节点加入路径
		dfs(node.Left, prefix, curSum)  // 在其左子树中递归寻找
		dfs(node.Right, prefix, curSum) // 在其右子树中递归寻找
		prefix[curSum]--
	}
	dfs(root, map[int]int{0: 1}, 0)
	return res
}

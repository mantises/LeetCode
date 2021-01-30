package _300_399

/*
The thief has found himself a new place for his thievery again. There is only one entrance to this
area, called the "root." Besides the root, each house has one and only one parent house. After a
tour, the smart thief realized that "all houses in this place forms a binary tree". It will
automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

Example 1:
Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.

Example 2:
Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.


*/
// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	levelSum := make([]int, 0)
	levelNode := make([]*TreeNode, 0)
	nextLevel := make([]*TreeNode, 0)
	levelNode = append(levelNode, root)
	for len(levelNode) > 0 {
		tmp := 0
		for _, v := range levelNode {
			tmp += v.Val
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}
		}
		levelSum = append(levelSum, tmp)
		levelNode, nextLevel = nextLevel, make([]*TreeNode, 0)
	}

	return robber(levelSum)
}

func robber(levelSum []int) int {
	if len(levelSum) == 0 {
		return 0
	}
	if len(levelSum) == 1 {
		return levelSum[0]
	}
	if len(levelSum) == 2 {
		return max(levelSum[0], levelSum[1])
	}
	first, second := levelSum[0], max(levelSum[0], levelSum[1])
	for i := 2; i < len(levelSum); i++ {
		first, second = second, max(first+levelSum[i], second)
	}
	return second
}

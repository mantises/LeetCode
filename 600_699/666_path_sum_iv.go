package _600_699

/*
If the depth of a tree is smaller than 5, then this tree can be represented by a list of
three-digits integers.

For each integer in this list:

The hundreds digit represents the depth D of this node, 1 <= D <= 4.
The tens digit represents the position P of this node in the level it belongs to, 1 <= P
<= 8. The position is the same as that in a full binary tree.
The units digit represents the value V of this node, 0 <= V <= 9.
Given a list of ascending three-digits integers representing a binary tree with the depth
smaller than 5, you need to return the sum of all paths from the root towards the leaves.

It's guaranteed that the given list represents a valid connected binary tree.

Example 1:
Input: [113, 215, 221]
Output: 12
Explanation:
The tree that the list represents is:
    3
   / \
  5   1

The path sum is (3 + 5) + (3 + 1) = 12.


Example 2:
Input: [113, 221]
Output: 4
Explanation:
The tree that the list represents is:
    3
     \
      1

The path sum is (3 + 1) = 4.

*/
func pathSum(nums []int) int {
	tree := make([]int, 16)
	for i := range tree {
		tree[i] = -1
	}
	for i := range nums {
		pos, dep := nums[i]/10%10-1, uint(nums[i]/100%10-1)
		tree[1<<dep-1+pos] = nums[i] % 10 // 2^2 是异或
	}
	return helper(tree, 0, 0)
}

func helper(tree []int, root, sum int) int {
	if tree[root] == -1 {
		return 0
	}
	if root >= 7 || tree[2*root+1] == -1 && tree[2*root+2] == -1 { // 最后一层或者叶节点直接返回
		return tree[root] + sum
	}
	return helper(tree, 2*root+1, sum+tree[root]) + helper(tree, 2*root+2, sum+tree[root])
}

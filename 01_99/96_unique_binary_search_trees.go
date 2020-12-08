package _01_99

/*
Given n, how many structurally unique BST (binary search trees) that store values 1 ... n?

Example:
Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BSTs:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

Constraints:
1 <= n <= 19

*/
// sub problem
func numTrees(n int) int {
	if n <= 1 {
		return n
	}
	res := make([]int, n+1)
	res[0], res[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			res[i] += res[j-1] * res[i-j]
		}
	}
	return res[n]
}

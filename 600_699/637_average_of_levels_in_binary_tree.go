package _600_699

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var avg float64
		avg, nodes = average(nodes)
		res = append(res, avg)
	}
	return res
}

func average(nodes []*TreeNode) (float64, []*TreeNode) {
	count := float64(len(nodes))
	var sum float64
	var nextNodes []*TreeNode
	for _, v := range nodes {
		sum += float64(v.Val)
		if v.Left != nil {
			nextNodes = append(nextNodes, v.Left)
		}
		if v.Right != nil {
			nextNodes = append(nextNodes, v.Right)
		}
	}
	return sum / count, nextNodes
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	rootVal := root.Val

	var preOrder func(*TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil || ans != -1 && node.Val >= ans {
			return
		}
		if node.Val > rootVal {
			ans = node.Val
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(root)
	return ans
}

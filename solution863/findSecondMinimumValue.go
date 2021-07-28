package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) (ans []int) {
	parents := map[int]*TreeNode{}
	var buildRelations func(*TreeNode)

	buildRelations = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			buildRelations(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			buildRelations(node.Right)
		}
	}

	buildRelations(root)

	var calAns func(*TreeNode, *TreeNode, int)
	calAns = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k {
			ans = append(ans, node.Val)
			return
		}

		if node.Left != from {
			calAns(node.Left, node, depth + 1)
		}
		if node.Right != from {
			calAns(node.Right, node, depth + 1)
		}

		if parents[node.Val] != from {
			calAns(parents[node.Val], node, depth + 1)
		}
	}

	calAns(target, nil, 0)
	return ans
}

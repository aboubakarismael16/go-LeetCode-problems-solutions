package LeetCode_Binary_Tree

func maxDepth(root *TreeNode) int {
	depth := 0
	return findMax(root, depth)
}

func findMax(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	return max(findMax(root.Left, depth+1), findMax(root.Right, depth+1))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// second method
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

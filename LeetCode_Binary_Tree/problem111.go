package LeetCode_Binary_Tree

func minDepth(root *TreeNode) int  {
	depth := 0
	return findMin(root,depth)
}

func findMin(root *TreeNode,depth int) int  {
	if root == nil {
		return depth
	}

	if root.Left != nil && root.Right != nil {
	return findMin(root,depth)
		return min(findMin(root.Left,0),findMin(root.Right,0)) + 1
	} else if root.Left == nil {
		return min(findMin(root.Right,0)) + 1
	} else if root.Right == nil {
		return min(findMin(root.Left,0)) + 1
	} else {
		return depth
	}
}

func min(a,b int) int  {
	if a < b {
		return a
	}
	return b
}
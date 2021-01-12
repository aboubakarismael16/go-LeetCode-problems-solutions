package LeetCode_Binary_Tree

func isBalanced(root *TreeNode) bool  {
	if root == nil {
		return true
	}

	return abs(heighHelper(root.Left,0) - heighHelper(root.Right,0)) < 2 && 
					isBalanced(root.Left) && isBalanced(root.Right)
}

func heighHelper(root *TreeNode,heigh int) int  {
	if root == nil {
		return heigh
	}

	return max(heighHelper(root.Left,heigh + 1),heighHelper(root.Right,heigh + 1))
}

func abs(a int) int  {
	if a < 0 {
		return 0 -  a 
	}
	return a
}

func max(a,b int) int {
	if  a < b {
		return b
	}
	return a
}
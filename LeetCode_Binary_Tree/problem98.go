package LeetCode_Binary_Tree

import "math"


func isValidBST(root *TreeNode) bool  {
	return isValidBSTHelper(root,math.MinInt32,math.MaxInt32)

}

func isValidBSTHelper(root *TreeNode,min,max int) bool {
	if root == nil {
		return true
	}

	return isValidBSTHelper(root.Left,min,root.Val-1) &&
	       isValidBSTHelper(root.Right,root.Val+1,max)
}
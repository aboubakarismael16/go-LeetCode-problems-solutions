package LeetCode_Binary_Tree

func rangeSumBST(root *TreeNode,low ,high int) int {
	res := 0
	if root == nil {
		return res
	}
	if root.Val > low && root.Val > high {
		return rangeSumBST(root.Left,low,high)
	}else if root.Val < low && root.Val < high {
		return rangeSumBST(root.Right,low,high)
	}else if root.Val >= low && root.Val <= high {
		res += root.Val
		res += rangeSumBST(root.Right,low,high)
		res += rangeSumBST(root.Left,low,high)
	}

	return res
}
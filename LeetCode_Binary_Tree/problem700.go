package LeetCode_Binary_Tree

func searchBST(root *TreeNode,val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val < val {
		return searchBST(root.Right,val)
	} else if root.Val > val {
		return searchBST(root.Right,val)
	}

	return nil
}
package LeetCode_Binary_Tree

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	inorder(root,&sum)
	return root
}

func inorder(root *TreeNode,sum *int)  {
	if root == nil {
		return
	}
	inorder(root.Right,sum)
	*sum = *sum + root.Val
	root.Val = *sum
	inorder(root.Left,sum)
}
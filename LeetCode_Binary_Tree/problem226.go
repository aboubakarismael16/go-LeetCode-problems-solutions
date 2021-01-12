package LeetCode_Binary_Tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root.left
	root.Left = root.Right
	root.Right = temp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
package LeetCode_Binary_Tree


func inOrderTraversal(root *TreeNode) []int  {
	if root == nil {
		return nil
	}

	var res []int
	res = append(res,inOrderTraversal(root.Left))
	res = append(res,root.Val)
	res = append(res,inOrderTraversal(root.Right))
	return res
}
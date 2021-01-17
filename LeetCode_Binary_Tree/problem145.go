package LeetCode_Binary_Tree


func postOrderTraversal(root *TreeNode) []int  {
	if root == nil {
		return nil
	}

	var res []int
	res = append(res,postOrderTraversal(root.Left)...)
	res = append(res,postOrderTraversal(root.Right)...)
	res = append(res,root.Val)

	return res
}
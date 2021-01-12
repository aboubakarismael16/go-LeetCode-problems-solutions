package LeetCode_Binary_Tree


func preOrderTraversal(root *treeNode) []int  {
	if root == nil {
		return nil
	}

	res := []int{}
	res = append(res,root.Val)
	res = append(res,preOrderTraversal(root.Left)...)
	res = append(res,preOrderTraversal(root.Right)...)

	return res
}

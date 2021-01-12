package LeetCode_Binary_Tree

type treeNode struct {
	Val int
	Left *treeNode
	Right *treeNode
}

func postOrderTraversal(root *treeNode) []int  {
	if root == nil {
		return nil
	}

	var res []int
	res = append(res,postOrderTraversal(root.Left)...)
	res = append(res,postOrderTraversal(root.Right)...)
	res = append(res,root.Val)

	return res
}
package LeetCode_Binary_Tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}

	if root == nil {
		return res
	}
	helper(root,"",&res)
	return res
}

func helper(root *TreeNode,path string,res*[]string)  {
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res,path)
		return
	}
	path += "->"
	if root.Left != nil {
		helper(root.Left,path,res)
	}
	if root.Right != nil {
		helper(root.Right,path,res)
	}
}
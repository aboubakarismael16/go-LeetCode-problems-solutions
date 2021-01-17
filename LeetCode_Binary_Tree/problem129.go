package LeetCode_Binary_Tree

func sumNumbers(root *TreeNode) int  {
	res := 0
	if root == nil {
		return res
	}

	return sumNumberHelper(root,0)
}

func sumNumberHelper(root *TreeNode,num int) int  {
	if root == nil {
		return 0
	}
	num = root.Val + num * 10
	if root.Left == nil && root.Right == nil {
		return num 
	}

	sum := 0
	sum += sumNumberHelper(root.Left,num)
	sum += sumNumberHelper(root.Right,num)

	return sum
}
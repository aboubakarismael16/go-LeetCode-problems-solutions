package LeetCode_Binary_Tree

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	sumOfLeftLeavesHelper(root,&res)

	return res
}

func sumOfLeftLeavesHelper(root *TreeNode,res *int)  {
	if root == nil {
		return
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res += root.Left.Val
	}

	sumOfLeftLeavesHelper(root.Left,res)
	sumOfLeftLeavesHelper(root.Right,res)
}
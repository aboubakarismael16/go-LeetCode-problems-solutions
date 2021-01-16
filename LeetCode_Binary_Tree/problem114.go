package LeetCode_Binary_Tree

func flatten(root *TreeNode)  {
	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			left := root.Left
			for left.Right != nil {
				left = left.Right
			}
			left.Right = root.Right
			root.Right = root.Left
			root.Left = nil
			root = root.Right
		}
	}
	
}
package LeetCode_Binary_Tree




func sortArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		root := &TreeNode{}
		root.Val = nums[0]
		return root
	}

	mid := len(nums) / 2
	root := &TreeNode{}
	root.Val = nums[mid]
	root.Left = sortArrayToBST(nums[:mid])
	root.Right = sortArrayToBST(nums[mid+1:])
	return root
	
}


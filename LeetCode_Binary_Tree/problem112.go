package LeetCode_Binary_Tree

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

//second method

var allSums []int

func hasPathSum2(root *TreeNode, sum int) bool {
	getAllSums(root, 0)
	for _, val := range allSums {
		if sum == val {
			allSums = []int{}
			return true
		}
	}

	allSums = []int{}
	return false
}

func getAllSums(root *TreeNode, currSum int) {
	if root != nil {
		currSum += root.Val
		if root.Left == nil && root.Right == nil {
			allSums = append(allSums, currSum)
		} else {
			getAllSums(root.Left, currSum)
			getAllSums(root.Right, currSum)
		}
	}
}

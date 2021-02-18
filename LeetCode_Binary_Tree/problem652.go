package LeetCode_Binary_Tree

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	cache := make(map[string]int)
	result := make([]*TreeNode, 0)
	helper(root, cache, &result)
	return result
}

func helper(root *TreeNode, cache map[string]int, result *[]*TreeNode) string {
	if root == nil {
		return "|"
	}
	left, right := helper(root.Left, cache, result), helper(root.Right, cache, result)
	key := strconv.Itoa(root.Val) + "|" + left + "|" + right
	if cache[key] == 1 {
		*result = append(*result, root)
	}
	cache[key]++
	return key
}
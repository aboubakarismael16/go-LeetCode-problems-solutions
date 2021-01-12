package LeetCode_Binary_Tree

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    
    if root.Val == sum && root.Left == nil && root.Right == nil {
        return true
    }
    
    return hasPathSum(root.Left,sum-root.Val) || hasPathSum(root.Right,sum-root.Val)
}
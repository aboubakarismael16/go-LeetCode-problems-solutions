package LeetCode_Binary_Tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
    return isSameTreeHelper(p,q,true)
}

func isSameTreeHelper(p *TreeNode, q *TreeNode,res bool) bool {
    if p == nil && q == nil {
        return res
    } else if (p == nil && q != nil) || (p != nil && q == nil) || p.Val != q.Val {
        return false
    }
    
    return isSameTreeHelper(p.Left,q.Left,res) && isSameTreeHelper(p.Right,q.Right,res)
}
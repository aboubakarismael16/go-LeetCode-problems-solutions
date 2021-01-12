package LeetCode_Binary_Tree


func isSubPath(head *ListNode,root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}

	return isSubPathHelper(head,root) || isSubPath(head,root.Left) || isSubPath(head,root.Right)
}

func isSubPathHelper(head *ListNode,root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil && head != nil {
		return false
	}
	if root.Val != head.Val {
		return false
	}

	return isSubPathHelper(head.Next,root.Left) || isSubPathHelper(head.Next,root.Right) 
}
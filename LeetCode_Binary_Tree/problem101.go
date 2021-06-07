package LeetCode_Binary_Tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	if root.Left.Val != root.Right.Val {
		return false
	}

	var q []*TreeNode
	q = append(q, root.Left)
	q = append(q, root.Right)

	for len(q) > 0 {
		lenght := len(q)
		level := []int{}

		for i := 0; i < lenght; i++ {
			node := q[0]
			var nodeVal int
			if node != nil {
				nodeVal = node.Val
				q = append(q, node.Left)
				q = append(q, node.Right)
			}

			level = append(level, nodeVal)
			q = q[1:]
		}

		left, right := 0, len(level)-1
		for left < right {
			if level[left] != level[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

// second method

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root, root)
}

func helper(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && helper(node1.Right, node2.Left) && helper(node1.Left, node2.Right)
}

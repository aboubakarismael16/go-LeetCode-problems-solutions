package LeetCode_Linked_List

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next 
	} 
	newHead := reverse(head,node)
	head.Next = reverseKGroup(node,k)

	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	pre := last
	for first != last {
		temp := first.Next
		first.Next = pre
		pre = first
		first = temp
	}

	return pre
}
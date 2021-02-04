package LeetCode_Linked_List

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil || val == 0 {
		return head 
	}

	newHead := &ListNode{0,head}
	pre := newHead
	cur := head 
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}

		cur = cur.Next
	}

	return newHead.Next 
}
package LeetCode_Linked_List

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pre := &ListNode{-1,head}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next 
		}
	}

	return pre.Next
}
package LeetCode_Linked_List

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{-1,head}
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		nextNext := next.Next 

		pre.Next = next 
		next.Next = cur 
		cur.Next = nextNext

		pre = cur
		cur = cur.Next 
	}

	return dummy.Next
}
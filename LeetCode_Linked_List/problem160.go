package LeetCode_Linked_List

func getIntersectionNode(headA,headB *listNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a,b := headA,headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return b
}
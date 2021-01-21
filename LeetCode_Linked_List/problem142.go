package LeetCode_Linked_List

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow,fast := head,head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}	
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	third := head
	for third != slow {
		slow = slow.Next
		third = third.Next
	}

	return slow
}
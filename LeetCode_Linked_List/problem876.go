package LeetCode_Linked_List

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow,fast := head,head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next != nil && fast.Next.Next == nil {
		slow = slow.Next
	}

	return slow
}

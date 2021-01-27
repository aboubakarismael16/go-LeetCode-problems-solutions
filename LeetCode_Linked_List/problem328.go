package LeetCode_Linked_List

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd,even,evenHead := head,head.Next,head.Next
	for even != nil && even.Next != nil {
		oddNext := even.Next 
		odd.Next = oddNext
		odd = odd.Next 

		evenNext := odd.Next 
		even.Next = evenNext 
		even = even.Next 
	}
	odd.Next = evenHead

	return head
}

// h
// o, e, o, e, o, e, o
// 2->1->3->5->6->4->7->NULL
// 
//2->3->6->7->1->5->4->NULL
package LeetCode_Linked_List

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{0,nil}
	pre.Next = head
	left := pre
	right := pre

	for i := 0; i < n+1; i++ {
		right = right.Next
	}
	for right != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next

	return pre.Next
}
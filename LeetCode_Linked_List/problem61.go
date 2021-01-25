package LeetCode_Linked_List

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	cur := head
	count := 0
	for cur.Next != nil {
		cur = cur.Next
		count++
	}
	count++

	cur.Next = head
	k = k % count
	n := count - k
	for n > 1 {
		head = head.Next 
		n--
	}

	newHead := head.Next
	head.Next = nil

	return newHead
}
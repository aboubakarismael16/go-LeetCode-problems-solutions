package LeetCode_Linked_List

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow,fast := head,head
	var pre *ListNode
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next

		next := slow.Next
		slow.Next = pre 
		pre = slow
		slow = next 
	}

	var left *ListNode
	var right *ListNode

	if fast.Next == nil {
		left = pre
		rgiht = slow.Next
	} else {
		right = slow.Next
		slow.Next = pre
		left = slow
	}

	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next 
		right = right.Next 
	}

	return true
}
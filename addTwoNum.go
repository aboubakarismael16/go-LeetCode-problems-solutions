package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

//second method to resolve this algorithm
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	pre :=&ListNode {0,nil}
	cur := pre
	carry := 0

	for l1 != nil || l2 != nil {
		a,b := 0,0
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}

		sum := a + b + carry
		carry = sum / 10
		sum = sum % 10

		cur.Next =&ListNode{sum,nil}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if (carry==1) {
		cur.Next = &ListNode{carry,nil}
	}
	return pre.Next

}


func main() {

	l1 := &ListNode{6,nil}
	l2 := &ListNode{5,nil}
	fmt.Println(addTwoNumbers2(l1,l2))
	
}
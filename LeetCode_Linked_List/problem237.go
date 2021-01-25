package LeetCode_Linked_List


//first method
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next 
}

//second method

//func deleteNode(node *ListNode) {
//     if node == nil {
//         return
//     }
	
//     cur := node
//     for cur.Next.Next != nil {
//         cur.Val = cur.Next.Val
//         cur = cur.Next
//     }
//     cur.Val = cur.Next.Val
//     cur.Next = nil
// }
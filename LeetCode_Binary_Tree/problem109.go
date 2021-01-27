package LeetCode_Binary_Tree


func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    
    if head.Next == nil && head != nil {
        return &TreeNode{
            Val : head.Val,
            Left : nil,
            Right : nil,
        }
    }
    
    middleNode,preNode := middleNodeAndPreNode(head)
    if middleNode == nil {
        return nil
    }
    if preNode != nil {
        preNode.Next = nil
    }
    if middleNode == head {
        head = nil
    }
    
    return &TreeNode{
        Val : middleNode.Val,
        Left : sortedListToBST(head),
        Right : sortedListToBST(middleNode.Next),
    }
    
}

func middleNodeAndPreNode(head *ListNode) (middle,pre *ListNode) {
    if head == nil || head.Next == nil {
        return nil,head
    }
    
    p1,p2 := head,head
    for p2.Next != nil && p2.Next.Next != nil {
        pre = p1
        p1 = p1.Next
        p2 = p2.Next.Next
    }
    
    return p1,pre
}

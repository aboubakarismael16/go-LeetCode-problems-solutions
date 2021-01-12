package LeetCode_Binary_Tree

func buildTree(inorder []int, postorder []int) *TreeNode {
    return buildTreeHelper(inorder,postorder)
}

func buildTreeHelper(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }
    
    rootVal := postorder[len(postorder)-1]
    index := 0
    for i,v := range inorder {
        if v == rootVal {
            index = i
        }
    }
    
    root := &TreeNode{}
    root.Val = rootVal
    root.Left = buildTreeHelper(inorder[:index] , postorder[:index] )
    root.Right = buildTreeHelper(inorder[index+1:] , postorder[index:len(postorder)-1] )
    
    return root
}

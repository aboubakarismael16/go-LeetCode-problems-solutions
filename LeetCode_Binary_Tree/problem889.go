package LeetCode_Binary_Tree

func constructFromPrePost(pre []int, post []int) *TreeNode {
    return buildTreeHelper(pre,post)
}

func buildTreeHelper(pre []int, post []int) *TreeNode {
    if len(pre) == 0 || len(post) == 0 {
        return nil
    }
    
    if len(pre) == 1 || len(post) == 1 {
        root := &TreeNode{}
        root.Val = pre[0]
        return root
    }
    
    index := 0
    for i,v := range post {
        if v == pre[1] {
            index = 1 + i
        }
    }
    
    root := &TreeNode{}
    root.Val = pre[0]
    root.Left = buildTreeHelper(pre[1:index+1],post[:index])
    root.Right = buildTreeHelper(pre[index+1:],post[index:len(post)-1])
    
    return root
}
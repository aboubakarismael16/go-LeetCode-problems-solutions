package LeetCode_Binary_Tree

type BSTIterator struct {
    values []int
    index int
}


func Constructor(root *TreeNode) BSTIterator {
    values := make([]int,0)
   
    inorder(root,&values)
    return BSTIterator{
        values : values,
        index :0,
    }
}

func inorder(root *TreeNode, values *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left,values)
    *values = append(*values,root.Val)
    inorder(root.Right,values)
}

func (this *BSTIterator) Next() int {
    if this.HasNext() {
        val := this.values[this.index]
        this.index++
        return val
    }
    
    return -1
}

func (this *BSTIterator) HasNext() bool {
    return this.index < len(this.values)
}

package LeetCode_Binary_Tree


// first method 

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	LeftToRight := true

	var q []*TreeNode
	q = append(q,root)

	for len(q) > 0 {
		lenght := len(q)
		level := []int{}

		for i := 0; i < lenght; i++ {
			node := q[0]
			q = q[1:]
            
            level = append(level,node.Val)

			if node.Left != nil {
				q = append(q,node.Left)
			}
			if node.Right != nil {
				q = append(q,node.Right)
            } 
        }

        var levelFinal []int
        if LeftToRight {
            levelFinal = level
        } else {
            for i := len(level) -1; i >= 0; i--{
                levelFinal = append(levelFinal,level[i])
            }
        }
        if len(levelFinal) > 0 {
            res = append(res,levelFinal)
        }
        LeftToRight = !LeftToRight
    }
	return res

}

// second method : recursive

func zigzagLevelOrder0(root *TreeNode) [][]int {
	res := [][]int{}
	search(root,0,&res)
	return res
}

func search(node *TreeNode, level int,res *[][]int)  {
	if node == nil {
		return
	}

	currLevel := level + 1
	for len(*res) < currLevel {
		*res = append(*res,[]int{})
	}

	if level % 2 == 0 {
		(*res)[level] = append((*res)[level],node.Val)
	} else {
		(*res)[level] = append([]int{node.Val},(*res)[level]...)
	}
	search(node.Left,currLevel,res)
	search(node.Right,currLevel,res)
}
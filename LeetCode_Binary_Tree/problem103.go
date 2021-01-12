package LeetCode_Binary_Tree

import "fmt"


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

func main() {
	l1 := TreeNode{
		Val : 3,
		Left : nil,
		Right : nil,
	}

	fmt.Println(zigzagLevelOrder(&l1))
}
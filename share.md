#LeetCode Probelms Resolutions

## LeetCode Binary Tree Search

###[problem94](https://leetcode.com/problems/binary-tree-inorder-traversal/)
```go
func inOrderTraversal(root *TreeNode) []int  {
	if root == nil {
		return nil
	}

	var res []int
	res = append(res,inOrderTraversal(root.Left)...)
	res = append(res,root.Val)
	res = append(res,inOrderTraversal(root.Right)...)
	return res
}

```

###[problem95](https://leetcode.com/problems/unique-binary-search-trees-ii/)
```go
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generateBSTree(1,n)
}

func generateBSTree(start,end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		tree = append(tree,nil)
		return tree
	}
	for i := start; i <= end; i++ {
		left := generateBSTree(start,i-1)
		right := generateBSTree(i+1,end)

		for _,l := range left {
			for _,r := range right {
				root := &TreeNode{
					Val : i,
					Left : l,
					Right :r,
				}
				tree = append(tree,root)
			}
		}
	}

	return tree
}

```

###[problem98](https://leetcode.com/problems/validate-binary-search-tree/)
```go
import "math"


func isValidBST(root *TreeNode) bool  {
	return isValidBSTHelper(root,math.MinInt32,math.MaxInt32)

}

func isValidBSTHelper(root *TreeNode,min,max int) bool {
	if root == nil {
		return true
	}

	return isValidBSTHelper(root.Left,min,root.Val-1) &&
	       isValidBSTHelper(root.Right,root.Val+1,max)
}

```

###[problem100](https://leetcode.com/problems/same-tree/)
```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return isSameTreeHelper(p,q,true)
}

func isSameTreeHelper(p *TreeNode, q *TreeNode,res bool) bool {
    if p == nil && q == nil {
        return res
    } else if (p == nil && q != nil) || (p != nil && q == nil) || p.Val != q.Val {
        return false
    }
    
    return isSameTreeHelper(p.Left,q.Left,res) && isSameTreeHelper(p.Right,q.Right,res)
}

```

###[problem101](https://leetcode.com/problems/symmetric-tree/)
```go
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }  
    if root.Left == nil && root.Right == nil {
        return true
    }
    if root.Left == nil || root.Right == nil {
        return false
    }
    if root.Left.Val != root.Right.Val {
        return false
    }
    
    var q []*TreeNode
    q = append(q,root.Left)
    q = append(q,root.Right)
    
    for len(q) > 0 {
        lenght := len(q)
        level := []int{}
        
        for i := 0; i<lenght; i++ {
            node := q[0]
            var nodeVal int
            if node != nil {
                nodeVal = node.Val
                q = append(q,node.Left)
                q = append(q,node.Right)
            }
            
            level = append(level,nodeVal)
            q = q[1:]
        }
        
        left,right := 0,len(level) - 1
        for left < right {
            if level[left] != level[right] {
                return false
            }
            left++
            right--
        }
    }
    return true
}

```

###[problem102](https://leetcode.com/problems/binary-tree-level-order-traversal/)
```go

//first method
func levelOrder(root *TreeNode) [][]int  {
	res := [][]int{}
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q,root)

	for len(q) > 0 {
		lenght := len(q)
		level := []int{}

		for i := 0 ; i < lenght; i++ {
			node := q[0]
			q = q[1:]
			level = append(level,node.Val)

			if node.Left != nil {
				q = append(q,node.Left)
			}
			if node.Right != nil {
				q = append(q,node.Right)
			}

			for len(level) > 0 {
				res = append(res,level)
			}
		}
	}

	return res
}

// second method 
func levelOrder(root *TreeNode) [][]int {
    levels := [][]int{}
    dfsLevel(root,-1,&levels) // recursive function
    return levels
}

// recursive function
func dfsLevel(node *TreeNode,level int,res *[][]int) {
    if node == nil {
        return 
    }
    
    currLevel := level + 1
    for len(*res) <= currLevel {
        *res = append(*res,[]int{})
        
    }
    (*res)[currLevel] = append((*res)[currLevel],node.Val)
    dfsLevel(node.Left,currLevel,res)
    dfsLevel(node.Right,currLevel,res)
}

```

###[problem103](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)
```go
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

```

###[problem104](https://leetcode.com/problems/maximum-depth-of-binary-tree/)
```go
func maxDepth(root *TreeNode) int {
    depth := 0
    return findMax(root,depth)
}

func findMax(root *TreeNode,depth int) int {
    if root == nil {
        return depth
    }
    
    return max(findMax(root.Left,depth + 1),findMax(root.Right,depth + 1))
}

func max(a,b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

```

###[problem105](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)
```go
func buildTree(preorder []int,inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}	

	return buildTreeHelper(preorder,inorder)
}

func buildTreeHelper(preorder []int,inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) ==  0 {
		return nil
	}

	rootVal := preorder[0]
	index := 0
	for i,v := range inorder {
		if v == rootVal {
			index = i
		}
	}

	root :=&TreeNode{}
	root.Val = rootVal
	root.Left  = buildTreeHelper(preorder[1:index+1],inorder[:index])
	root.Right = buildTreeHelper(preorder[index+1:],inorder[index+1:])

	return root
	
}

```

###[problem106](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
```go
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


```

###[problem107](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/)
```go
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q,root)
	for len(q) > 0 {
		lenght := len(q)
		level := []int{}

		for i := 0; i < lenght ; i++ {
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
		if len(level) > 0 {
			res = append(res,level)
		}
	}
	result := [][]int{}
	for i := len(res)-1; i >= 0; i-- {
		result = append(result,res[i])
	}

	return result
}


```

###[problem108](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)
```go
func sortArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		root := &TreeNode{}
		root.Val = nums[0]
		return root
	}

	mid := len(nums) / 2
	root := &TreeNode{}
	root.Val = nums[mid]
	root.Left = sortArrayToBST(nums[:mid])
	root.Right = sortArrayToBST(nums[mid+1:])
	return root
	
}



```

###[problem109](FindLongString.gohttps://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/)

``` go
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
		return nil
	}
	//           h
    //head =   [-10,-3,0,5,9]
    //                  h
    //output = [0,-3,9,-10,null,5]
	if head.Next == nil && head != nil {
		return &TreeNode {
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

func middleNodeAndPreNode(head *ListNode) (middle,pre *ListNode)  {
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

//           p1
//           p2         
//head =   [-10,-3,0,5,9]
//          pre       
//output = [0,-3,9,-10,null,5]

```


###[problem110](https://leetcode.com/problems/balanced-binary-tree/)
```go
func isBalanced(root *TreeNode) bool  {
	if root == nil {
		return true
	}

	return abs(heighHelper(root.Left,0)-heighHelper(root.Right,0)) < 2 &&
			isBalanced(root.Left) && isBalanced(root.Right)
}

func heighHelper(root *TreeNode,heigh int) int  {
	if root == nil {
		return heigh
	}

	return max(heighHelper(root.Left,heigh+1),heighHelper(root.Right,heigh+1))
	
}

func abs(a int) int  {
	if a < 0 {
		return 0 -  a 
	}
	return a
}

func max(a,b int) int {
	if  a < b {
		return b
	}
	return a
}

```

###[problem111](https://leetcode.com/problems/minimum-depth-of-binary-tree/)
```go
func minDepth(root *TreeNode) int  {
	depth := 0
	return findMin(root,depth)
}

func findMin(root *TreeNode,depth int) int  {
	if root == nil {
		return depth
	}
	if root.Left != nil && root.Right != nil {
		return min(findMin(root.Left,depth),findMin(root.Right,depth)) + 1
	} else if root.Left == nil {
		return findMin(root.Right,depth) + 1
	} else if root.Right == nil {
		return findMin(root.Left,depth) + 1
	} else {
		return depth
	}
}

func min(a,b int) int  {
	if a < b {
		return a
	}
	return b
}

```

###[problem112](https://leetcode.com/problems/path-sum/)
```go
func hasPathSum(root *TreeNode, sum int) bool {
   if root == nil {
       return false
   }

   if root.Val == sum && root.Left == nil && root.Right == nil {
       return true
   }

   return hasPathSum(root.Left,sum-root.Val) || hasPathSum(root.Right,sum-root.Val)
}

```

###[problem114](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/)
```go
func flatten(root *TreeNode)  {
	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			left := root.Left
			for left.Right != nil {
				left = left.Right
			}
			left.Right = root.Right
			root.Right = root.Left
			root.Left = nil
			root = root.Right
		}
	}
	
}

```

###[problem129](https://leetcode.com/problems/sum-root-to-leaf-numbers/)
```go
func sumNumbers(root *TreeNode) int  {
	res := 0
	if root == nil {
		return res
	}

	return sumNumberHelper(root,0)
}

func sumNumberHelper(root *TreeNode,num int) int  {
	if root == nil {
		return 0
	}
	num = root.Val + num * 10
	if root.Left == nil && root.Right == nil {
		return num 
	}

	sum := 0
	sum += sumNumberHelper(root.Left,num)
	sum += sumNumberHelper(root.Right,num)

	return sum
}

```

###[problem144](https://leetcode.com/problems/binary-tree-preorder-traversal/)
```go
func preOrderTraversal(root *treeNode) []int  {
	if root == nil {
		return nil
	}

	res := []int{}
	res = append(res,root.Val)
	res = append(res,preOrderTraversal(root.Left)...)
	res = append(res,preOrderTraversal(root.Right)...)

	return res
}


```

###[problem145](https://leetcode.com/problems/binary-tree-postorder-traversal/)
```go
func postOrderTraversal(root *TreeNode) []int  {
	if root == nil {
		return nil
	}

	var res []int
	res = append(res,postOrderTraversal(root.Left)...)
	res = append(res,postOrderTraversal(root.Right)...)
	res = append(res,root.Val)

	return res
}

```

###[problem173](https://leetcode.com/problems/binary-search-tree-iterator/)
```go
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


```

###[problem199](https://leetcode.com/problems/binary-tree-right-side-view/)
```go
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q,root)
	for len(q) > 0 {
		lenght := len(q)
		level := []int{}

		for i := 0; i<lenght; i++ {
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
		if len(level) > 0 {
			res = append(res,level[len(level)-1])
		}
	}

	return res
	
}

```

###[problem226](https://leetcode.com/problems/invert-binary-tree/)
```go
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

```

###[problem235](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
```go
func lowestCommonAncestor(root,p,q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left,p,q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right,p,q)
	}
	return root
}

```

###[problem236](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)
```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}

```

###[problem257](https://leetcode.com/problems/binary-tree-paths/)
```go
import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}

	if root == nil {
		return res
	}
	helper(root,"",&res)
	return res
}

func helper(root *TreeNode,path string,res*[]string)  {
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res,path)
		return
	}
	path += "->"
	if root.Left != nil {
		helper(root.Left,path,res)
	}
	if root.Right != nil {
		helper(root.Right,path,res)
	}
}

```


###[problem404](https://leetcode.com/problems/sum-of-left-leaves/)
```go
func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	sumOfLeftLeavesHelper(root,&res)

	return res
}

func sumOfLeftLeavesHelper(root *TreeNode,res *int)  {
	if root == nil {
		return
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res += root.Left.Val
	}

	sumOfLeftLeavesHelper(root.Left,res)
	sumOfLeftLeavesHelper(root.Right,res)
}

```


###[problem429](https://leetcode.com/problems/n-ary-tree-level-order-traversal/)
```go
type Node struct {
	Val int
	Children []*Node
}

func levelOrder(root *Node) [][]int  {
	res := [][]int{}
	if root == nil {
		return res 
	}

	var q []*Node 
	q = append(q,root)

	for len(q) > 0 {
		lenght := len(q)
		level := []int{}
		
		for i := 0; i < lenght; i++ {
			node := q[0]
			q = q[1:]
			level= append(level,node.Val)

			for _,v := range node.Children {
				q = append(q,v)
			}
		}
		if len(level) > 0 {
			res = append(res,level)
		}
	}

	return res	
}


```


###[problem450](https://leetcode.com/problems/delete-node-in-a-bst/)
```go
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == key {
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        }
        left := root.Left
        for left.Right != nil {
            left = left.Right
        }
        left.Right = root.Right
        root = root.Left 
    } else if root.Val < key {
        root.Right = deleteNode(root.Right,key)
    } else if root.Val > key {
        root.Left = deleteNode(root.Left,key)
    } 
    
    return root
}

```


###[problem501](https://leetcode.com/problems/find-mode-in-binary-search-tree/)
```go
func findMode(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	dic := make(map[int]int)
	inorder(root,&dic)
	max := 0
	for _,v := range dic {
		if v > max {
			max = v
		}
	}
	for k,v := range dic {
		if v == max {
			res = append(res,k)
		}
	}

	return res
}

func inorder(root *TreeNode, dic *map[int]int)  {
	if root == nil {
		return
	}

	inorder(root.Left,dic)
	(*dic)[root.Val]++
	inorder(root.Right,dic)
}

```


###[problem513](https://leetcode.com/problems/find-bottom-left-tree-value/)
```go
func findBottomLeftValue(root *TreeNode) int {
	res := -1
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q,root)
	for len(q) > 0 {
		lenght := len(q)
		level := []int{}

		for i := 0; i< lenght; i++ {
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

		if len(level) > 0 {
			res = level[0]
		}
	}

	return res
}

```

###[problem515](https://leetcode.com/problems/find-largest-value-in-each-tree-row/)
```go
func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q,root)
	for len(q) > 0 {
		length := len(q)
		level := []int{}

		for i := 0; i < length; i++ {
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

		maxValue := max(level)
		res = append(res,maxValue)
	}

	return res
}

func max(level []int) int  {
	if len(level) == 0 {
		return -1
	}

	res := level[0]
	for _,v := range level {
		if v > res {
			res = v
		}
	}

	return res
}

```

###[problem538](https://leetcode.com/problems/convert-bst-to-greater-tree/)
```go
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	inorder(root,&sum)
	return root
}

func inorder(root *TreeNode,sum *int)  {
	if root == nil {
		return
	}
	inorder(root.Right,sum)
	*sum = *sum + root.Val
	root.Val = *sum
	inorder(root.Left,sum)
}

```

###[problem700](https://leetcode.com/problems/search-in-a-binary-search-tree/)
```go
func searchBST(root *TreeNode,val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val < val {
		return searchBST(root.Right,val)
	} else if root.Val > val {
		return searchBST(root.Right,val)
	}

	return nil
}

```

###[problem889](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/)
```go
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

```

###[problem938](https://leetcode.com/problems/range-sum-of-bst/)
```go
func rangeSumBST(root *TreeNode,low ,high int) int {
	res := 0
	if root == nil {
		return res
	}
	if root.Val > low && root.Val > high {
		return rangeSumBST(root.Left,low,high)
	}else if root.Val < low && root.Val < high {
		return rangeSumBST(root.Right,low,high)
	}else if root.Val >= low && root.Val <= high {
		res += root.Val
		res += rangeSumBST(root.Right,low,high)
		res += rangeSumBST(root.Left,low,high)
	}

	return res
}

```


###[problem997](https://leetcode.com/problems/find-the-town-judge/)
```go
func findJudge(N int,trust [][]int) int {
	if N == 1 {
		return N
	}

	dic := make(map[int]int)
	for _,v := range trust {
		dic[v[1]]++
	}

	res := -1
	for k,v := range dic {
		if v == N-1 {
			res  = k
			break
		}
	}

	if v > -1 {
		for _,v := range trust {
			if v[0] == res {
				return -1 
			}
		}
	}

	return res
}

```

###[problem1367](https://leetcode.com/problems/linked-list-in-binary-tree/)
```go
func isSubPath(head *ListNode,root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}

	return isSubPathHelper(head,root) || isSubPath(head,root.Left) || isSubPath(head,root.Right)
}

func isSubPathHelper(head *ListNode,root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil && head != nil {
		return false
	}
	if root.Val != head.Val {
		return false
	}

	return isSubPathHelper(head.Next,root.Left) || isSubPathHelper(head.Next,root.Right) 
}

```


##LeetCode Binary Search

###[problem33](https://leetcode.com/problems/search-in-rotated-sorted-array/)

```go
func search(nums []int,target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if target >= nums[0] {
			if nums[mid] < target && nums[0] <= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1 
			}
		} else if target < nums[0] {
			if nums[mid] < target || nums[0] <= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

```

###[problem34](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

```go
func searchRange(nums []int,target int) []int {
	
	return []int{searchRangeFirst(nums,target),searchRangeLast(nums,target)}
}

func searchRangeFirst(nums []int,target int) int {
	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) {
				return mid
			}
		} else {
			right = mid - 1
		}
	}
	return -1
}

func searchRangeLast(nums []int,target int) int {
	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) {
				return mid
			}
		} else {
			right = mid - 1
		}
	}
	return -1
}

```


###[problem35](https://leetcode.com/problems/search-insert-position/)

```go
func searchInsert(nums []int,target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left,right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid -1
		}
	}

	return left
}

```


###[problem69](https://leetcode.com/problems/sqrtx/)

```go
// first method

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	left,right := 1,x
	for left <= right {
		mid = left + (right-left) / 2
		if mid == x / mid {
			return mid
		} else if mid > x / mid {
			right = mid - 1
		} else {
			left = mid +1
		}
	}
	
	return left - 1
}

//second method: Newton Iterative method

func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}

	return r
}

```


###[problem153](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)

```go
//first method
func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	left,right := 0,len(nums)-1
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid +1
		}
	}

	return nums[left]
}

//second method

func findMin2(nums []int) int {
	min := nums[0]
	for _,num := range nums[1:] {
		if min > num {
			min = num
		}
	}

	return min

```


###[problem162](https://leetcode.com/problems/find-peak-element/)

```go
func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left,right := 0,len(nums)-1
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			right = mid 
		} else {
			left = mid + 1
		}
	}

	return left
}

```


###[problem367](https://leetcode.com/problems/valid-perfect-square/)

```go
func isPerfectSquare(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}

	left,right := 0,num
	for left <= right {
		mid := left + (right-left)/2
		if mid == num/mid && num % mid == 0 {
			return true
		} else if mid == num / mid && num % mid > 0 {
			left = mid +1
		} else if mid < num / mid {
			left = mid +1
		} else if mid > num /mid {
			right = mid - 1
		}
	}

	return false
}

```


###[problem374](https://leetcode.com/problems/guess-number-higher-or-lower/)

```go
func guessNumber(n int) int {
	left,right := 1,n
	for left <= right {
		mid := left + (right-left) / 2
		res := guess(mid)

		if res == 0 {
			return mid
		} else if res == -1 {
			right = mid -1
		} else {
			left = mid +1
		}
	}
	
	return -1
}

func guess

```


###[problem704](https://leetcode.com/problems/binary-search/)

```go
func search(nums []int,target int) int  {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	
	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right - left) /2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid -1
		} else {
			left = mid + 1
		}
	}

	return -1
}

```

###[problem744](https://leetcode.com/problems/find-smallest-letter-greater-than-target/)

```go
func nextGreatesLetter(letters []byte,target []byte) byte {
	left,right := 0,len(letters)
    for left < right {
        mid := left + (right-left)/2
        if letters[mid] <= target {
            left = mid + 1
        } else if letters[mid] > target {
            right = mid
        }
    }
    
    if right >= len(letters) {
        return letters[0]
    }
    return letters[right]
}

```

## LeetCode Array

###[problem001](https://leetcode.com/problems/two-sum/)

```go
func TwoSum(nums []int,target int) []int  {
	res :=[]int{-1,-1}

	lookup := make(map[int]int)
	for i, num := range nums {
		temp := target - num 
		if _,ok := lookup[temp]; ok {
			res[0] = lookup[temp]
			res[1] = i
			return res
		}
		lookup[num] = i
	}
	return res
	
}

```

###[problem066](https://leetcode.com/problems/plus-one/)

```go
func plusOne(digits []int) []int {
	res := make([]int,0)
	if digits == nil || len(digits) == 0 {
		return res
	}

	carry := 1
	for i := len(digits)-1; i >= 0; i-- {
		sum := digits[i] + carry
		carry = sum / 10
		digits[i] = sum % 10 
	}
	if carry == 1 {
		res = []int{1}
		digits = append(res,digits...)
	}

	return digits
}

```

###[problem121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

```go
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	minPrice,maxProfit := prices[0],0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i]-minPrice > maxProfit) {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

```

###[problem122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)

```go
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}

```

###[problem169](https://leetcode.com/problems/majority-element/)

```go
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	threshold := 1 + len(nums)/2
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val,ok := m[nums[i]]
		if ok {
			val++
			if val >= threshold {
				return nums[i]
			} else {
				m[nums[i]] = val
			}
		} else {
			m[nums[i]] = 1
		}
	}
	return 0
}

// second method : O(n)

func majorityElement1(nums []int) int {
	m := make(map[int]int)
	for _,v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v 
		}
	}

	return 0
}

```

###[problem389](https://leetcode.com/problems/find-the-difference/)

```go
func findTheDifference(s string,t string) byte {
	var res byte
	if len(s) >= len(t) {
		return res
	}
	S := []byte(s)
	T := []byte(t)

	dic := make(map[byte]int)
	for i := 0; i <len(S); i++ {
		dic[S[i]]++
	}
	for i := 0; i < len(T); i++ {
		val,ok := dic[T[i]]
		if !ok {
			res = T[i]
			break
		} else {
			if val == 0 {
				res = T[i]
				break
			}
			dic[T[i]]--
		}
	}
	return res
}

```







## LeetCode Linked List

###[problem002](https://leetcode.com/problems/add-two-numbers/)

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{0,nil}
	cur := pre
	carry := 0
	

	for l1 != nil || l2 != nil {
		a,b := 0,0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum := a + b + carry
		carry = sum / 10
		sum = sum % 10

		cur.Next = &ListNode{sum,nil}
		cur = cur.Next

		// if l1 != nil {
		// 	l1 = l1.Next
		// }
		// if l2 != nil {
		// 	l2 =l2.Next
		// }
	}

	if carry == 1 {
		cur.Next = &ListNode{carry,nil}
	}

	return pre.Next

}

```

###[problem19](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{0,nil}
	pre.Next = head
	left := pre
	right := pre

	for i := 0; i < n+1; i++ {
		right = right.Next
	}
	for right != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next

	return pre.Next
}

```

###[problem21](https://leetcode.com/problems/merge-two-sorted-lists/)

```go

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{0,nil}
	cur := pre

	for l1 != nil && l2 != nil {
		temp := l1.Val
		if l1.Val > l2.Val {
			temp = l2.Val
		}
		cur.Next = &ListNode{temp,nil}
		cur = cur.Next

		if l1.Val > l2.Val {
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return pre.Next
}

```

###[problem23](https://leetcode.com/problems/merge-k-sorted-lists/)

```go

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])

	return Helper(left,right)
}

func Helper(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = Helper(l1.Next,l2)
		return l1
	}
	l2.Next = Helper(l1,l2.Next)
	return l2
}

```

###[problem24](https://leetcode.com/problems/swap-nodes-in-pairs/)

```go

func swapPairs(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}

	dummy := &ListNode{-1,head}
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		nextNext := next.Next

		pre.Next = next
		next.Next = cur
		cur.Next = nextNext

		pre = cur
		cur = cur.Next
	}

	return dummy.Next
}

```

###[problem25](https://leetcode.com/problems/reverse-nodes-in-k-group/submissions/)

```go

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next 
	} 
	newHead := reverse(head,node)
	head.Next = reverseKGroup(node,k)

	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	pre := last
	for first != last {
		temp := first.Next
		first.Next = pre
		pre = first
		first = temp
	}

	return pre
}

```

###[problem61](https://leetcode.com/problems/rotate-list/)

```go

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

```

###[problem82](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/)

```go

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
    dummy := &ListNode{-1,head}
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			cur = cur.Next
            pre = pre.Next
            continue
		} 
        for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
            cur = cur.Next
        }
        cur = cur.Next
        pre.Next = cur
	}

	return dummy.Next
}

```

###[problem83](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)

```go

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pre := &ListNode{-1,head}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next 
		}
	}

	return pre.Next
}

```

###[problem141](https://leetcode.com/problems/linked-list-cycle/)

```go

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow,fast := head,head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

```

###[problem142](https://leetcode.com/problems/linked-list-cycle-ii/)

```go

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow,fast := head,head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}	
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	third := head
	for third != slow {
		slow = slow.Next
		third = third.Next
	}

	return slow
}

```

###[problem]()

```go

```


###[problem]()

```go

```


###[problem]()

```go

```


###[problem]()

```go

```



###[problem328](https://leetcode.com/problems/odd-even-linked-list/)

```go

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd,even,evenHead := head,head.Next,head.Next
	for even != nil && even.Next != nil {
		oddNext := even.Next 
		odd.Next = oddNext
		odd = odd.Next 

		evenNext := odd.Next 
		even.Next = evenNext 
		even = even.Next 
	}
	odd.Next = evenHead

	return head
}

// h
// o, e, o, e, o, e, o
// 2->1->3->5->6->4->7->NULL
// 
//2->3->6->7->1->5->4->NULL

```


## LeetCode String

###[problem14](https://leetcode.com/problems/longest-common-prefix/)

```go

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	res := strs[0]
	for i := 1; i < len(strs); i++ {
		resLenght := len(res)
		tempLength := len(strs[i])
		length := min(resLenght,tempLength)
		j := 0
		for j < length {
			if res[j] != strs[i][j] {
				break
			}
			j++
		}
		res = res[:j]
		if len(res) == 0 {
			return ""
		}
	}
	
	return res
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

```


###[problem58](https://leetcode.com/problems/length-of-last-word/)

```go
func LengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	chars := []byte(s)
	first := -1
	for i := len(s)-1; i >= 0; i-- {
		if first == -1 {
			if chars[i] == ' ' {
				continue
			} else {
				first = i 
			}
		} else {
			if chars[i] == ' ' {
				return first - i 
			}
		}
	}

	return first + 1
}

// "Hello World"
// "Hello World        "
// "a"
// "      "
// "  wwe  "
```

###[problem383](https://leetcode.com/problems/ransom-note/)

```go
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	r := []byte(ransomNote)
	m := []byte(magazine)

	dic := make(map[byte]int)
	for i := 0; i < len(m); i++ {
		dic[m[i]]++
	}

	for i := 0; i < len(r); i++ {
		val,ok := dic[r[i]]
		if ok {
			if val == 0 {
				return false 
			} else {
				dic[r[i]]--
			}
		} else {
			return false
		}
	}

	return true
}
```

###[problem1374](https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/)

```go
func generateTheString(n int) string {
	if n == 0 {
		return ""
	}
	if n % 2 == 0 {
		// pair
		return buildString(n-1) + "b"
	} else {
		//impair
		return buildString(n)
	}
}

func buildString(n int) string  {
	res := ""
	for i := 0; i < n ; i++ {
		res += "a"
	}

	return res 
}
```

###[problem1422](https://leetcode.com/problems/maximum-score-after-splitting-a-string/)

```go
func maxScore(s string) int {
	res := 0
	if len(s) == 0 {
		return 0
	}

	left := 0
	if s[0] == '0' {
		left = 1
	}

	right := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			right++
		}
	}

	res = left + right
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			left++
		}
		if s[i] == '1' {
			right--
		}
		temp := left + right
		if temp > res {
			res = temp
		}
	}

	return res
}
```


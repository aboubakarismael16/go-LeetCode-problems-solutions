package LeetCode

func FindRoot(p int, arr []int) int {
    for p != arr[p] {
        p = arr[p]
    }
    return p
}

func Union(p, q int, arr []int) {
    arr[q] = p
}

func findRedundantConnection(edges [][]int) []int {
    arr := make([]int, len(edges) + 1)
    for index := range arr {
        arr[index] = index
    }
    for _, edge := range edges {
        pRoot := FindRoot(edge[0], arr)
        qRoot := FindRoot(edge[1], arr)
        if pRoot == qRoot {
            return edge
        } else {
            Union(pRoot, qRoot, arr)
        }
    }
    return nil
}
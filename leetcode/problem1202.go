package LeetCode


func smallestStringWithSwaps(s string, pairs [][]int) string {
    uf := NewUF(len(s))
    for _,pair := range pairs {
        uf.Union(pair[0],pair[1])
    }
    dic := map[int][]byte{}
    for i := 0; i < len(s); i++ {
        r := uf.Find(i)
        dic[r] = append(dic[r],s[i])
    }
    for _,v := range dic {
        sort.Slice(v,func(i,j int) bool {
            return v[i] < v[j]
        })
    }
    res := []byte(s)
    for i := 0; i < len(s); i++ {
        r := uf.Find(i)
        bytes := dic[r]
        res[i] = bytes[0]
        dic[r] = bytes[1:len(bytes)]
    }
    return string(res)
}


type UF struct {
    arr []int
}


func NewUF(size int) *UF {
    
    var arr []int
    for i :=0; i < size; i++ {
        arr = append(arr,i)
    }
    
    return &UF{arr}
}

func (u *UF) Union(a,b int) {
    
    pa,pb := u.Find(a), u.Find(b)
    if pa != pb {
        u.arr[pa] = pb
    }
}

func (u *UF) Find(n int) int {
    if n!= u.arr[n] {
        u.arr[n] = u.Find(u.arr[n])
    }
    return u.arr[n]
}
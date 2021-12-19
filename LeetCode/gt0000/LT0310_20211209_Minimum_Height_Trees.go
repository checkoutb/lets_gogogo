// package sdq
package main

import (
    "fmt"
)



// .. 树， 叶子结点是度为1的结点， 度为1的结点不可能是 root， 因为使用 度为1的结点的父结点 作为root，深度必然 <= 叶子结点作为root

// 循环 减去 度1的叶子。 直到 只剩2个。


// root 是 最长的path 的中点。。。。。。



// Runtime: 452 ms, faster than 5.13% of Go online submissions for Minimum Height Trees.
// Memory Usage: 8.3 MB, less than 35.90% of Go online submissions for Minimum Height Trees.
// 只有硬算吧。
// hint1, 感觉最多2个， 所以是任意找一个结点，然后 计算 哪个子树最深， 用最深的这个子树的根 作为根 再算一遍？ 当所有的子树都是同一深度或者都计算过了，就return。
func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 0 {
        return []int{0}
    }
    mndepth := 1000000
    map2 := map[int]int{}
    map3 := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        t2, t3 := edges[i][0], edges[i][1]
        map3[t2] = append(map3[t2], t3)
        map3[t3] = append(map3[t3], t2)
    }

    q := []int{0}

    for len(q) > 0 {
        // fmt.Println(q)
        szq := len(q)
        for i := 0; i < szq; i++ {
            mxdp := 0
            tmp := []int{}
            for _, v := range map3[q[i]] {
                t2 := dfsa310(v, map3, q[i])
                if t2 + 1 > mxdp {
                    mxdp = t2 + 1
                    tmp = []int{}
                    if _, ok := map2[v]; !ok {
                        tmp = append(tmp, v)
                    }
                } else if t2 + 1 == mxdp {
                    tmp = append(tmp, v)
                }
            }
            q = append(q, tmp...)
            if mxdp < mndepth {
                mndepth = mxdp
            }
            map2[q[i]] = mxdp
        }
        q = q[szq : ]
    }

    ans := []int{}
    for k, v := range map2 {
        if v == mndepth {
            ans = append(ans, k)
        }
    }
    // fmt.Println(map2)
    return ans
}

func dfsa310(node int, map3 map[int][]int, parent int) int {
    ans := 1
    for _, child := range map3[node] {
        if child != parent {
            t2 := dfsa310(child, map3, node)
            if t2 + 1 > ans {
                ans = t2 + 1
            }
        }
    }
    return ans
}


func main_LT0310_20211209() {
// func main() {

    fmt.Println("ans:")

    // n := 4
    // arr := [][]int{{1,0},{1,2},{1,3}}

    n := 2
    arr := [][]int{{1,0}}

    fmt.Println(findMinHeightTrees(n, arr))

}

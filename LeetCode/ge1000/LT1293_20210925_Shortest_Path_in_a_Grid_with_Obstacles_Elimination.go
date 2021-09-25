
package main

import (
    "fmt"
    "container/list"
)

// D D

// 得想办法去哪里弄个utils来。

// queue := [][]int{{0,0}, nil}         // nil 是第二个元素。。 应该是用来 作为 step++ 的条件。

// m,n  := len(grid), len(grid[0])
// leftK := make([][]int, m)            // s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
// for i := range leftK {
//     leftK[i] = make([]int, n)
//     for j := range leftK[i] {
//         leftK[i][j] = -1
//     }
// }

// for len(queue) > 0 {
//     q := queue[0]
//     queue = queue[1:]
    // if q == nil {
    //     step++  

// queue = append(queue, qn)




// Runtime: 39 ms, faster than 58.82% of Go online submissions for Shortest Path in a Grid with Obstacles Elimination.
// Memory Usage: 7.2 MB, less than 82.35% of Go online submissions for Shortest Path in a Grid with Obstacles Elimination.
func shortestPath(grid [][]int, k int) int {
    ans := 0
    sz1, sz2 := len(grid), len(grid[0])
    // vvm := [][]map[int]int{}
    // vvm := [2][2]map[int]int{}
    // vvm := make([2][2]int, sz1)
    // vector<vector<int>> ...

    // var vvm [][]map[int]int
    // for i := 0; i < sz1; i++ {
    //     t2 := make([]map[int]int, sz2)
    //     vvm = append(vvm, t2)
    // }

    vvm := make([][]map[int]int, sz1)
    for i := range vvm {
        vvm[i] = make([]map[int]int, sz2)
        for j := range vvm[i] {
            vvm[i][j] = make(map[int]int)
        }
    }

    // fmt.Println(vvm[0][0][1])
    // fmt.Println(len(vvm[0][0]))

    // bfs or dfs
    // 没有queue，。。 package list...  

    que1 := list.New()
    // que1.PushBack([]int{0,0})               // 首字母大写的驼峰。。  []int{0,0} is object (interface{}'s instance)...cao,bushi,其他地方报错了，导致这里就没有检查。。或者说 这种是对的，只不过拿不出来？.buzhidao
    // map2 := vvm[0][0]
    // map2[k] = 0
    vvm[0][0][k] = 0        // kao, meiyou init...   nil...
    que1.PushBack(makeABCD(0, 0, k, 0))

    // for list.Len(que1) > 0 {}        // undefined: list.Len
    // func New() *List     func (l *List) Len() int        第一个形参 是 对象自己？ 怎么区分 这个 是 对象方法 还是 静态方法？
                // 好像没有 静态方法， 连类都没有。。  反正就是怎么区分。。。==， New 该不会是 关键方法吧。。 不是。 new() 是内置函数
            // * is pointer.
            // 看错了 。。。 Len 的形参是空的， (l *List)是调用者。。。整天倒装，怎么不把形参列表也倒装下。。。

    arr := [][]int{{0,1},{1,0},{-1,0},{0,-1}}

    for que1.Len() > 0 {
        // fmt.Println(que1.Len())
        // var arr2 [2]int = que1.Front().Value         // [2]int 不是 interface{}。
        // 感觉不如自己写链表
        // t1 := arr2[0]
        // t2 := arr2.Value[1]
        // var t5 int = que1.Front().Value          // int is not interface{} 。。。 
        // t1, t2 := t5 / 100, t5 % 100
        // que1.remove(t5)

        // asd := que1.Front().Value.([]int)       /// .... vscode bu bao cuo

        var t4 ABCD = que1.Front().Value.(ABCD)      // cannot use que1.Front().Value (type interface {}) as type AB in assignment: need type assertion
                                        // 是不是上面也是这个原因。。 yinggai shide ...
        t1, t2 := t4.a, t4.b
        z1, z2 := t4.c, t4.d
        que1.Remove(que1.Front())       // 。。。 感觉好像无效啊，死循环啊。或者 代码逻辑错了。 写错了， 剩余翻转次数 和 步数 没有写。

        for _, ele := range arr {
            a1, a2 := t1 + ele[0], t2 + ele[1]
            if a1 >= 0 && a2 >= 0 && a1 < sz1 && a2 < sz2 {
                c1, c2 := z1, z2 + 1
                if grid[a1][a2] == 1 {
                    c1--
                }
                if (c1 < 0) {
                    continue
                }
                var map2 *map[int]int = &vvm[a1][a2]            // 加不加& 好像都可以。。     & 导致 下面需要 *
                for k, v := range *map2 {        // 还剩多少次翻转， 步数
                    if (k >= c1 && v <= c2) {       // 。。哈哈，用坐标 和 翻转次数，步数 在做比较。。
                        goto AAA;
                    }
                }
                // que1.PushBack([]int{a1, a2})
                vvm[a1][a2][c1] = c2
                que1.PushBack(makeABCD(a1, a2, c1, c2))
                AAA:
                continue
            }
        }
    }
    ans = 10000000
    for _, v := range vvm[sz1 - 1][sz2 - 1] {           // no &, no *
        if ans > v {
            ans = v
        }
    }
    if ans < 10000000 {
        return ans
    } else {
        return -1
    }
}
type ABCD struct {
    a int
    b int
    c int
    d int
}
func makeABCD(a int, b int, c int, d int) ABCD {
    // return AB{a: a, b: b};      // ...buxiebuzhidao,
    return ABCD{a, b, c, d}
}

// func makeABP(a2 int, b2 int) *AB {
//     ab := new(AB)           // *AB
//     // var ab = AB{a: a, b: b}
//     ab.a = a2
//     ab.b = b2
//     return ab
// }
// arr := {{0,1},{1,0},{-1,0},{0,-1}};      function 外不能声明  [][]int{xxxx}


func main_LT1293_20210925() {
// func main() {

    // var vvi [][]int = [][]int{{0,0,0},{1,1,0},{0,0,0},{0,1,1},{0,0,0}};
    // k := 1

    vvi := [][]int{{0,1,1},{1,1,0},{1,0,0}}
    k := 1

    ans := shortestPath(vvi, k)

    fmt.Println("ans:")

    fmt.Println(ans)
}

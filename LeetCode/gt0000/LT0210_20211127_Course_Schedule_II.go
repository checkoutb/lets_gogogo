// package sdq
package main

import (
    "fmt"
)


// for i := 0; i != len(out); i++ {
//     node := out[i]
//     for _,v := range vectors[node]{
//         ins[v] --
//         if ins[v] == 0 {
//             out = append(out, v)         // .
//         }
//     }
// }
// queue, 数组元素会增加。


// Runtime: 12 ms, faster than 79.20% of Go online submissions for Course Schedule II.
// Memory Usage: 6.4 MB, less than 68.98% of Go online submissions for Course Schedule II.
func findOrder(numCourses int, prerequisites [][]int) []int {
    indegree := make([]int, numCourses)
    map2 := map[int][]int{}
    for _, v := range prerequisites {
        indegree[v[0]]++
        if _, ok := map2[v[1]]; !ok {
            map2[v[1]] = []int{}
        }
        map2[v[1]] = append(map2[v[1]], v[0])
    }
    // ans := make([]int, numCourses)
    ans := []int{}
    chg := true
    for chg {
        chg = false
        for i, v := range indegree {
            if v == 0 {
                chg = true
                ans = append(ans, i)
                if _, ok := map2[i]; ok {
                    for _, nxt := range map2[i] {
                        indegree[nxt]--
                    }
                }
                indegree[i] = -1
            }
        }
    }
    if len(ans) == numCourses {
        return ans
    } else {
        return []int{}
    }
}


func main_LT0210_20211127() {
// func main() {

    fmt.Println("ans:")


}

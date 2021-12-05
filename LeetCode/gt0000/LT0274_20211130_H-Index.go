// package sdq
package main

import (
    "fmt"
    "sort"
)



// for l < r {
//     mid := (l + r) >> 1
//     if citations[n-1-mid] > mid {
//         l = mid + 1
//     } else {
//         r = mid
//     }
// }



// Runtime: 0 ms, faster than 100.00% of Go online submissions for H-Index.
// Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for H-Index.
// n份论文中 h份 >=h次引用， n-h份 低于h次引用
// 二分 不了。  多个满足的值。
func hIndex(citations []int) int {
    sort.Ints(citations)
    for i := 0; i < len(citations); i++ {
        cnt := len(citations) - i
        if citations[i] >= cnt {
            return cnt
        }
    }
    return 0

    // ans := -1
    // for i := len(citations) - 1; i >= 0; i-- {
    //     cnt := len(citations) - i
    //     if citations[i] >= cnt {
    //         ans = cnt
    //     }
    // }
    // return ans

    // l, r := 0, len(citations) - 1
    // for l < r {
    //     mid := (l + r) / 2          // fengshu
    //     t2 := citations[mid]        // yy cishu
    //     fmt.Println(l, mid, r, t2)
    //     if t2 >= mid {
    //         r = mid
    //     } else {
    //         l = mid + 1
    //     }
    // }
    // return citations[l]
}


func main_LT0274_20211130() {
// func main() {

    fmt.Println("ans:")

    arr := []int{3,0,6,1,5}
    // arr := []int{1,3,1}

    fmt.Println(hIndex(arr))

}

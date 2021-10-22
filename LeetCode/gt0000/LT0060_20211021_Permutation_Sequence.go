// package sdq
package main

import (
    "fmt"
    "strconv"
)





// gg .... 

// 多少个数。
func getPermutation(n int, k int) string {
    arr := make([]int, n + 1)           // 1...n      0==not used
    arr2 := make([]int, n + 1)      // 1, 1*2, 3!, 4! ... 9!
    ans := ""
    arr2[1] = 1
    for i := 2; i <= n; i++ {
        arr2[i] = arr2[i - 1] * i
    }
    arr2[0] = 1
    fmt.Println(arr2)
    for i := n - 1; i >= 0; i-- {
        if k >= arr2[i] {
            t2 := k / arr2[i] + 1
            for j := 1; j <= n; j++ {
                if arr[j] == 0 {
                    t2--
                    if t2 == 0 {
                        ans += strconv.Itoa(j)
                        arr[j] = 1
                        break
                    }
                }
            }
        } else {
            for j := 1; j <= n; j++ {
                if arr[j] == 0 {
                    ans += strconv.Itoa(j)
                    arr[j] = 1
                    break
                }
            }
        }
        // if k >= arr2[i] {
        //     t2 := k / arr2[i] + 1
        //     // if k % arr2[i] == 0 {
        //     //     t2++
        //     // }
        //     k -= (t2 - 1) * arr2[i]
        //     // if k > 0 || t2 == 0 {
        //     //     t2++
        //     // }
        //     // fmt.Println(t2, ", ", k)
        //     for j := 1; j <= n; j++ {
        //         if arr[j] == 0 {
        //             t2--
        //             if t2 == 0 {
        //                 ans += strconv.Itoa(j)
        //                 arr[j] = 1
        //                 break
        //             }
        //         }
        //     }
        // }
    }
    for j := 1; j <= n; j++ {
        if arr[j] == 0 {
            ans += strconv.Itoa(j)
            // break
        }
    }
    return ans
}


func main_LT0060_20211021() {
// func main() {

    fmt.Println("ans:")

    // n, k := 3, 3     // 213
    n, k := 4, 9     // 2314
    // n, k := 3, 1     // 123
    // n, k := 2, 2     // 21

    ans := getPermutation(n, k)

    fmt.Println(ans)

}

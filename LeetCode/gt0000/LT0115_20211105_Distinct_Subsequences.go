// package sdq
package main

import (
    "fmt"
)






// error

// “” == “”

// dp[len(s)][len(t)]
// s[i] == t[j] -> dp[i][j] = dp[i-1][j-1]  + ???          [0][x] = 1
//      else  max(1,2,3 ? )
// 到 [i][j] 为止，  i 可能取不到的，  j必须的，因为是为了生成 j。
func numDistinct(s string, t string) int {
    sz1, sz2 := len(s), len(t)
    arr := make([][]int, sz1)
    for i, _ := range arr {
        arr[i] = make([]int, sz2)
    }
    for i := 0; i < sz1; i++ {
        if s[i] == t[0] {
            if i == 0 {
                arr[i][0] = 1
            } else {
                arr[i][0] = 1 + arr[i - 1][0]
            }
        }
    }
    for i := 1; i < sz1; i++ {              // 1 - 0
        for j := 1; j < sz2; j++ {          // 1 - 0
            if s[i] == t[j] {
                arr[i][j] = arr[i - 1][j - 1] + arr[i][j - 1]           // [i][j - 1] 是指 s取到， t没有取到。  错了， discuss是 多一行的那种， 所以是 本次相同，那么可以 影响到 下一次的。
                // if i > 0 {
                    // arr[i][j] = arr[i - 1][j - 1] + arr[i - 1][j]
                // } else {
                //     arr[i][j] = arr[i - 1][j]
                // }
            } else {
                arr[i][j] = arr[i][j - 1]
                // if i > 0 {
                //     arr[i][j] = arr[i - 1][j]
                // }
            }
        }
    }

    for _, v := range arr {
        fmt.Println(v)
    }

    return arr[sz1 - 1][sz2 - 1]
}


func main_LT0115_20211105() {
// func main() {

    fmt.Println("ans:")

    // s, t := "rabbbit", "rabbit"
    s, t := "babgbag", "bag"

    ans := numDistinct(s, t)

    fmt.Println(ans)

}

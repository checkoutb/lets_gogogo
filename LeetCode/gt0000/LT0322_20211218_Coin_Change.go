// package sdq
package main

import (
    "fmt"
    "sort"
)


// for _, coin := range coins {
//     for x := coin; x < len(dp); x++ {
//         dp[x] = Min(dp[x], 1 + dp[x - coin])
//     }
// }

// for (int c : coins)
//     for (int a=c; a<=amount; a++)
//         need[a] = min(need[a], need[a-c] + 1);
// need 初始全部 amount + 1
// 感觉这样有点慢。因为感觉 need[a-c] 大概率 是 amount+1 。。。这个要看coin， 有1的话，就不是了。

// Runtime: 8 ms, faster than 94.12% of Go online submissions for Coin Change.
// Memory Usage: 6.5 MB, less than 97.79% of Go online submissions for Coin Change.
// amount < 10^4
// coin INT_MAX. int64..
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    sort.Ints(coins)
    arr := make([]int, amount + 1)
    for _, v := range coins {
        if v <= amount {
            arr[v] = 1
        }
    }
    for i := 1; i <= amount; i++ {
        if arr[i] != 0 {
            for _, v := range coins {
                if i + v <= amount {
                    if arr[i + v] == 0 {
                        arr[i + v] = arr[i] + 1
                    } else if arr[i] + 1 < arr[i + v] {
                        arr[i + v] = arr[i] + 1
                    }
                } else {
                    break
                }
            }
        }
    }
    if arr[amount] != 0 { return arr[amount] } else { return -1 }
}


func main_LT0322_20211218() {
// func main() {

    fmt.Println("ans:")

    arr := []int{474,83,404,3}
    t2 := 264

    ans := coinChange(arr, t2)

    fmt.Println(ans)

}

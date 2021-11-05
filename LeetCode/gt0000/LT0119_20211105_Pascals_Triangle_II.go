// package sdq
package main

import (
    "fmt"
)


// for i := 1; i < rowIndex+1; i++ {
//     for j := i; j > 0; j-- {
//         dp[j]+=dp[j-1]
//     }
// }
// 错位 相加。


// for (int j = 1; j <= rowIndex - 2; j++)
// {
//     result.add((int)(s = (rowIndex - j) * s / j));
// }
// 这个是哪里弄的？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？ 


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle II.
// Memory Usage: 2 MB, less than 74.01% of Go online submissions for Pascal's Triangle II.
// 0 index
func getRow(rowIndex int) []int {
    ans := []int{1}
    for rowIndex > 0 {
        rowIndex--
        sz1 := len(ans)
        arr := make([]int, sz1 + 1)
        arr[0], arr[sz1] = 1, 1
        for i := 1; i < sz1; i++ {
            arr[i] = ans[i] + ans[i - 1]
        }
        ans = arr
    }
    return ans
}


func main_LT0119_20211105() {
// func main() {

    fmt.Println("ans:")


}

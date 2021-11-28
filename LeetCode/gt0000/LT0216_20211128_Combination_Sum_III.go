// package sdq
package main

import (
    "fmt"
)





// Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum III.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Combination Sum III.
func combinationSum3(k int, n int) [][]int {
    ans := [][]int{}
    dfsa216(n, 0, k, &ans, []int{})
    return ans
}

func dfsa216(remain, lst, cnt int, ans *[][]int, arr []int) {
    if cnt == 0 || remain < 0 {
        if remain == 0 {
            arr2 := make([]int, len(arr))
            copy(arr2, arr)
            *ans = append(*ans, arr2)
        }
        return
    }

    // if remain == 0 && cnt == 0 {
    //     *ans = append(*ans, arr)
    //     return
    // }
    // if cnt == 0 {
    //     return
    // }
    for num := lst + 1; num <= (10 - cnt); num++ {
        arr = append(arr, num)
        dfsa216(remain - num, num, cnt - 1, ans, arr)
        arr = arr[0 : len(arr) - 1]
    } 
}


func main_LT0216_20211128() {
// func main() {

    fmt.Println("ans:")


}

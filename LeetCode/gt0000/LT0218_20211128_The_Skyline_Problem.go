// package sdq
package main

import (
    "fmt"
)


// gg


// 高度, 结束
// 后续 且在 right内 第一个高于它的值。 如果right内没有比它高的，那么 它的right 就是 下一个点。
func getSkyline(buildings [][]int) [][]int {
    stk := [][]int{ []int{buildings[0][0], buildings[0][2]} }
    ans := [][]int{}
    sz1 := len(buildings)
    deal := 0
    // ans = append(ans, []int{buildings[0][0], buildings[0][2]})
    // for i := 1; i < sz1; i++ {
    for deal < sz1 {
        for j := deal + 1; j < sz1; j++ {
            if buildings[j][0] < buildings[deal][1] {
                if buildings[j][2] > buildings[deal][2] {
                    deal = j
                    break
                }
            } else {
                break
            }
        }
    }
    return ans
}

func main_LT0218_20211128() {
// func main() {

    fmt.Println("ans:")


}

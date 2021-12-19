// package sdq
package main

import (
    "fmt"
)






// Runtime: 700 ms, faster than 44.74% of Go online submissions for Range Sum Query 2D - Immutable.
// Memory Usage: 18.1 MB, less than 20.18% of Go online submissions for Range Sum Query 2D - Immutable.
type NumMatrix struct {
    Arr [][]int
}
func Constructor(matrix [][]int) NumMatrix {
    sz1, sz2 := len(matrix), len(matrix[0])
    arr := make([][]int, sz1)
    for i, _ := range arr {
        arr[i] = make([]int, sz2)
    }
    arr[0][0] = matrix[0][0]
    for i := 1; i < sz1; i++ {
        arr[i][0] = arr[i - 1][0] + matrix[i][0]
    }
    for j := 1; j < sz2; j++ {
        arr[0][j] = arr[0][j - 1] + matrix[0][j]
    }
    for i := 1; i < sz1; i++ {
        for j := 1; j < sz2; j++ {
            arr[i][j] = arr[i - 1][j] + arr[i][j - 1] - arr[i - 1][j - 1] + matrix[i][j]
        }
    }
    return NumMatrix{ Arr : arr }
}
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    ans := this.Arr[row2][col2]
    if row1 > 0 {
        ans -= this.Arr[row1 - 1][col2]
    }
    if col1 > 0 {
        ans -= this.Arr[row2][col1 - 1]
    }
    if col1 > 0 && row1 > 0 {
        ans += this.Arr[row1 - 1][col1 - 1]
    }
    return ans
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main_LT0304_20211206() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
    "fmt"
)




// for (int i = 0; i < matrix.length / 2; i++) {
//     for (int j = i; j < matrix.length - i - 1; j++) {


    // reverse(matrix.begin(), matrix.end());
    // for (int i = 0; i < matrix.size(); ++i) {
    //     for (int j = i + 1; j < matrix[i].size(); ++j)
    //         swap(matrix[i][j], matrix[j][i]);
    // }
// 先上下， 然后 按主对角线 swap。


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate Image.
// Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Rotate Image.
// 顺时针 90度
func rotate(matrix [][]int)  {
    sz1 := len(matrix)
    mxi, mxj := sz1 / 2, (sz1 + 1) / 2
    for i := 0; i < mxi; i++ {
        for j := 0; j < mxj; j++ {
            a1,b1 := i, j
            a2,b2 := j, sz1 - 1 - i
            a3,b3 := sz1 - 1 - i, sz1 - 1 - j
            a4,b4 := sz1 - 1 - j, i
            matrix[a1][b1], matrix[a2][b2], matrix[a3][b3], matrix[a4][b4] = matrix[a4][b4], matrix[a1][b1], matrix[a2][b2], matrix[a3][b3]
        }
    }
}


func main_LT0048_20211019() {
// func main() {

    fmt.Println("ans:")

    // arr := [][]int{{1,2,3},{4,5,6},{7,8,9}}
    arr := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}

    rotate(arr)

    for _, ar := range arr {
        for _, e := range ar {
            fmt.Print(e, ", ")
        }
        fmt.Println()
    }

}

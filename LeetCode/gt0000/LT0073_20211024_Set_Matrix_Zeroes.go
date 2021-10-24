// package sdq
package main

import (
    "fmt"
    "sort"
)


// 使用第一行，第一列 来记录 这列，这行 是否为0。
// [0,0]这个点 重叠了 行与列的信息， 所以需要一个 额外的var 来保存。
// int col0 = 1, rows = matrix.size(), cols = matrix[0].size();
// for (int i = 0; i < rows; i++) {
//     if (matrix[i][0] == 0) col0 = 0;
//     for (int j = 1; j < cols; j++)
//         if (matrix[i][j] == 0)
//             matrix[i][0] = matrix[0][j] = 0;
// }
// for (int i = rows - 1; i >= 0; i--) {
//     for (int j = cols - 1; j >= 1; j--)
//         if (matrix[i][0] == 0 || matrix[0][j] == 0)
//             matrix[i][j] = 0;
//     if (col0 == 0) matrix[i][0] = 0;
// }



// Runtime: 12 ms, faster than 87.77% of Go online submissions for Set Matrix Zeroes.
// Memory Usage: 6.3 MB, less than 99.69% of Go online submissions for Set Matrix Zeroes.
func setZeroes(matrix [][]int)  {
    cols := []int{}
    rows := []int{}
    sz1, sz2 := len(matrix), len(matrix[0])
    for i := 0; i < sz1; i++ {
        for j := 0; j < sz2; j++ {
            if matrix[i][j] == 0 {
                rows = append(rows, i)
                cols = append(cols, j)
            }
        }
    }
    sort.Ints(cols)

    for i := 0; i < len(cols); i++ {
        if i == 0 {
            for k := 0; k < sz1; k++ {
                matrix[k][cols[i]] = 0
            }
            for k := 0; k < sz2; k++ {
                matrix[rows[i]][k] = 0
            }
        } else {
            if cols[i] != cols[i - 1] {
                for k := 0; k < sz1; k++ {
                    matrix[k][cols[i]] = 0
                }
            }
            if rows[i] != rows[i - 1] {
                for k := 0; k < sz2; k++ {
                    matrix[rows[i]][k] = 0
                }
            }
        }
    }

}

// func dfsa73(matrix [][]int, cols []int, idx int) {
//     if idx == len(matrix) {
//         break
//     }
//     b2 := false
//     for j := 0; j < len(matrix[0]); j++ {
//         if matrix[idx][j] == 0
//     }
// }

func main_LT0073_20211024() {
// func main() {

    fmt.Println("ans:")

    arr := [][]int{{1,1,1},{1,0,1},{1,1,1}}

    setZeroes(arr)

    fmt.Println(arr)

}

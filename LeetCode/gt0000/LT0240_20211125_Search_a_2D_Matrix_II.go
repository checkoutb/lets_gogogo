// package sdq
package main

import (
    "fmt"
)


// int row = matrix.length-1;
// int col = 0;
// while (row >= 0 && col < matrix[0].length) {
//     if (matrix[row][col] > target) {
//         row--;
//     } else if (matrix[row][col] < target) {
//         col++;
//     } else { // found it
//         return true;
//     }
// }
// 这个是从 左下角开始的， 左下角作为 root，  把图旋转下，  让 左下角在 root 的位置 就可以 看到一颗 树


// 把 右上角作为root， 对于每个结点，col--就是 比它小的结点， row++ 就是比它大的结点。


// bool searchMatrix(int** A, int m, int n, int target) {
//     int x = ~target;
//     while (m && n && (x = A[0][n-1]) != target)
//         x < target ? A++, m-- : n--;
//     return x == target;
// }


// Runtime: 24 ms, faster than 78.74% of Go online submissions for Search a 2D Matrix II.
// Memory Usage: 6.8 MB, less than 66.14% of Go online submissions for Search a 2D Matrix II.
// 任意 矩形 中的 所有值 都是 >= 它的 左上角的值的。  好像没用。
// 记得 好像是 for for， 只不过 内层的for 不需要 从0开始。 大约就是 O(sz1+sz2) ... bu ,还是要回退的。。
func searchMatrix(matrix [][]int, target int) bool {
    sz1, sz2 := len(matrix), len(matrix[0])
    j := 0
    for i := 0; i < sz1; i++ {
        t2 := j
        for t2 >= 0 && matrix[i][t2] > target {
            t2--
        }
        if t2 > 0 {
            if matrix[i][t2 - 1] == target {
                return true
            }
        }
        for j < sz2 && matrix[i][j] < target {
            j++
        }
        if j < sz2 && matrix[i][j] == target {
            return true
        }
        j = t2
        if j < 0 {
            j = 0
        }
    }
    return false
}


func main_LT0240_20211125() {
// func main() {

    fmt.Println("ans:")


}

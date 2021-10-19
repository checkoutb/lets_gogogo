// package sdq
package main

import (
    "fmt"
)

// if (grid[i][j] == '.') {
//     continue;
// }
// int v = grid[i][j] - '0';
// int pos = 1 << v;
// if ((r[i] & pos) > 1) {
//     return false;
// }
// r[i] |= pos;
// if ((c[j] & pos) > 1) {
//     return false;
// }
// c[j] |= pos;
// int matrix = (i/3)*3 + j/3;
// if ((m[matrix] & pos) > 1) {
//     return false;
// }
// m[matrix] |= pos;
// int 是 32bit， 所以可以 1<<v.. 等于就是  每行一个 int， int的后9位 对应 board的9列。



// '4' in row 7 is encoded as "(4)7".
// '4' in column 7 is encoded as "7(4)".
// '4' in the top-right block is encoded as "0(4)2".
// 做一个 编码， 然后 set ，重复就是 false。







// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Sudoku.
// Memory Usage: 2.7 MB, less than 96.78% of Go online submissions for Valid Sudoku.
// 我在想，这个byte 是 5 还是 '5' 。。。好像都可以，那就是 '5' 了。。 ..   '.'
func isValidSudoku(board [][]byte) bool {
    // var t234 byte = '5'
    // var t435 byte = 5
    // fmt.Println(t234, ", ", t435)       // 53, 5

    ans := true
    arr := make([]int, 9)
    arr2 := make([]int, 9)
    arr3 := make([]int, 9)
    for i := 0; i < 9; i++ {
        arr[i], arr2[i], arr3[i] = -1, -1, -1
    }
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            a1,a2,a3,a4 := i / 3 * 3, i % 3 * 3, j / 3, j % 3
            t1 := board[a1 + a3][a2 + a4]
            if t1 != '.' {
                t1 = t1 - '1'
                if arr3[t1] == i {
                    return false
                }
                arr3[t1] = i
            }
            if board[i][j] != '.' {
                t2 := board[i][j] - '1'
                if arr[t2] == i {
                    return false
                }
                arr[t2] = i
            }
            if board[j][i] != '.' {
                t3 := board[j][i] - '1'
                if arr2[t3] == i {
                    return false
                }
                arr2[t3] = i
            }
        }
    }
    return ans
}



func main_LT0036_20211018() {
// func main() {

    fmt.Println("ans:")

    // arr2 := [][]byte{{'8','3','.','.','7','.','.','.','.'},          // , 必须在这行末尾。
    //                 {'6','.','.','1','9','5','.','.','.'},
    //                 {'.','9','8','.','.','.','.','6','.'},
    //                 {'8','.','.','.','6','.','.','.','3'},
    //                 {'4','.','.','8','.','3','.','.','1'},
    //                 {'7','.','.','.','2','.','.','.','6'},
    //                 {'.','6','.','.','.','.','2','8','.'},
    //                 {'.','.','.','4','1','9','.','.','5'},
    //                 {'.','.','.','.','8','.','.','7','9'}}

    arr2 := [][]byte{{'.','.','4','.','.','.','6','3','.'},
                    {'.','.','.','.','.','.','.','.','.'},
                    {'5','.','.','.','.','.','.','9','.'},
                    {'.','.','.','5','6','.','.','.','.'},
                    {'4','.','3','.','.','.','.','.','1'},
                    {'.','.','.','7','.','.','.','.','.'},
                    {'.','.','.','5','.','.','.','.','.'},
                    {'.','.','.','.','.','.','.','.','.'},
                    {'.','.','.','.','.','.','.','.','.'}}
    

    ans := isValidSudoku(arr2)

    fmt.Println(ans)

}

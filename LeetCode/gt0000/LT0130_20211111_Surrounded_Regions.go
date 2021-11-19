// package sdq
package main

import (
    "fmt"
)


// 一次描边 {边O -> Z}，    然后for-for 遍历一遍 { Z->O, O->X }


// Runtime: 28 ms, faster than 24.19% of Go online submissions for Surrounded Regions.
// Memory Usage: 16 MB, less than 5.28% of Go online submissions for Surrounded Regions.
func solve(board [][]byte)  {
    sz1, sz2 := len(board), len(board[0])
    for i := 0; i < sz2; i++ {
        if board[0][i] == 'O' {
            filla130(board, 0, i, 'O', 'Z')
        }
        if board[sz1 - 1][i] == 'O' {
            filla130(board, sz1 - 1, i, 'O', 'Z')
        }
    }
    for i := 1; i < sz1; i++ {
        filla130(board, i, 0, 'O', 'Z')
        filla130(board, i, sz2 - 1, 'O', 'Z')
    }
    // fmt.Println(board)
    for i := 1; i < sz1 - 1; i++ {
        for j := 1; j < sz2 - 1; j++ {
            filla130(board, i, j, 'O', 'X')
        }
    }

    for i := 0; i < sz2; i++ {
        filla130(board, 0, i, 'Z', 'O')
        filla130(board, sz1 - 1, i, 'Z', 'O')
    }
    for i := 1; i < sz1 - 1; i++ {
        filla130(board, i, 0, 'Z', 'O')
        filla130(board, i, sz2 - 1, 'Z', 'O')
    }
}

func filla130(board [][]byte, i, j int, ori, new byte) {
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != ori {
        return
    }
    board[i][j] = new
    // -1 0   0 -1   1 0     0 1
    arr := []int{-1,0,1,0,-1}
    for k := 0; k < 4; k++ {
        filla130(board, i + arr[k], j + arr[k + 1], ori, new)
    }
}


func main_LT0130_20211111() {
// func main() {

    fmt.Println("ans:")

    arr := [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}

    solve(arr)

    fmt.Println(arr)

}

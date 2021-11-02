// package sdq
package main

import (
    "fmt"
)



// board[x][y] = 0;


// Runtime: 219 ms, faster than 24.15% of Go online submissions for Word Search.
// Memory Usage: 2.3 MB, less than 22.85% of Go online submissions for Word Search.
// 要快的话,是不是可以 一个3维数组, [i][j][idx] 表明 board[i][j] 是 word[idx] 时, 是否可行.  只用设置一次, 成功就直接返回了.
func exist(board [][]byte, word string) bool {
    sz1, sz2 := len(board), len(board[0])
    vst := make([][]bool, sz1)
    for i, _ := range vst {
        vst[i] = make([]bool, sz2)
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if dfsa79(board, word, 0, i, j, vst) {
                return true
            }
        }
    }
    return false
}

func dfsa79(board [][]byte, word string, idx int, i, j int, vst [][]bool) (ans bool) {
    if idx == len(word) {
        return true
    }
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || vst[i][j] {
        return false
    }
    if board[i][j] != word[idx] {
        return false
    }
    vst[i][j] = true

    dir := [][]int{{1,0},{0,1},{-1,0},{0,-1}}
    for _, ar := range dir {
        if dfsa79(board, word, idx + 1, i + ar[0], j + ar[1], vst) {
            return true
        }
    }
    vst[i][j] = false

    // ans := false         // no new variables on left side of :=
    ans = false     // 返回的形参 已经声明了.
    return
}

func main_LT0079_20211025() {
// func main() {

    fmt.Println("ans:")

    arr := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
    // w := "ABCCED"
    // w := "SEE"
    w := "ABCB"

    ans := exist(arr, w)

    fmt.Println(ans)

}

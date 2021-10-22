// package sdq
package main

import (
    "fmt"
)





// if(board[i][j] == 'Q' && (x + j == y + i || x + y == i + j || x == i))
// xy 是校验的点

// 只需要校验边， 就是把 已有的q 映射到边上， 后续的 也把自己映射到边上，看是否重叠。
// 边的话，应该是 垂直向下， 垂直向下的 左右45度。 所以需要 3个 一维bool数组。
// 45度，数组长度是 2n
// if (flag_col[col] && flag_45[row + col] && flag_135[n - 1 + col - row]) {


// 不需要check 那么多， 只需要 往下 attack 就可以了。

// Runtime: 7 ms, faster than 47.92% of Go online submissions for N-Queens.
// Memory Usage: 3.7 MB, less than 34.72% of Go online submissions for N-Queens.
// 好像只需要算一半， 棋盘上下翻转一下 就是 新的 ans 了。。
// 负数代表q， 0代表空， 正数代表会被攻击。
// slice 默认指针的， 为什么 一直都是 *[][]string ?
func solveNQueens(n int) [][]string {
    ans := [][]string{}
    board := make([][]int, n)
    for i, _ := range board {
        board[i] = make([]int, n)
    }
    dfsa51(n, &ans, board, 0)
    return ans
}

func dfsa51(n int, ans *[][]string, board [][]int, idx int) {
    if idx == n {
        arr := []string{}
        for _, ar := range board {
            str := ""
            for _, e := range ar {
                if e < 0 {
                    str += "Q"
                } else {
                    str += "."
                }
            }
            arr = append(arr, str)
        }
        *ans = append(*ans, arr)
        return
    }
    for i := 0; i < n; i++ {
        if board[idx][i] == 0 {
            board[idx][i] = -1
            attack(board, idx, i, idx, true, n)
            dfsa51(n, ans, board, idx + 1)
            attack(board, idx, i, idx, false, n)
            board[idx][i] = 0
        }
    }
}

// b2 true |  false ^
func attack(board [][]int, i int, j int, sft int, b2 bool, n int) {
    st := 1 << sft
    for b := 0; b < n; b++ {
        if b == j {
            continue
        }
        if b2 {
            board[i][b] |= st
        } else {
            board[i][b] ^= st
        }
    }
    for a := 0; a < n; a++ {
        if a == i {
            continue
        }
        if b2 {
            board[a][j] |= st
        } else {
            board[a][j] ^= st
        }
    }
    for k := 1; k < n; k++ {
        // a1,b1,a2,b2,a3,b3,a4,b4 := 
        arr := []int{
            i + k, j + k,
            i - k, j - k,
            i + k, j - k,
            i - k, j + k}
        for u := 0; u < 8; u += 2 {
            a, b := arr[u], arr[u + 1]
            if a >= 0 && a < n && b >= 0 && b < n {
                if b2 {
                    board[a][b] |= st
                } else {
                    board[a][b] ^= st
                }
            }
        }
    }
}


func main_LT0051_20211020() {
// func main() {

    fmt.Println("ans:")

    // n := 4
    n := 1

    ans := solveNQueens(n)

    for _, s := range ans {
        fmt.Println(s)
    }

}

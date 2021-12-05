// package sdq
package main

import (
    "fmt"
)


// for _,newPos := range [8][2]int{ {r-1,c-1},{r-1,c},{r-1,c+1},{r+1,c-1},{r+1,c},{r+1,c+1},{r,c-1},{r,c+1}} {


// for (int x = Math.max(i - 1, 0); x <= Math.min(i + 1, m - 1); x++) {
//     for (int y = Math.max(j - 1, 0); y <= Math.min(j + 1, n - 1); y++) {
//         lives += board[x][y] & 1;
//     }
// }

// [2nd bit, 1st bit] = [next state, current state]
// board[i][j] & 1
// board[i][j] >> 1




// Runtime: 0 ms, faster than 100.00% of Go online submissions for Game of Life.
// Memory Usage: 2 MB, less than 49.41% of Go online submissions for Game of Life.
// 活的， <2 die, 2,3 live, >3 die
// die, 3 活
func gameOfLife(board [][]int)  {
    sz1, sz2 := len(board), len(board[0])
    for i := 0; i < sz1; i++ {
        for j := 0; j < sz2; j++ {
            cnt := howManyNeighbours(board, i, j)
            if board[i][j] == 0 {
                if cnt == 3 {
                    board[i][j] += 2
                }
            } else {
                if cnt != 2 && cnt != 3 { board[i][j] += 2 }
                // switch cnt {
                // case 2:
                //     fallthrough
                // case 3:
                //     ;
                // default:
                //     board[i][j] += 2
                // }
            }
        }
    }
    for i := 0; i < sz1; i++ {
        for j := 0; j < sz2; j++ {
            if board[i][j] > 1 { board[i][j] = (board[i][j] + 1) % 2 }
        }
    }
}

func howManyNeighbours(arr [][]int, i, j int) int {
    cnt := 0
    for a := i - 1; a <= i + 1; a++ {
        for b := j - 1; b <= j + 1; b++ {
            cnt += getVala289(arr, a, b)
        }
    }
    return cnt - arr[i][j]
}

// invalid.val == 0
// % 2    2:die->live   3:live->die
func getVala289(arr [][]int, i, j int) int {
    if i < 0 || j < 0 || i >= len(arr) || j >= len(arr[0]) { return 0 }
    return arr[i][j] % 2
}


func main_LT0289_20211203() {
// func main() {

    fmt.Println("ans:")


}

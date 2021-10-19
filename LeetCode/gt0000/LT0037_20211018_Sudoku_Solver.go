// package sdq
package main

import (
    "fmt"
)


// 硬算 简单而快乐。。。  不要任何辅助数组，不要任何记录，就硬算。每次都硬算。。。 确实，刷新状态 太恶心了。。。







// tmd, tu le.....

// Runtime: 7 ms, faster than 41.13% of Go online submissions for Sudoku Solver.
// Memory Usage: 2.5 MB, less than 18.15% of Go online submissions for Sudoku Solver.
// 把每行，每列，每块， 做一个id。 priority_queue, <还有多少空格 , id> 。。。 no
// 每个格子 一个 id， pri_que <还有多少种可能， id>   bitmap,  order by bitmap.1的个数
func solveSudoku(board [][]byte)  {
    arr := [9][9]Info{}
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            arr[i][j] = Info{i,j,9,0}
        }
    }
    map2 := map[int][]*Info{}     // row           array...
    map3 := map[int][]*Info{}     // col
    map5 := map[int][]*Info{}     // block
    emp := 0
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' {
                arr[i][j].row = i
                arr[i][j].col = j
                map2[i] = append(map2[i], &arr[i][j])
                map3[j] = append(map3[j], &arr[i][j])
                t2 := i / 3 * 3 + j / 3
                map5[t2] = append(map5[t2], &arr[i][j])
                emp++
            }
        }
    }

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] != '.' {
                t3 := board[i][j] - '0'
                t1 := 1 << t3
                t2 := i / 3 * 3 + j / 3
                for _, v := range map2[i] {
                    if ((*v).bit & t1) == 0 {
                        (*v).cnt = v.cnt - 1
                        (*v).bit = v.bit | t1
                        // fmt.Println("543 .", v.bit, ", ", v.cnt, ", ", v)
                    }
                }
                for _, v := range map3[j] {
                    if (v.bit & t1) == 0 {
                        v.cnt = v.cnt - 1
                        v.bit = v.bit | t1
                    }
                }
                for _, v := range map5[t2] {
                    if (v.bit & t1) == 0 {
                        v.cnt = v.cnt - 1
                        v.bit = v.bit | t1
                    }
                }
            }
        }
    }

    // for i := 0; i < 9; i++ {
    //     for j := 0; j < 9; j++ {
    //         if (board[i][j] == '.') {
    //             fmt.Println(arr[i][j].row, ", ", arr[i][j].col, ", ", arr[i][j].cnt, ", ", arr[i][j].bit, ", ", i, ", ", j)
    //         }
    //     }
    // }

    b2 := dfsa37(map2, map3, map5, board, emp)

    fmt.Println(b2)

}

type Info struct {
    row int
    col int
    cnt int             // 还有多少种可能
    bit int         // 1<<x | ....
}

func dfsa37(map2 map[int][]*Info, map3 map[int][]*Info, map5 map[int][]*Info, board [][]byte, emp int) bool {
    if (emp == 0) {
        return true
    }
    var info *Info = nil
    mncnt := 10
    // any := false
    lt0 := false
// fmt.Println("...")
    for _, v := range map2 {
        for _, e := range v {
            // any = true
            if e.cnt < 0 {
                lt0 = true
            }
            if e.cnt < mncnt && e.cnt > 0 && board[e.row][e.col] == '.' {       // .................
                mncnt = e.cnt
                info = e
            }
        }
    }
    if lt0 {
// fmt.Println("lt0")
        return false
    }
    if info == nil {
        return false
        // return !any
    }
//     if !any {
// // fmt.Println("TTT")
//         return true
//     }
// fmt.Println(info.row, ", ", info.col, ", ", info.bit, ", ", info.cnt)
    // if mncnt == 10 {
    //     return true
    // }
    for i := 1; i <= 9; i++ {
        if (info.bit & (1 << i)) == 0 {
            t2 := 1 << i
            r := info.row
            c := info.col
            t3 := r / 3 * 3 + c / 3
            lt02 := false
            // evict := []int{}
            arrchg := []*Info{}                     // ......................................
            for _, e12 := range map2[r] {
                if (e12.bit & t2) == 0 {
                    e12.bit = e12.bit | t2
                    e12.cnt = e12.cnt - 1
                    if e12.cnt < 0 {
                        lt02 = true
                    }
                    arrchg = append(arrchg, e12)
                }
                // if e12.cnt == 0 {
                //     evict = append(evict, idx)
                // }
            }
//             lst := len(map2[r]) - 1
// // fmt.Println(lst, ".....lst ")
//             for i := len(evict) - 1; i >= 0; i-- {
//                 val := evict[i]
//                 map2[r][val], map2[r][lst] = map2[r][lst], map2[r][val]
//                 lst--
//             }
            // for _, val := range evict {
            //     if val < lst {
            //         map2[r][val], map2[r][lst] = map2[r][lst], map2[r][val]
            //     }
            //     lst--
            // }
            // tarr2 := map2[r][lst + 1 : ]
            // map2[r] = map2[r][0 : lst + 1]
            
            
            for _, e12 := range map3[c] {
                if (e12.bit & t2) == 0 {
                    e12.bit = e12.bit | t2
                    e12.cnt = e12.cnt - 1
                    arrchg = append(arrchg, e12)
                }
            }
            for _, e12 := range map5[t3] {
                if (e12.bit & t2) == 0 {
                    e12.bit = e12.bit | t2
                    e12.cnt = e12.cnt - 1
                    arrchg = append(arrchg, e12)
                }
            }
            // fmt.Println(r, ", ", c, ", ", board[r][c] - '0')
            if !lt02 {
                board[r][c] = byte(i + '0')
                b2 := dfsa37(map2, map3, map5, board, emp - 1);
                if b2 {
                    return true
                }
                board[r][c] = '.'                           // ...........................
            }
            for _, e12 := range arrchg {
                e12.bit = e12.bit ^ t2
                e12.cnt = e12.cnt + 1
            }
            // map2[r] = append(map2[r], tarr2...)
            // for _, e12 := range map2[r] {
            //     e12.bit = e12.bit ^ t2
            //     e12.cnt = e12.cnt + 1
            // }
            // for _, e12 := range map3[c] {
            //     e12.bit = e12.bit ^ t2
            //     e12.cnt = e12.cnt + 1
            // }
            // for _, e12 := range map5[t3] {
            //     e12.bit = e12.bit ^ t2
            //     e12.cnt = e12.cnt + 1
            // }
        }
    }
// fmt.Println("FFFF")
    return false
}

func main_LT0037_20211018() {
// func main() {

    fmt.Println("ans:")

    // arr := [][]byte{{'5','3','.','.','7','.','.','.','.'},
    //                 {'6','.','.','1','9','5','.','.','.'},
    //                 {'.','9','8','.','.','.','.','6','.'},
    //                 {'8','.','.','.','6','.','.','.','3'},
    //                 {'4','.','.','8','.','3','.','.','1'},
    //                 {'7','.','.','.','2','.','.','.','6'},
    //                 {'.','6','.','.','.','.','2','8','.'},
    //                 {'.','.','.','4','1','9','.','.','5'},
    //                 {'.','.','.','.','8','.','.','7','9'}}


// [["5","1","9","7","4","8","6","3","2"],
// ["7","8","3","6","5","2","4","1","9"],
// ["4","2","6","1","3","9","8","7","5"],
// ["3","5","7","9","8","6","2","4","1"],
// ["2","6","4","3","1","7","5","9","8"],
// ["1","9","8","5","2","4","3","6","7"],
// ["9","7","5","8","6","3","1","2","4"],
// ["8","3","2","4","9","1","7","5","6"],
// ["6","4","1","2","7","5","9","8","3"]]

    arr := [][]byte{{'.','.','9','7','4','8','.','.','.'},
                    {'7','.','.','.','.','.','.','.','.'},
                    {'.','2','.','1','.','9','.','.','.'},
                    {'.','.','7','.','.','.','2','4','.'},
                    {'.','6','4','.','1','.','5','9','.'},
                    {'.','9','8','.','.','.','3','.','.'},
                    {'.','.','.','8','.','3','.','2','.'},
                    {'.','.','.','.','.','.','.','.','6'},
                    {'.','.','.','2','7','5','9','.','.'}}
    

    for _, ar := range arr {
        for _, e := range ar {
            if e == '.' {
                fmt.Print(".", "  ")
                continue
            }
            fmt.Print(e - '0', "  ")
        }
        fmt.Println()
    }

    solveSudoku(arr)

    for _, ar := range arr {
        for _, e := range ar {
            if e == '.' {
                fmt.Print(".", "  ")
                continue
            }
            fmt.Print(e - '0', "  ")
        }
        fmt.Println()
    }


    a := 8
    a = a | 2
    fmt.Println(a, ", ")
    a = a ^ 2
    fmt.Println(a, ". ")

}

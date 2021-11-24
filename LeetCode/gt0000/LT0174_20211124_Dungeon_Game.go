// package sdq
package main

import (
    "fmt"
)








// Runtime: 4 ms, faster than 92.86% of Go online submissions for Dungeon Game.
// Memory Usage: 3.7 MB, less than 47.14% of Go online submissions for Dungeon Game.
// 是不是二分求血量？  还是反推？
func calculateMinimumHP(dungeon [][]int) int {
    sz1, sz2 := len(dungeon), len(dungeon[0])
    arr := make([][]int, sz1)
    for i, _ := range arr {
        arr[i] = make([]int, sz2)
    }
    arr[sz1 - 1][sz2 - 1] = 1 - dungeon[sz1 - 1][sz2 - 1]            // 进入这个格子，至少需要 多少血，才能不死。
    if arr[sz1 - 1][sz2 - 1] <= 0 {
        arr[sz1 - 1][sz2 - 1] = 1
    }
    for i := sz1 - 2; i >= 0; i-- {
        arr[i][sz2 - 1] = arr[i + 1][sz2 - 1] - dungeon[i][sz2 - 1]
        if arr[i][sz2 - 1] <= 0 {
            arr[i][sz2 - 1] = 1
        }
    }
    for i := sz2 - 2; i >= 0; i-- {
        arr[sz1 - 1][i] = arr[sz1 - 1][i + 1] - dungeon[sz1 - 1][i]
        if arr[sz1 - 1][i] <= 0 {
            arr[sz1 - 1][i] = 1
        }
    }
    for i := sz1 - 2; i >= 0; i-- {
        for j := sz2 - 2; j >= 0; j-- {
            t2 := arr[i + 1][j] - dungeon[i][j]
            t3 := arr[i][j + 1] - dungeon[i][j]
            if t2 <= 0 {
                t2 = 1
            }
            if t3 <= 0 {
                t3 = 1
            }
            if t2 < t3 {
                arr[i][j] = t2
            } else {
                arr[i][j] = t3
            }
        }
    }
    return arr[0][0]
}

func main_LT0174_20211124() {
// func main() {

    fmt.Println("ans:")

    arr := [][]int{{-2,-3,3},{-5,-10,1},{10,30,-5}}

    ans := calculateMinimumHP(arr)

    fmt.Println(ans)
}

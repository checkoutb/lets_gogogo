// package sdq
package main

import (
    "fmt"
)








// Runtime: 9 ms, faster than 24.18% of Go online submissions for Number of Islands.
// Memory Usage: 4 MB, less than 59.74% of Go online submissions for Number of Islands.
// flood-fill
func numIslands(grid [][]byte) int {
    sz1, sz2 := len(grid), len(grid[0])
    ans := 0
    for i := 0; i < sz1; i++ {
        for j := 0; j < sz2; j++ {
            if grid[i][j] == '1' {
                ans++
                floodfilla200(grid, i, j)
            }
        }
    }
    return ans
}
var dirs []int = []int{1,0,-1,0,1}
func floodfilla200(grid [][]byte, i, j int) {
    if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == '0' { return }
    grid[i][j] = '0'
    for a := 0; a < 4; a++ {
        floodfilla200(grid, i + dirs[a], j + dirs[a + 1])
    }
}

func main_LT0200_20211125() {
// func main() {

    fmt.Println("ans:")


}

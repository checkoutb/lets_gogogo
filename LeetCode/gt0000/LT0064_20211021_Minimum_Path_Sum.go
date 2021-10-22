// package sdq
package main

import (
    "fmt"
)






// Runtime: 9 ms, faster than 52.96% of Go online submissions for Minimum Path Sum.
// Memory Usage: 4.1 MB, less than 76.32% of Go online submissions for Minimum Path Sum.
func minPathSum(grid [][]int) int {
    sz1, sz2 := len(grid), len(grid[0])
    
    for i := 1; i < sz1; i++ {
        grid[i][0] += grid[i - 1][0]
    }
    for i := 1; i < sz2; i++ {
        grid[0][i] += grid[0][i - 1]
    }

    for i := 1; i < sz1; i++ {
        for j := 1; j < sz2; j++ {
            // if i == 1 && j == 2 {
            //     fmt.Println(grid[i - 1][j], ", ", grid[i][j - 1], ", ", grid[i - 1][j] > grid[i][j - 1])
            // }
            if grid[i - 1][j] > grid[i][j - 1] {
                // fmt.Println("111")
                grid[i][j] += grid[i][j - 1]
            } else {
                grid[i][j] += grid[i - 1][j]                // [j-1][i]。。。 .......
                // fmt.Println("222, ", grid[i][j])
            }
        }
    }
    // for _, ar := range grid {
    //     fmt.Println(ar)
    // }
    return grid[sz1 - 1][sz2 - 1]
}


func main_LT0064_20211021() {
// func main() {

    fmt.Println("ans:")

    arr := [][]int{{1,3,1},{1,5,1},{4,2,1}}

    ans := minPathSum(arr)

    fmt.Println(ans)

}

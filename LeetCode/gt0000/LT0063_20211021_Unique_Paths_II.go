// package sdq
package main

import (
    "fmt"
)




// if obstacleGrid[i][0] == 1 {
//     dp[i][0] = 0
// } else {
//     dp[i][0] = dp[i-1][0]
// }


// vector<vector<int> > dp(m + 1, vector<int> (n + 1, 0));
// dp[0][1] = 1;
// for (int i = 1; i <= m; i++)
//     for (int j = 1; j <= n; j++)
//         if (!obstacleGrid[i - 1][j - 1])
//             dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
// return dp[m][n];
// 当前的下标 左上不是石头， 说明 当前下标的 左侧 上方 是 路径。 。不， 这里跳过了一些东西， 就是 石头 右/下侧 的空白 也是有 路径的，但是这里就放弃了。 
// 石头 右侧的 会被 石头的 右 右 测 继承。  而不会被 石头的 右下侧集成。




// 0 空白， 1石头  .   右下角 会不会是 石头？ 左上角呢。。    都会。。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    sz1, sz2 := len(obstacleGrid), len(obstacleGrid[0])
    if obstacleGrid[sz1 - 1][sz2 - 1] == 1 || obstacleGrid[0][0] == 1 {
        return 0
    }
    // for _, ar := range obstacleGrid {
    //     fmt.Println(ar)
    // }
    // fmt.Println("==============")
    obstacleGrid[0][0] = 1
    b2 := true
    for i := 1; i < sz1; i++ {
        if obstacleGrid[i][0] == 1 {
            obstacleGrid[i][0] = 0
            b2 = false
            // break                // 导致 后面的 石头 没有被清空。。.后面是 不遍历 第一列 第一行的， 所以 第一列 第一行的 石头被当作了 可用路径
        } else {
            if b2 {
                obstacleGrid[i][0] = 1
            }
        }
    }

    b2 = true
    for i := 1; i < sz2; i++ {
        if obstacleGrid[0][i] == 1 {
            obstacleGrid[0][i] = 0
            b2 = false
            // break
        } else {
            if b2 {
                obstacleGrid[0][i] = 1
            }
        }
    }

    for i := 1; i < sz1; i++ {
        for j := 1; j < sz2; j++ {
            if obstacleGrid[i][j] == 1 {
                obstacleGrid[i][j] = 0
            } else {
                obstacleGrid[i][j] = obstacleGrid[i - 1][j] + obstacleGrid[i][j - 1]
            }
        }
    }
    
    // for _, ar := range obstacleGrid {
    //     fmt.Println(ar)
    // }

    return obstacleGrid[sz1 - 1][sz2 - 1]
}


func main_LT0063_20211021() {
// func main() {

    fmt.Println("ans:")

    // arr := [][]int{{0,0,0},{0,1,0},{0,0,0}}
    // arr := [][]int{{0,1},{0,0}}
    arr := [][]int{{0,0},{0,0},{0,0},{0,0},{0,0},{0,0},{1,0},{0,0},{0,0},
    {0,0},{0,0},{0,0},{1,0},{0,0},{0,0},{0,0},{0,0},{0,1},
    {0,0},{0,0},{1,0},{0,0},{0,0},{0,1},{0,0},{0,0},{0,0},{0,0},
    {0,0},{0,0},{0,0},{0,1},{0,0},{0,0},{0,0},{0,0},{1,0},{0,0},
    {0,0},{0,0},{0,0}}

    // arr := [][]int{{0,0},{1,0},{0,0},{0,1},{0,0},{0,0}}


    ans := uniquePathsWithObstacles(arr)

    fmt.Println(ans)

}

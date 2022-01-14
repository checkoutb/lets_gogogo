// package sdq
package main

import (
    "fmt"
)




// memo := make([][]int, row)
// for i:=0;i<row;i++ {
//     memo[i] = make([]int, col)
// }
// maxFound := 1 
//
// var dp func(i, j int) int
// dp = func(i, j int) int {
//     max := 1
//     if memo[i][j] > 0 {
//         return memo[i][j]
//     }
//    
//     if i+1<row && matrix[i][j] < matrix[i+1][j] {
//         count := 1 + dp(i+1, j)
//         if count > max {
//             max = count
//         }
//     }




// Runtime: 40 ms, faster than 63.70% of Go online submissions for Longest Increasing Path in a Matrix.
// Memory Usage: 7.6 MB, less than 49.63% of Go online submissions for Longest Increasing Path in a Matrix.
// 200*200 如果每个点都计算 感觉可能tle。 不 应该不会。 并不是重新计算。
func longestIncreasingPath(matrix [][]int) int {
    sz1, sz2 := len(matrix), len(matrix[0])
    ans := 0
    arr := make([][]int, sz1)
    for i, _ := range arr {
        arr[i] = make([]int, sz2)
    }
    for i := 0; i < sz1; i++ {
        for j := 0; j < sz2; j++ {
            t2 := dfsa329a(matrix, i, j, -1, &arr)
            if t2 > ans {
                ans = t2
            }
        }
    }
    return ans
}

// find bigger, return len(path)
// <= prev 跳过了 来源。  是不是 只需要计算一次？  应该是的。
func dfsa329a(matrix [][]int, i, j int, prev int, arr *[][]int) int {
    if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) || matrix[i][j] <= prev {
        return 0
    }
    if (*arr)[i][j] != 0 {
        return (*arr)[i][j]
    }
    a1,a2,a3,a4 :=  dfsa329a(matrix, i + 1, j, matrix[i][j], arr),
                    dfsa329a(matrix, i - 1, j, matrix[i][j], arr),
                    dfsa329a(matrix, i, j + 1, matrix[i][j], arr),
                    dfsa329a(matrix, i, j - 1, matrix[i][j], arr)

    (*arr)[i][j] = mx329(a1, mx329(a2, mx329(a3, a4))) + 1
    return (*arr)[i][j]
}

func mx329(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main_LT0329_20220113() {
// func main() {

    fmt.Println("ans:")

    // arr := [][]int{{9,9,4},{6,6,8},{2,1,1}}
    arr := [][]int{{1}}

    ans := longestIncreasingPath(arr)

    fmt.Println(ans)

}

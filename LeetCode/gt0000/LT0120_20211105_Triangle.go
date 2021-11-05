// package sdq
package main

import (
    "fmt"
)


// for i := len(triangle) - 2; i >= 0; i-- {
//     for j := 0; j < len(triangle[i]); j++ {
//         triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
//     }
// }


// vector<int> minPathSum = triangle[n-1];
// for(int i = n-2; i >=0; i--)
//     for(int j = 0; j <= i; j++)
//     {
//         minPathSum[j] = triangle[i][j] + min(minPathSum[j], minPathSum[j+1]);
//     }
// return minPathSum[0];


// Runtime: 4 ms, faster than 92.81% of Go online submissions for Triangle.
// Memory Usage: 3.3 MB, less than 96.23% of Go online submissions for Triangle.
// ？ dp。。 标准dp。。
func minimumTotal(triangle [][]int) int {
    sz1 := len(triangle)
    for i := 1; i < sz1; i++ {
        sz2 := len(triangle[i])
        triangle[i][0], triangle[i][sz2 - 1] = triangle[i][0] + triangle[i - 1][0], triangle[i][sz2 - 1] + triangle[i - 1][sz2 - 2]
        for j := 1; j < len(triangle[i]) - 1; j++ {
            if triangle[i - 1][j - 1] > triangle[i - 1][j] {
                triangle[i][j] += triangle[i - 1][j]
            } else {
                triangle[i][j] += triangle[i - 1][j - 1]
            }
        }
    }
    ans := 200 * 10001
    for _, v := range triangle[sz1 - 1] {
        if v < ans {
            ans = v
        }
    }
    return ans
}

func main_LT0120_20211105() {
// func main() {

    fmt.Println("ans:")


}

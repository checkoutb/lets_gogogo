// package sdq
package main

import (
    "fmt"
)


// vector<int> vi(n, 0);
// vi[0] = 1;
// while (m-- > 0)
// {
//     for (int i = 1; i < n; ++i)
//     {
//         vi[i] = vi[i - 1] + vi[i];
//     }
// }
// return vi[n - 1];


// (m+n)! / (m! * n!)



// Runtime: 4 ms, faster than 9.09% of Go online submissions for Unique Paths.
// Memory Usage: 2.4 MB, less than 17.25% of Go online submissions for Unique Paths.
func uniquePaths(m int, n int) int {
    arr := make([][]int, m)
    for i := 0; i < m; i++ {
        arr[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        arr[i][0] = 1
    }
    for i := 0; i < n; i++ {
        arr[0][i] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            arr[i][j] = arr[i - 1][j] + arr[i][j - 1]
        }
    }
    fmt.Println(arr)
    return arr[m - 1][n - 1]
}


func main_LT0062_20211021() {
// func main() {

    fmt.Println("ans:")

    // m, n := 3, 7
    // m, n := 3, 2
    m, n := 3, 3

    ans := uniquePaths(m, n)

    fmt.Println(ans)

}

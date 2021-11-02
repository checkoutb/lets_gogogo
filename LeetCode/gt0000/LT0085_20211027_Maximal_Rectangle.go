// package sdq
package main

import (
    "fmt"
)


// left(i,j) = max(left(i-1,j), cur_left), cur_left can be determined from the current row
// right(i,j) = min(right(i-1,j), cur_right), cur_right can be determined from the current row
// height(i,j) = height(i-1,j) + 1, if matrix[i][j]=='1';
// height(i,j) = 0, if matrix[i][j]=='0'
// buzhidaoshuosha, geiningpigechaba



// Runtime: 4 ms, faster than 78.21% of Go online submissions for Maximal Rectangle.
// Memory Usage: 4.1 MB, less than 71.79% of Go online submissions for Maximal Rectangle.
// ....... byte == char...
func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
    sz1, sz2 := len(matrix), len(matrix[0])
    ans := 0
    arr := make([][]int, sz1)
    for i, _ := range arr {
        arr[i] = make([]int, sz2)
    }
    for i := sz1 - 1; i >= 0; i-- {
        for j := sz2 - 1; j >= 0; j-- {
            if matrix[i][j] != '0' {
                if j == sz2 - 1 {
                    arr[i][j] = 1
                } else {
                    arr[i][j] = arr[i][j + 1] + 1
                }
            }
        }
    }
    for i := 0; i < sz1; i++ {
        for j := 0; j < sz2; j++ {
            if arr[i][j] != 0 {
                t2 := int(arr[i][j])
                cnt := 1
                if t2 * cnt > ans {
                    ans = t2 * cnt
                }
                for k := i + 1; k < sz1 && arr[k][j] != 0; k++ {
                    cnt++
                    if int(arr[k][j]) < t2 {
                        t2 = int(arr[k][j])
                    }
                    if (t2 * cnt > ans) {
                        ans = t2 * cnt
                    }
                }
            }
        }
    }
    return ans
}


func main_LT0085_20211027() {
// func main() {

    fmt.Println("ans:")

    // arr := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
    arr := [][]byte{{'1','0','1'}}

    ans := maximalRectangle(arr)

    fmt.Println(ans)
}

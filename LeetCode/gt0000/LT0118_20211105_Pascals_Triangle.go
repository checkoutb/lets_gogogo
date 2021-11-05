// package sdq
package main

import (
    "fmt"
)


// res = [[1]]
// for i in range(1, numRows):
//     res += [map(lambda x, y: x+y, res[-1] + [0], [0] + res[-1])]
// return res[:numRows]




// Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
// Memory Usage: 2.1 MB, less than 99.81% of Go online submissions for Pascal's Triangle.
func generate(numRows int) [][]int {
    ans := [][]int{{1}}
    for i := 1; i < numRows; i++ {
        arr := make([]int, i + 1)
        arr[0], arr[i] = 1, 1
        for j := 1; j < i; j++ {
            arr[j] = ans[i - 1][j] + ans[i - 1][j - 1]
        }
        ans = append(ans, arr)
    }
    // for numRows > 1 {
    //     numRows--
    //     sz1 := len(ans)
    //     // arr := []int{ans[sz1 - 1][0]}            // length = 1,2,3,4,5,6,7...
    //     arr := []int{1}

    //     arr = append(arr, 1)
    // }
    return ans
}


func main_LT0118_20211105() {
// func main() {

    fmt.Println("ans:")


}

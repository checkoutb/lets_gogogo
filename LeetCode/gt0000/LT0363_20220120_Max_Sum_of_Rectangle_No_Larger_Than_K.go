// package sdq
package main

import (
    "fmt"
)



//gg

// 。。。
func maxSumSubmatrix(matrix [][]int, k int) int {
    sz1, sz2 := len(matrix), len(matrix[0])
    ans := 0
    larr := make([][]int, sz1)
    for i, _ := range larr {
        larr[i] = make([]int, sz2)
    }
    for j := sz2 - 2; j >= 0; j-- {
        larr[sz1 - 1][j] = larr[sz1 - 1][j + 1] + matrix[sz1 - 1][j]
    }
    sfx := 0
    for i := sz1 - 2; i >= 0; i-- {
        sfx = 0
        for j := sz2 - 1; j >= 0; j-- {
            sfx += matrix[i][j]
            larr[i][j] = larr[i + 1][j] + sfx
        }
    }

    fmt.Println(larr)




    return ans
}

func main_LT0363_20220120() {
// func main() {

    fmt.Println("ans:")


}

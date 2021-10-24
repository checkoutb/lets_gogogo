// package sdq
package main

import (
    "fmt"
)



// 。。。。。。


// error
// 只有长度 > 的时候，才会考虑删除。 而不是跳过？
// 求 公共子subseq    len(wrod1) - len(subseq) == ans 。。 不， Example2 不是这种。
// 还是 subseq， 然后 再处理。
// 感觉可以直接dp。
func minDistance(word1 string, word2 string) int {
    sz1, sz2 := len(word1), len(word2)
    arr := make([][]int, sz1)
    for i, _ := range arr {
        arr[i] = make([]int, sz2)
    }
    for i := 0; i < sz1; i++ {
        for j := 0; j < sz2; j++ {
            if word1[i] == word2[j] {
                if i > 0 && j > 0 {
                    arr[i][j] = arr[i - 1][j - 1]
                } else {
                    if i == 0 {
                        arr[i][j] = j           // j + 1 - 1
                    } else if j == 0 {
                        arr[i][j] = arr[i - 1][j]
                    }
                }
            } else {
                if i > 0 && j > 0 {
                    arr[i][j] = arr[i - 1][j - 1] + 1
                } else {
                    if i == 0 {
                        arr[i][j] = j + 1
                    } else {
                        // arr[i][j] = i - j + 1
                        arr[i][j] = arr[i - 1][j] + 1
                    }
                }
            }
        }
    }
    fmt.Println(arr)
    return arr[sz1 - 1][sz2 - 1]
}


func main_LT0072_20211023() {
// func main() {

    fmt.Println("ans:")

    s1, s2 := "horse", "ros"

    ans := minDistance(s1, s2)

    fmt.Println(ans)
}

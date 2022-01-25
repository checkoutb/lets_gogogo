// package sdq
package main

import (
    "fmt"
    "strings"
)


// 保存 从头到自己的 路径的长度。  len()是深度。



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Absolute File Path.
// Memory Usage: 2.2 MB, less than 41.18% of Go online submissions for Longest Absolute File Path.
// \t
func lengthLongestPath(input string) int {
    stk := []Node388{}
    arr := strings.Split(input, "\n")
    ans := 0
    for i := 0; i < len(arr); i++ {
        sz1 := 0
        isF := false
        depth := 0
        for j := 0; j < len(arr[i]); j++ {
            if arr[i][j] == '\t' {
                sz1 = 1
                depth++
            } else {
                sz1++
                if arr[i][j] == '.' {
                    isF = true
                }
            }
        }
        for len(stk) > 0 && stk[len(stk) - 1].dep >= depth {
            stk = stk[0 : len(stk) - 1]
        }
        if isF {
            for _, v := range stk {
                sz1 += v.sz1
            }
            if sz1 > ans {
                ans = sz1
            }
        } else {
            stk = append(stk, Node388{ depth, sz1 })
        }
    }
    return ans
}

type Node388 struct {
    dep int
    sz1 int         // + "/"
}

// func dfsa388(s string, idx *int, depth int) int {
//     if *idx >= len(s) {
//         return 0
//     }
//     ans := 0
//     if s[*idx] == '\n' {
//         *idx++
//         for s[*idx] == '\t' {
//             *idx++
//         }

//     }

//     return ans
// }

// func main_LT0388_20220124() {
func main() {

    fmt.Println("ans:")

    // s := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
    // s := "a"
    s := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"

    ans := lengthLongestPath(s)

    fmt.Println(ans)

}

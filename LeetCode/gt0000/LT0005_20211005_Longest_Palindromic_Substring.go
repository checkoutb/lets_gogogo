// package sdq
package main

import (
    "fmt"
)

// 剩余长度小于 ans.length.
// 保存ans的 开始下标，长度， 减少string截取 。 不过，我这个是 切片吧？

// dp


// Runtime: 8 ms, faster than 71.23% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 2.8 MB, less than 75.02% of Go online submissions for Longest Palindromic Substring.
func longestPalindrome(s string) string {
    sz1 := len(s)
    ans := ""
    // mxlen := 0
    for i := 0; i < sz1; i++ {
        len2 := 1
        for j := 1; (i - j) >= 0 && (i + j) < sz1 && s[i - j] == s[i + j]; j++ {
            len2 += 2
        }
        if len2 > len(ans) {
            ans = s[i - len2 / 2 : i + len2 / 2 + 1]
        }
        if (i + 1) < sz1 && s[i + 1] == s[i] {
            len2 = 2
            for j := 1; (i - j) >= 0 && (i + j + 1) < sz1 && s[i - j] == s[i + j + 1]; j++ {
                len2 += 2
            }
            if (len2 > len(ans)) {
                // fmt.Println(len2)
                // fmt.Println(i)
                ans = s[i - len2 / 2 + 1 : i + len2 / 2 + 1]
            }
        }
    }
    return ans
}


func main_LT0005_20211005() {
// func main() {

    fmt.Println("ans:")

    // s := "babad"
    // s := "cbbd"
    // s := "a"
    // s := "ac"
    s := "aaaa"

    fmt.Println(longestPalindrome(s))
}

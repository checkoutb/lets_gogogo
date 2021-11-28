// package sdq
package main

import (
    "fmt"
)


// .. KMP





// def shortestPalindrome(self, s):
//     r = s[::-1]
//     for i in range(len(s) + 1):
//         if s.startswith(r[i:]):
//             return r[:i] + s
// 反转s， 找到最大的重叠部分  r[x : -1]  与 s[0 : x] 相等，说明r[0:x]就是 前缀。


// Runtime: 124 ms, faster than 64.71% of Go online submissions for Shortest Palindrome.
// Memory Usage: 3 MB, less than 97.06% of Go online submissions for Shortest Palindrome.
// 前面加char，使得 回文。
// 加的char 必然可以回文。
// 所以只要找到 使得 s[0 : x] 是一个回文，的最大 x
// O(n^2) ?      有快速的， Manacher 。。。 但是。。。
func shortestPalindrome(s string) string {
    sz1 := len(s)
    idx := -1
    for i := sz1 - 1; i >= 0; i-- {
        l, r := 0, i
        for l < r && s[l] == s[r] {
            l++
            r--
        }
        if l >= r {
            idx = i
            break
        }
    }
    pre := make([]byte, sz1 - idx - 1)
    for i := 0; i < len(pre); i++ {
        pre[i] = s[sz1 - 1 - i]
    }
    // fmt.Println(pre, idx)
    ans := string(pre) + s
    return ans
}


func main_LT0214_20211128() {
// func main() {

    fmt.Println("ans:")

    // s := "aacecaaa"
    s := "abcd"

    ans := shortestPalindrome(s)

    fmt.Println(ans)

}

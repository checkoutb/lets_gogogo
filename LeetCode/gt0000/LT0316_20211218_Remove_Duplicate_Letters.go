// package sdq
package main

import (
    "fmt"
)





// Runtime: 3 ms, faster than 46.88% of Go online submissions for Remove Duplicate Letters.
// Memory Usage: 2.7 MB, less than 18.75% of Go online submissions for Remove Duplicate Letters.
// 后面有的，外加最后一次出现。   hint， bitmask... 我想的是map/arr。。。
func removeDuplicateLetters(s string) string {
    t2 := 0
    for i := 0; i < len(s); i++ {
        t2 |= (1 << int(s[i] - 'a'))
    }
    cnt := 0
    for i := 0; i < 26; i++ {
        if t2 & (1 << i) != 0 {
            cnt++
        }
    }
    arr := make([]int, len(s))
    arr[len(s) - 1] = (1 << int(s[len(s) - 1] - 'a'))
    for i := len(s) - 2; i >= 0; i-- {
        arr[i] = arr[i + 1] | (1 << int(s[i] - 'a'))
    }
    ans := ""
    idx := 0
    // fmt.Println(arr)
    for len(ans) < cnt {
        AAA:
        for i := 0; i < 26; i++ {
            if t2 & (1 << i) != 0 {
                bt := byte(('a' + i))
                // fmt.Println(bt)
                for j := idx; j < len(s); j++ {
                    if s[j] == bt && (arr[j] & t2) == t2 {
                        idx = j + 1
                        ans += string(bt)
                        t2 &= (^(1 << i))              // not ~i..
                        // fmt.Println(ans, t2)
                        break AAA
                    }
                }
            }
        }
        // AAA:
        // continue
    }
    return ans
}


func main_LT0316_20211218() {
// func main() {

    fmt.Println("ans:")

    s := "bcabc"
    // s := "cbacdcbc"

    ans := removeDuplicateLetters(s)

    fmt.Println(ans)

}

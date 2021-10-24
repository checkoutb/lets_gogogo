// package sdq
package main

import (
    "fmt"
)


// dou cha bu duo. or ...


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Window Substring.
// Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Minimum Window Substring.
func minWindow(s string, t string) string {
    // arr := make([]int, 123)
    arr := [123]int{}
    ans := ""
    sza := 10000000
    cnt := 0
    for _, v := range t {
        if arr[v] == 0 {
            cnt++
        }
        arr[v]++
    }
    tag1 := -12345678
    for i, _ := range arr {
        if arr[i] == 0 {
            arr[i] = tag1
        }
    }
    // fmt.Println(arr)
    st := 0
    for i := 0; i < len(s); i++ {
        if arr[s[i]] == tag1 {continue}
        arr[s[i]]--
        if arr[s[i]] == 0 {
            cnt--
        }
        if cnt == 0 {
            // fmt.Println("1")
            for st <= i {
                if arr[s[st]] == tag1 {st++; continue}
                if arr[s[st]] == 0 {
                    t2 := i - st + 1
                    if t2 < sza {
                        ans = s[st : i + 1]
                        sza = i - st + 1
                    }
                    cnt++
                    arr[s[st]]++
                    st++
                    break
                }
                arr[s[st]]++
                st++
            }
        }
    }

    return ans
}


func main_LT0076_20211024() {
// func main() {

    fmt.Println("ans:")

    s, t := "ADOBECODEBANC", "ABC"
    // s, t := "a", "a"
    // s, t := "a", "aaa"

    ans := minWindow(s, t)

    fmt.Println(ans)
}

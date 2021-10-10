// package sdq
package main            // 一个文件夹只能一个包。。本来想的是，减少 加载的东西。

import (
    "fmt"
    "reflect"
)



// hist := map[rune]int{}
// for i, r := range s {
//     if old, ok := hist[r]; ok && old >= start {
//         start = old + 1
//     }
//     hist[r] = i
//     maxLen = max(maxLen, i - start + 1)
// }

// hash := make(map[rune]int)

// for idx, c := range str {
//     if _, okay := m[c]; okay == true && m[c] >= left {


// Runtime: 9 ms, faster than 52.16% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 2.8 MB, less than 84.93% of Go online submissions for Longest Substring Without Repeating Characters.
// no repeat
func lengthOfLongestSubstring(s string) int {
    st, _ := 0, 0
    ans := 0
    arr := [130]int{};           // last index ... map
    for i := 0; i < 130; i++ {
        arr[i] = -1
    }
    for i, ch := range s {
        if arr[ch] != -1 {
            t2 := arr[ch]
            for st <= t2 {
                arr[s[st]] = -1
                st++
            }
        }
        arr[ch] = i
        if (i - st + 1) > ans {
            ans = i - st + 1
        }
    }
    return ans
}


func main_LT0003_20211005() {
// func main() {

    // s := "abcabcbb"
    // s := "bbbbb"
    s := "pwwkew"
    // s := ""

    t2 := []rune(s)

    fmt.Println(s[2])
    // t2 rune = s[2]

    fmt.Println(reflect.TypeOf(s))
    fmt.Println(reflect.TypeOf(s[2]))
    fmt.Println(reflect.TypeOf(t2))
    fmt.Println(reflect.TypeOf(t2[2]))

    fmt.Println("ans:")

    fmt.Println(lengthOfLongestSubstring(s))

}

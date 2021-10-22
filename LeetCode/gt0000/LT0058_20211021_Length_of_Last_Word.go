// package sdq
package main

import (
    "fmt"
)


// arr := strings.Split(strings.Trim(s, " "), " ")
// return len(arr[len(arr)-1])


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
// Memory Usage: 2.3 MB, less than 28.81% of Go online submissions for Length of Last Word.
func lengthOfLastWord(s string) int {
    idx := len(s) - 1
    for s[idx] == ' ' {
        idx--
    }
    ans := 0
    for idx >= 0 && s[idx] != ' ' {
        idx--
        ans++
    }
    return ans
}


func main_LT0058_20211021() {
// func main() {

    fmt.Println("ans:")


}

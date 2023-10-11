// package sdq
package main

import (
    "fmt"
)

// D D

// str := make([]byte, len(words))
// for i := 0; i < len(words); i++ {
//     str[i] = words[i][0]
// }
// return string(str)==s



// Runtime11 ms
// Beats
// 12.53%
// Memory3.6 MB
// Beats
// 100%
func isAcronym(words []string, s string) bool {
    if len(s) != len(words) {
        return false;
    }
    for i := 0; i < len(s); i++ {
        if words[i][0] != s[i] {
            return false;
        }
    }
    return true;
}


func main_LT2828_20231010() {
// func main() {

    fmt.Println("ans:")


}

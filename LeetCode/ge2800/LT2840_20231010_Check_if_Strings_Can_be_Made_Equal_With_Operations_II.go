// package sdq
package main

import (
    "fmt"
)

// D D

// var odd [26]int
// var even [26]int

// for i := range s1{
//     if i % 2 == 0{
//         odd[s1[i]-'a']++





// Runtime8 ms
// Beats
// 95.74%
// Memory7 MB
// Beats
// 34.4%
func checkStrings(s1 string, s2 string) bool {
    vi := make([]int, 26)
    for i := 0; i < len(s1); i += 2 {
        vi[s1[i] - 'a']--
        vi[s2[i] - 'a']++
    }
    for _, cnt := range vi {
        if cnt != 0 {
            return false
        }
    }
    for i := 1; i < len(s1); i += 2 {
        vi[s1[i] - 'a']--
        vi[s2[i] - 'a']++
    }
    for _, cnt := range vi {
        if cnt != 0 {
            return false
        }
    }
    return true
}


func main_LT2840_20231010() {
// func main() {

    fmt.Println("ans:")


}

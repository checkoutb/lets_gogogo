package sdq
// package main

import (
    "fmt"
    // "slices"     // 1.20.4 not contain slices..
    // "golang.org/x/exp/slices"    // connect refuse..
)

// D D

// func rev[T any](aa []T) {
//     i := 0
//     j := len(aa)-1
//     for i < j {
//         aa[i], aa[j] = aa[j], aa[i]
//         i++
//         j--
//     }
// }
// 
// generic ?



// Runtime3 ms
// Beats
// 83.55%
// Memory3.1 MB
// Beats
// 66.45%
func finalString(s string) string {
    var vc []byte = []byte{}

    for _, ch := range s {
        if ch == 'i' {
            // slices.Reverse(vc);
            // var st, en = 0, vc.len()
            var st, en = 0, len(vc) - 1
            for st < en {
                vc[st], vc[en] = vc[en], vc[st]
                st, en = st + 1, en - 1
            }
        } else {
            vc = append(vc, byte(ch));            // ch is rune
        }
    }

    return string(vc);
}

func main_LT2810_20231010() {
// func main() {

    var s string = "string";

    fmt.Println("ans:", finalString(s));


}

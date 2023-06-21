// package sdq
package main

import (
    "fmt"
)


// Runtime0 ms
// Beats
// 100%
// Memory5.2 MB
// Beats
// 98.67%

func partitionString(s string) int {
    var vb = make([]bool, 26);
    var ans int = 1;
    for _, ch := range s {
        if (vb[ch - 'a']) {
            ans++;
            for i := range vb {
                vb[i] = false;
            }
        }
        vb[ch - 'a'] = true;
    }
    return ans;
}


func main_LT2405_20230620() {
// func main() {

    fmt.Println("ans:", partitionString("abacaba"));


}

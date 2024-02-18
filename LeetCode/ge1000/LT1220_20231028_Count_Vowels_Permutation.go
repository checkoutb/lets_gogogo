// package sdq
package main

import (
	"fmt"
)

// Runtime2 ms
// Beats
// 100%
// Memory1.9 MB
// Beats
// 100%
func countVowelPermutation(n int) int {
    a,e,i,o,u := 1,1,1,1,1
    const MOD = 1e9 + 7

    for n > 1 {
        n,a,e,i,o,u = n-1,((e+u)%MOD+i)%MOD,(a+i)%MOD,(e+o)%MOD,i,(o+i)%MOD
    }
    return ((a + e) % MOD + ((i + o) % MOD + u) % MOD) % MOD
}

func main_LT1220_20231028() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
	"fmt"
)

// D D

// var ret int
// for i:=0; i<=limit; i++ {            // first child got i
//     left := n-i
//     m := min(limit, left)
//     for j:=0; j<=m; j++ {            // second child got j
//         left2 := left-j
//         if left2<=limit {            // third child got left2
//             ret++
//         }
//     }
// }
// return ret



// Runtime8 ms
// Beats
// 35.19%
// Memory3.4 MB
// Beats
// 6.48%
func distributeCandies(n int, limit int) int {
	var vi []int = make([]int, n+1)
	vi[n] = 1
	var v2 []int
	for i := 0; i < 3; i++ { // 3 children
		v2 = make([]int, n+1)
		for j := 0; j <= n; j++ { // candy
			mxk := min(limit, j)
			for k := 0; k <= mxk; k++ {
				v2[j-k] += vi[j]
			}
		}
		vi = v2
		fmt.Println(vi)
	}
	return vi[0]
}

func main_LT2928_20231205() {
	// func main() {

	// a, b := 5, 2
	a, b := 3, 3

	fmt.Println("ans:", distributeCandies(a, b))

}

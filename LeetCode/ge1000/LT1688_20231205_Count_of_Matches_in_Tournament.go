// package sdq
package main

import (
	"fmt"
)

// Runtime1 ms
// Beats
// 70.31%
// Memory2 MB
// Beats
// 37.50%
func numberOfMatches(n int) int {
	var ans int = 0
	for n > 1 {
		ans += n / 2
		n = n/2 + n%2
	}
	return ans
}

func main_LT1688_20231205() {
	// func main() {

	fmt.Println("ans:")

}

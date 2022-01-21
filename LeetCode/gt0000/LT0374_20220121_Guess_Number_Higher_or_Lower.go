// package sdq
package main

import (
    "fmt"
)



// def guessNumber(self, n):
// class C: __getitem__ = lambda _, i: -guess(i)
// return bisect.bisect(C(), -1, 1, n)



// Runtime: 2 ms, faster than 28.03% of Go online submissions for Guess Number Higher or Lower.
// Memory Usage: 2 MB, less than 48.41% of Go online submissions for Guess Number Higher or Lower.
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
//
//
func guessNumber(n int) int {
    l, r := 1, n
    for l <= r {
        mid := l + (r - l) / 2              // 2^31
        t2 := guess(mid)
        if t2 == 0 {
            return mid
        } else if t2 < 0 {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return l
}

func main_LT0374_20220121() {
// func main() {

    fmt.Println("ans:")


}

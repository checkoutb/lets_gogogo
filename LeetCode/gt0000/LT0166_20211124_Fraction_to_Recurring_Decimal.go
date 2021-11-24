// package sdq
package main

import (
    "fmt"
    "strconv"
)


// fmt.Sprintf("%d", n)
// int -> string

// (numerator > 0) ^ (denominator > 0) ? "-" : ""


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Fraction to Recurring Decimal.
// Memory Usage: 3.1 MB, less than 11.76% of Go online submissions for Fraction to Recurring Decimal.
// [INT_MIN, INT_MAX] .. int is int64.
func fractionToDecimal(numerator int, denominator int) string {
    map2 := map[int]int{}
    b2 := (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0)
    if numerator < 0 {
        numerator = -numerator
    }
    if denominator < 0 {
        denominator = -denominator
    }
    t2 := numerator / denominator
    ans := ""
    if b2 {
        ans = "-"
    }
    ans += strconv.Itoa(t2)
    numerator -= t2 * denominator
    if numerator != 0 {
        ans += "."
        numerator *= 10
    }
    for numerator != 0 {
        if numerator < denominator {
            numerator *= 10
            ans += "0"
        } else {
            if _, ok := map2[numerator]; ok {
                t3 := map2[numerator]
                ans = ans[0 : t3] + "(" + ans[t3 : len(ans)] + ")"
                break
            } else {
                map2[numerator] = len(ans)
            }
            t2 := numerator / denominator
            numerator -= t2 * denominator
            ans += strconv.Itoa(t2)
            numerator *= 10
        }
    }
    return ans
}

func main_LT0166_20211124() {
// func main() {

    fmt.Println("ans:")

    // a, b := 2, 3
    // a, b := 4, 333
    a, b := -50, 8

    ans := fractionToDecimal(a, b)

    fmt.Println(ans)
}

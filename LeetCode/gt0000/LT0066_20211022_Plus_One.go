// package sdq
package main

import (
    "fmt"
)


// for (int i=digits.size(); i--; digits[i] = 0)                // .... i-- ; = 0
// if (digits[i]++ < 9)
//     return digits;
// digits[0]++;
// digits.push_back(0);
// return digits;


// Runtime: 4 ms, faster than 11.69% of Go online submissions for Plus One.
// Memory Usage: 2.2 MB, less than 42.37% of Go online submissions for Plus One.
func plusOne(digits []int) []int {
    carry, idx := 1, len(digits) - 1
    for carry != 0 && idx >= 0 {
        digits[idx] += carry
        carry = 0
        if digits[idx] > 9 {
            carry = 1
            digits[idx] %= 10
        }
        idx--
    }
    if carry == 1 {
        digits = append([]int{1}, digits...)
    }
    return digits
}

func main_LT0066_20211022() {
// func main() {

    fmt.Println("ans:")


}

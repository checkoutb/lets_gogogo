// package sdq
package main

import (
    "fmt"
)

// for {
//     res = a ^ b
//     rm = a&b << 1
//     if rm == 0 { break }
//     a = res
//     b = rm
// }

// if (ta ^ tb ^ t1)

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Two Integers.
// Memory Usage: 2 MB, less than 53.85% of Go online submissions for Sum of Two Integers.
func getSum(a int, b int) int {
    ans := 0
    bt := 1
    car := 0
    for a != 0 || b != 0 {
        // showBinary371(a)
        // showBinary371(b)
        // showBinary371(ans)
        // fmt.Println(car)
        b1, b2 := a & bt, b & bt
        a &= (^bt)
        b &= (^bt)
        t2 := 1
        if b1 != 0 { t2 <<= 1 }          // 2           // != 0  not > 0....
        if b2 != 0 { t2 <<= 1 }      // 4
        if car != 0 { t2 <<= 1 }     // 8
        car = 0
        if t2 >= (1 << 2) {     // 4
            car = 1
        }
        if t2 != (1 << 2) && t2 != 1 {                      // t2 != 1
            ans |= bt
        }
        bt <<= 1

        // fmt.Println(a, b, bt, ans, b1, b2, car)
    }
    // showBinary371(a)
    // showBinary371(b)
    // showBinary371(ans)
    // fmt.Println(car)
    if car > 0 {         // over-flow.bu...
        ans |= bt
    }
    return ans
}

func showBinary371(n int) {
    for i := 63; i >= 0; i-- {
        if (n & (1 << i)) == 0 {
            fmt.Print("0")
        } else {
            fmt.Print("1")
        }
        // fmt.Print(a & (1 << i), ", ")
    }
    fmt.Println()
}

func main_LT0371_20220120() {
// func main() {

    fmt.Println("ans:")

    // a, b := 1, 2
    // a, b := 2, 3
    // a, b := -1, 3
    // a, b := -1, -1111
    a, b := 4, 5

    // for i := 63; i >= 0; i-- {
    //     if (a & (1 << i)) == 0 {
    //         fmt.Print("0")
    //     } else {
    //         fmt.Print("1")
    //     }
    //     // fmt.Print(a & (1 << i), ", ")
    // }
    // fmt.Println()

    t2 := getSum(a, b)
    showBinary371(t2)
    fmt.Println(t2)

}

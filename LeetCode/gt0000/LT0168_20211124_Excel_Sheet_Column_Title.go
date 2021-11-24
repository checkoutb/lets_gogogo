// package sdq
package main

import (
    "fmt"
)






// return n == 0 ? "" : convertToTitle((n - 1) / 26) + (char) ((n - 1) % 26 + 'A');


// gg

/*
A - 1
B - 2

AA - 27 = A*Z + A
AZ - A*Z + Z = 26+26 = 52
BA - B*Z + A = 2*26+1 = 53
BZ - 2*26+26 = 78
ZA - 26*26+1 = 671
ZY - 26*26+25 = 701

Z 26        -  AA  26+1
ZZ 26*26+26     -  AAA  1*26*26+1*26+1
ZZZ 26*26*26 + 26*26 + 26    -  AAAA  1*26*26*26 + 1*2626 + 1*26 + 1

26
26*(26+1)
26*(26*26+26+1)     26*(26*(26+1) + 1)

a = 26*(a+1)

*/ 
func convertToTitle(columnNumber int) string {
    t2 := 26
    for t2 < columnNumber {
        t2 = (t2 + 1) * 26
    }
    t2 = t2 / 26 - 1
    ans := ""
    for columnNumber > 0 {
        if t2 == 0 {
            t2 = 1
        }
        t3 := columnNumber / t2
        if columnNumber % t2 == 0 && t2 != 1 {
            t3--
        }
        fmt.Println(t2, columnNumber, t3)
        columnNumber -= t3 * t2
        t2 = t2 / 26 - 1
        ans += string('A' + t3 - 1)
    }
    return ans
}


// error
// // world without 0..bu，应该还有一道 反过来的，更恐怖。。
// // 这个就是26 26×26 INT_MAX
// // 好像 0 是最大值。被整除的时候。。 这道也很难。。
// func convertToTitle(columnNumber int) string {
//     t2 := 26
//     for t2 < columnNumber {
//         t2 *= 26
//     }
//     t2 /= 26
//     ans := ""
//     for columnNumber > 0 {
//         t3 := columnNumber / t2

//         fmt.Println(t3, columnNumber, t2)

//         columnNumber -= t3 * t2
//         t2 /= 26
//         ans += string('A' + t3 - 1)
//     }
//     return ans
// }

func main_LT0168_20211124() {
// func main() {

    fmt.Println("ans:")

    // a := 1
    // a := 28
    // a := 701
    // a := 2147483647

    arr := []int{
        51, 
        52, 
        53,
        701,
        2147483647,
    }

    for i := 0; i < len(arr); i++ {
        ans := convertToTitle(arr[i])
        fmt.Println(ans)
    }


}

// package sdq
package main

import (
    "fmt"
)

// int value = b[0];
// for (int i = 1; i <= size; i++)
// {
//     value = (value * 10 + b[i]) % 1140;
// }
// 。。。

// 欧拉，费马，中国余数


// Runtime: 12 ms, faster than 63.16% of Go online submissions for Super Pow.
// Memory Usage: 4.2 MB, less than 15.79% of Go online submissions for Super Pow.
// a^b = a^(b/2) * a^(b/2)
// a^b%1337 = (a%1337)^b
// a^(10*x+y) = a^(10*x) * a^y
func superPow(a int, b []int) int {
    sz1 := len(b)
    arr := make([]int, sz1)
    a %= 1337
    idx := sz1 - 1
    arr[idx] = a            // a^1
    ans := 1
    idx--
    for idx >= 0 {
        arr[idx] = 1
        for i := 0; i < 10; i++ {
            arr[idx] = (arr[idx] * arr[idx + 1] % 1337)
        }
        idx--
    }
    for i := 0; i < sz1; i++ {
        for j := 0; j < b[i]; j++ {
            ans = ans * arr[i] % 1337
        }
    }
    return ans
}

func main_LT0372_20220121() {
// func main() {

    fmt.Println("ans:")


}

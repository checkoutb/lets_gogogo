// package sdq
package main

import (
    "fmt"
)



// if(n == 0) return 1;
// int res = 1, dp = 9, val = 9;
// for(int i = 2; i <= n; i++) 
//     res += dp, dp *= val, val--;
// return res+dp;
// .. i=2 的时候 dp 变成 9×9,  i=3 dp=9×9×8 。。。


// for i := 2; i <= n; i++ {
//     count *= 11 - i
//     res += count
// }


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Count Numbers with Unique Digits.
// Memory Usage: 2 MB, less than 37.14% of Go online submissions for Count Numbers with Unique Digits.
// 排列。 9取1,9取2,9取3...9取n      9取0 是1。。         9! / (9-m)!
// ... 10...  第一个不能是0 。。。
func countNumbersWithUniqueDigits(n int) int {
    ans := 0
    arr := make([]int, 11)
    arr[0] = 1          // idx!
    for i := 1; i <= 10; i++ {
        arr[i] = i * arr[i - 1]
    }

    // [0,1)   0
    for i := 0; i < n; i++ {
        ans += 9 * (arr[9]/arr[9 - i])
    }

    return ans + 1          // + 0
}


func main_LT0357_20220120() {
// func main() {

    fmt.Println("ans:")

    // n := 2
    n := 0

    fmt.Println(countNumbersWithUniqueDigits(n))

}

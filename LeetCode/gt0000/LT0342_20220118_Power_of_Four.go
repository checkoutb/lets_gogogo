// package sdq
package main

import (
    "fmt"
)




// return (n > 0) && (n&(n-1) == 0) && (n&0x55555555 != 0)
// ... 只有 1个 bit。 所以 直接 枚举所有 可能。。。

// return num > 0 && (num & (num - 1)) == 0 && (num - 1) % 3 == 0;
// (1) 4^n - 1 = (2^n + 1) * (2^n - 1)
// (2) among any 3 consecutive numbers, there must be one that is a multiple of 3
// (2^n-1), (2^n), (2^n+1), one of them must be a multiple of 3, and (2^n) cannot be the one, therefore either (2^n-1) or (2^n+1) must be a multiple of 3
// 3个连续数 必然有一个是 3的倍数。。。。。。。。。。。。

// double dou = log(num) / log(4);


// Runtime: 5 ms, faster than 9.20% of Go online submissions for Power of Four.
// Memory Usage: 2.3 MB, less than 29.89% of Go online submissions for Power of Four.
// >=4 && (x&(4)!=0 || x&(4^2)!= 0 || ...)      // 2^31 4^30...
// x&(x-1) == 0 &&  ... 怎么获得 最右 1bit的位置？ 我记得好像有
// 4^0 == 1
func isPowerOfFour(n int) bool {
    if n < 1 || (n & (n - 1) != 0) {
        return false
    }
    for n > 1 {
        n >>= 2
    }
    return n == 1
}


func main_LT0342_20220118() {
// func main() {

    fmt.Println("ans:")


}

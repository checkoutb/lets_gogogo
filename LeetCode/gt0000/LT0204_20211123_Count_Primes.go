// package sdq
package main

import (
    "fmt"
)


// nonPrimes := make([]bool, n)
// for p := 2; p*p < n; p++ {
//     if !nonPrimes[p] {
//         for i := p*p; i < n; i += p {
//             nonPrimes[i] = true
//         }
//     }
// }
// count := 0
// for i:=2; i< len(nonPrimes); i++ {
//     if !nonPrimes[i] {
//         count++
//     }
// }


// Runtime: 141 ms, faster than 22.22% of Go online submissions for Count Primes.
// Memory Usage: 7.5 MB, less than 99.52% of Go online submissions for Count Primes.
// 好像只能 拿个表 来筛选啊。 不可能for上去吧。
// 还有， 减去2的倍数，减去3的倍数 +6个倍数，但是。。。
func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }
    sz1 := n >> 1
    arr := make([]bool, sz1)
    ans := 1
    for i := 1; i < sz1; i++ {
        if arr[i] == false {
            ans++
            t3 := i * 2 + 1     // 1-3 2-5 3-7
            // fmt.Println(t3)
            for t2 := t3; t2 < sz1 << 1; t2 += t3 {
                if t2 & 1 == 1 {
                    arr[t2 >> 1] = true
                }
                // fmt.Println(" , ", t2, t2 >> 1)
            }
        }
    }
    // fmt.Println(arr)
    return ans
}


// func main_LT0204_20211123() {
func main() {

    fmt.Println("ans:")

    n := 10

    ans := countPrimes(n)

    fmt.Println(ans)

}

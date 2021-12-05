// package sdq
package main

import (
    "fmt"
)

// for _, num := range [3]int{2, 3, 5} {
//     for n != 0 && n%num == 0 {
//         n = n / num
//     }
// }


// for (int i=2; i<6 && num; i++)
//     while (num % i == 0)
//         num /= i;
// return num == 1;
// 2，3，4，5，   4的倍数会被2除掉。。


// 就算不是质数 也是硬算，只要能乘起来，就肯定能除起来。
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Ugly Number.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Ugly Number.
// 质数，应该是硬算吧？
func isUgly(n int) bool {
    if n <= 0 {
        return false
    }
    for n % 5 == 0 {
        n /= 5
    }
    for n % 3 == 0 {
        n /= 3
    }
    for n % 2 == 0 {
        n /= 2
    }
    // lst := n - 1
    // for lst != n {
    //     lst = n
    //     for n % 5 == 0 {
    //         n /= 5
    //     }
    //     for n % 3 == 0 {
    //         n /= 3
    //     }
    //     for 
    // }
    return n == 1
}

func main_LT0263_20211129() {
// func main() {

    fmt.Println("ans:")


}

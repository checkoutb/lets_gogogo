// package sdq
package main

import (
    "fmt"
)


// while(fast != 1 && slow != fast) {
//     slow = getNext(slow);
//     fast = getNext(getNext(fast));
// }
// 这也是找环。。。。


// Floyd's Cycle detection algorithm       这个就是快慢指针
// Brent's Cycle detection algorithm


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
// Memory Usage: 2.2 MB, less than 21.17% of Go online submissions for Happy Number.
// 1
// 10  3^2 + 1^2
// 是不是所有的 100..00 都可以变成2个数的平方 的和？
// 2 4 16 37   58   99 162 41 17 50 25 29 85(==58)
func isHappy(n int) bool {
    map2 := map[int]bool{}
    for true {
        if _, ok := map2[n]; !ok {
            map2[n] = true
            t2 := 0
            for n > 0 {
                t2 += (n % 10) * (n % 10)
                n /= 10
            }
            n = t2
        } else {
            break
        }
    }
    return n == 1
}


func main_LT0202_20211123() {
// func main() {

    fmt.Println("ans:")


}

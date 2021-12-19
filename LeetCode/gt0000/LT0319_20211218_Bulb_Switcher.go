// package sdq
package main

import (
    "fmt"
)


// func bulbSwitch(n int) int {
//     return int(math.Sqrt(float64(n)))
// }



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Bulb Switcher.
// Memory Usage: 1.9 MB, less than 37.84% of Go online submissions for Bulb Switcher.
// n盏灯，n次循环
// 1 1  ->  1
// 2 2  11 10 -> 1
// 3  111 101 100 -> 1
// 4  1111 1010  1000 1001 -> 2
// 5 11111 10101 10001 10011 10010  - 2
// 6 111111 101010  100011 100111 100101 100100
// sqrt?
// 7 1111111 1010101 1000111 1001111  1001011 1001001 1001000
// 11111111 10101010  10001110 10011111 1001000
// 111111111 101010101 100011100 100111110 
// 后一半 round，每次round，影响一个。
// ok， 因为  因式分解， 只有 当数是 某个数的平方时， 因子总数才是 奇数。 奇数是亮。
func bulbSwitch(n int) int {
    ans := 0
    for i := 1; i * i <= n; i++ {
        ans++
    }
    return ans
}


func main_LT0319_20211218() {
// func main() {

    fmt.Println("ans:")


}

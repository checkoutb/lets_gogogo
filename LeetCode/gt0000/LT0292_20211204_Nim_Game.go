// package sdq
package main

import (
    "fmt"
)



// boolean[] results = {true,true,true,false};
// for (int i=4; i<n; i++){
//     results[i%4] = !(results[(i-1)%4]&&results[(i-2)%4]&&results[(i-3)%4]);
// }
// return results[(n-1)%4];
// 由于n < 21亿， 所以 tle的。 ???


// ... 我拿了以后  留下4个 就赢了。  所以 第一次 就看能不能修成 4的倍数， 然后 每次 对方拿以后，我修成4。   如果最开始就是4n 就输了。
func canWinNim(n int) bool {
    return n % 4 != 0
}

func main_LT0292_20211204() {
// func main() {

    fmt.Println("ans:")


}

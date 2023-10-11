// package sdq
package main

import (
    "fmt"
)


// D D

// x := purchaseAmount % 10
// if x >= 5 {
//     purchaseAmount = purchaseAmount - x + 10
// } else {
//     purchaseAmount -= x
// }


// Runtime1 ms
// Beats
// 82.67%
// Memory1.9 MB
// Beats
// 96%
// 5->10
func accountBalanceAfterPurchase(purchaseAmount int) int {
    if purchaseAmount % 10 >= 5 {
        purchaseAmount = (purchaseAmount + 10) / 10 * 10;
    } else {
        purchaseAmount = (purchaseAmount) / 10 * 10;
    }
    return 100 - purchaseAmount;
}


func main_LT2806_20231010() {
// func main() {

    fmt.Println("ans:")


}

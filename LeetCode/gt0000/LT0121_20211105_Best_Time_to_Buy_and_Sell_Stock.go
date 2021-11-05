// package sdq
package main

import (
    "fmt"
)





// int maxCur = 0, maxSoFar = 0;
// for(int i = 1; i < prices.length; i++) {
//     maxCur = Math.max(0, maxCur += prices[i] - prices[i-1]);
//     maxSoFar = Math.max(maxCur, maxSoFar);
// }
// return maxSoFar;
// Kadane


// Runtime: 124 ms, faster than 61.93% of Go online submissions for Best Time to Buy and Sell Stock.
// Memory Usage: 8.4 MB, less than 49.43% of Go online submissions for Best Time to Buy and Sell Stock.
func maxProfit(prices []int) int {
    sz1 := len(prices)
    mx := prices[sz1 - 1]
    ans := 0
    for i := sz1 - 2; i >= 0; i-- {
        if prices[i] > mx {
            mx = prices[i]
        } else {
            if mx - prices[i] > ans {
                ans = mx - prices[i]
            }
        }
    }
    return ans
}


func main_LT0121_20211105() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
    "fmt"
)


// int n = A.length, lastBuy = -A[0], lastSold = 0;
// if (n == 0) return 0;                       
// for (int i = 1; i < n; i++) {
//     int curBuy = Math.max(lastBuy, lastSold - A[i]);
//     int curSold = Math.max(lastSold, lastBuy + A[i]);
//     lastBuy = curBuy;
//     lastSold = curSold;
// }

// int i = 0, buy, sell, profit = 0, N = prices.length - 1;
// while (i < N) {
//     while (i < N && prices[i + 1] <= prices[i]) i++;
//     buy = prices[i];
//     while (i < N && prices[i + 1] > prices[i]) i++;
//     sell = prices[i];
//     profit += sell - buy;
// }



// Runtime: 4 ms, faster than 90.73% of Go online submissions for Best Time to Buy and Sell Stock II.
// Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Best Time to Buy and Sell Stock II.
// 如果 大于前面 就 赚， 小于前面 就不操作。
func maxProfit(prices []int) int {
    ans := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i - 1] {
            ans += prices[i] - prices[i - 1]
        }
    }
    return ans
}


func main_LT0122_20211108() {
// func main() {

    fmt.Println("ans:")


}

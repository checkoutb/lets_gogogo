// package sdq
package main

import (
    "fmt"
)




// for _, i := range prices { 
//     release2 = max(release2, hold2+i) 
//     hold2 = max(hold2, release1-i)    
//     release1 = max(release1, hold1+i) 
//     hold1 = max(hold1, -i)            
// }



// states[][0]: one buy
// states[][1]: one buy, one sell
// states[][2]: two buys, one sell
// states[][3]: two buy, two sells
//
// int states[2][4] = {INT_MIN, 0, INT_MIN, 0}; // 0: 1 buy, 1: one buy/sell, 2: 2 buys/1 sell, 3, 2 buys/sells
// states[next][0] = max(states[cur][0], -prices[i]);
// states[next][1] = max(states[cur][1], states[cur][0]+prices[i]);
// states[next][2] = max(states[cur][2], states[cur][1]-prices[i]);
// states[next][3] = max(states[cur][3], states[cur][2]+prices[i]);
// swap(next, cur);


// dp[k, i] = max(dp[k, i-1], prices[i] - prices[j] + dp[k-1, j-1]), j=[0..i-1]


// release2 = Math.max(release2, hold2+i);     // The maximum if we've just sold 2nd stock so far.
// hold2    = Math.max(hold2,    release1-i);  // The maximum if we've just buy  2nd stock so far.
// release1 = Math.max(release1, hold1+i);     // The maximum if we've just sold 1nd stock so far.
// hold1    = Math.max(hold1,    -i);          // The maximum if we've just buy  1st stock so far. 



// Runtime: 112 ms, faster than 91.76% of Go online submissions for Best Time to Buy and Sell Stock III.
// Memory Usage: 9.6 MB, less than 55.69% of Go online submissions for Best Time to Buy and Sell Stock III.
// 2次买卖
// 前面执行的一次买卖 的最大收益  +  现在(当前下表)开始第二次 买入 
// 第一次买卖就是 记录 最小值，然后用 当前值减去它， 就是 第一次买卖的 收益，然后 if > 1st_MAX_Profit {}
// 第二次买卖也是类似，不，好像复杂度上去了。 主要是每个下标 都需要 作为开始。。 不  好像可以 先从后往前计算 第二次的(记录最大值)， 然后再从前往后 计算第一次的。
func maxProfit123(prices []int) int {
    ans := 0
    sz1 := len(prices)

    arr := make([]int, sz1)         // 第二次   保存的就是 这天买的话，最大利润 , 不，要后续的最大
    mx := prices[sz1 - 1]
    mx2 := 0
    for i := sz1 - 2; i >= 0; i-- {
        if prices[i] >= mx {
            mx = prices[i]
            arr[i] = mx2
        } else {
            if mx - prices[i] > mx2 {
                mx2 = mx - prices[i]
            }
            arr[i] = mx2
        }
    }
    fmt.Println(arr)
    mn := prices[0]
    for i := 1; i < sz1; i++ {
        if prices[i] > mn {
            if prices[i] - mn + arr[i] > ans {
                ans = prices[i] - mn + arr[i]
            }
        } else {
            mn = prices[i]
        }
    }
    return ans
}


func main_LT0123_20211108() {
// func main() {

    fmt.Println("ans:")

    arr := []int{6,1,3,2,4,7}       // 7

    ans := maxProfit123(arr)

    fmt.Println(ans)
}

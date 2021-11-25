// package sdq
package main

import (
    "fmt"
)

// 。。 标准的dp。。

// for tr := 1; tr <= k; tr++ {
//     cache := make([]int, len(prices))
//     maxDiff := 0
//     for i, currPrice := range prices {
//         if i == 0 {
//             cache[i] = 0
//             maxDiff = memo[tr-1][i] - currPrice
//             continue
//         }
//         profit := currPrice + maxDiff
//         cache[i] = int(math.Max(float64(profit), float64(cache[i-1])))
//         maxDiff = int(math.Max(float64(maxDiff), float64(memo[tr-1][i] - currPrice)))
//     }
//     memo[tr] = cache
// }
// cache是数组， memo就是 第几次交易，在第几个价格之前 的 最大profit。


/**
 * dp[i, j] represents the max profit up until prices[j] using at most i transactions. 
 * dp[i, j] = max(dp[i, j-1], prices[j] - prices[jj] + dp[i-1, jj]) { jj in range of [0, j-1] }
 *          = max(dp[i, j-1], prices[j] + max(dp[i-1, jj] - prices[jj]))
 * dp[0, j] = 0; 0 transactions makes 0 profit
 * dp[i, 0] = 0; if there is only one price data point you can't make any transaction.
 */
//


//   int maxProfit(vector<int>& p, int res = 0) {
//     for (int i = 1; i < p.size(); ++i) res += max(0, p[i] - p[i - 1]);
//     return res;
//   }
//   int maxProfit(int k, vector<int>& prices) {
//     if (k >= prices.size() / 2) return maxProfit(prices);
//     vector<int> buys(k + 1, INT_MIN), sells(k + 1, 0);
//     for (auto p : prices)
//       for (auto i = 1; i <= k; ++i) {
//         buys[i] = max(buys[i], sells[i - 1] - p);
//         sells[i] = max(sells[i], buys[i] + p);  
//       }
//     return sells[k];
//   }




// Runtime: 10 ms, faster than 21.92% of Go online submissions for Best Time to Buy and Sell Stock IV.
// Memory Usage: 2.6 MB, less than 49.32% of Go online submissions for Best Time to Buy and Sell Stock IV.
// 最多k次交易，并且交易不能重叠。
// 循环k？  第一次只需要 min  max
// 第二次，是在第一次的基础上，再for-for。太慢了。好像只需要for一次。
// 但是第三次就。。2+1 还是 1+2 。 第四次是3+1 还是 1+3 还是 2+2
// 还是说 把所有赚钱的交易列举，然后 合并下来满足 k
// 应该是合并;     循环k，第二次并不是基于第一次的。
func maxProfit(k int, prices []int) int {
    sz1 := len(prices)
    arr := [][]int{}
    for i := 0; i < sz1 - 1; i++ {
        if prices[i] < prices[i + 1] {
            t2 := i
            i++
            for i < sz1 && prices[i] > prices[i - 1] {
                i++
            }
            i--
            arr = append(arr, []int{t2, i})
        }
    }
    fmt.Println(arr)
    for len(arr) > k {
        // 合并或者删除。  2-11,2-11,10-22 -> 2-11,2-22    2-11,2-11,5-10 -> 2-11,2-11
        // 合并是和后面的一起计算， 移除是看 最小利润。
        // 不不不，实际上是 一件事的2个方面， 如果 后一次的卖出 大于 本次的卖出，那么肯定合并， 否则就是尝试移除。  或者移除自己。。
        // 11-12,1-22,1-22 -> 1-22,1-22
        // 11-15,1-22,1-22,21-23 -> 11-15,1-22,1-23
        // 还是说 就是移除最小利润，然后 看被移除的交易的  买入 和 卖出价格  是否能被 后一次 ，前一次 交易 使用。
        // 1-22,5-24,5-6 -> 1-22,5-24
        // 5-11,6-12,5-7 -> 5-12,5-7   不是移除最小利润。 .... ，，，，这里错了，，就是最小利润。。感觉下面 合并的白写了。。
        // 最小利润， 前向合并导致减少的利润(前一个的利润+本次的利润-合并后的利润)， 后向合并导致减少的利润。
        dec := 10000000
        rmidx := -1
        tag := -1
        for i := 0; i < len(arr); i++ {
            t1, t2 := arr[i][0], arr[i][1]
            if prices[t2] - prices[t1] < dec {
                rmidx = i
                dec = prices[t2] - prices[t1]
                tag = 0
            }
            if i > 0 {
                t3, t4 := arr[i - 1][0], arr[i - 1][1]
                a := prices[t2] - prices[t1] + prices[t4] - prices[t3] - (prices[t2] - prices[t3])          // .
                if a < dec {
                    rmidx = i
                    dec = a
                    tag = 1
                }
                // if prices[t4] > prices[t2] {        // 卖出价大
                // }
            }
            if i < len(arr) - 1 {
                t3, t4 := arr[i + 1][0], arr[i + 1][1]
                a := prices[t2] - prices[t1] + prices[t4] - prices[t3] - (prices[t4] - prices[t1])
                if a < dec {
                    rmidx = i
                    dec = a
                    tag = 2
                }
            }
        }
        switch tag {
        // case 0:
        //     arr = arr[0 : rmidx] + arr[rmidx + 1 : ]
        case 1:
            arr[rmidx - 1][1] = arr[rmidx][1]
        case 2:
            arr[rmidx + 1][0] = arr[rmidx][0]
        }
        // arr = arr[0 : rmidx] + arr[rmidx + 1 : ]     /// ...no operator+ for slice..
        arr = append(arr[0 : rmidx], arr[rmidx + 1 : ]...)
    }
    ans := 0
    for _, v := range arr {
        ans += prices[v[1]] - prices[v[0]]
    }
    return ans
}

func main_LT0188_20211125() {
// func main() {

    fmt.Println("ans:")

    // k := 2
    // arr := []int{2,4,1}

    k := 2
    // arr := []int{3,2,6,5,0,3,1,111}
    arr := []int{5,11,6,12,5,7}

    ans := maxProfit(k, arr)

    fmt.Println(ans)

}

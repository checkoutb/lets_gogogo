// package sdq
package main

import (
    "fmt"
)



// gg



// dp
// 还有一种， 肯定是低谷买，高峰卖， 所以买的时候 判断下，是 早一天卖，还是晚一天买。 哪种 利润降低的少。 如果 低谷的前一天不是 高峰，可以直接买。 如果前一天是高峰，就要判断 是早卖，还是晚买
// dp更难啊。
func maxProfit(prices []int) int {
    sz1 := len(prices)
    arr1 := make([]int, sz1)        // buy
    arr2 := make([]int, sz1)        // sell
    arr1[0], arr1[1] = -prices[0], -prices[1]
    arr2[1] = prices[1] - prices[0]
    if arr2[1] < 0 {
        arr2[1] = 0
    }
    for i := 2; i < len(prices); i++ {          // 今天买，那么2种， 今天价格比昨天更低，那么 昨天的buy+昨天的价格-今天的价格
                                            // 更高，那么就是 昨天的利润。 或者 前天卖 + 昨天买的价格。 不不不，都需要max下。 要不要昨天的利润
                                            // 今天卖， 今天的价格比昨天更高。  或者 前天买，今天卖。
        arr1[i] = arr2[i - 2] - prices[i]
        if arr1[i - 1] + prices[i - 1] - prices[i] > arr1[i] {
            arr1[i] = arr1[i - 1] + prices[i - 1] - prices[i]
        }

        arr2[i] = arr1[i - 2] + prices[i]
        if arr2[i - 1] + prices[i] - prices[i - 1] > arr2[i] {
            arr2[i] = arr2[i - 1] + prices[i] - prices[i - 1]
        }
        fmt.Println(arr1, " - ", arr2)
    }
    return arr2[sz1 - 1]


    // ppb, pps, pb, ps := 0,0,0,0
    // ans := 0

    // for i := 0; i < len(prices); i++ {

    // }

    // return ans
}

func main_LT0309_20211209() {
// func main() {

    fmt.Println("ans:")

    arr := []int{1,2,3,0,2}

    fmt.Println(maxProfit(arr))

}

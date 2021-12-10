// package sdq
package main

import (
    "fmt"
)



// https://leetcode.com/problems/burst-balloons/discuss/76228/Share-some-analysis-and-explanations

/*
我们只知道 第一个扎的，和最后一个扎的 气球的coin， 第一个扎的就是[i-1]*[i]*[i+1], 最后一个扎的就是[j]
考虑n个气球，把它们作为 最后一个扎的。  这样就分成了2部分 并且边界清晰  比如扎[i]， 那么 就是 [0:i-1] + [i] + [i+1:] ,并且 [0:i-1]的边界是 1 和 [i],  [i+1:]的边界是[i] 和 1
有边界以后，就是独立的问题了。
*/

// nums[0] = nums[n++] = 1;          原数组也扩展边界了，  所以 left right 就是边界， 所以可以 [left]*[i]*[right]，边界不会被扎破。
// for (int k = 2; k < n; ++k)
//     for (int left = 0; left < n - k; ++left) {
//         int right = left + k;
//         for (int i = left + 1; i < right; ++i)
//             dp[left][right] = Math.max(dp[left][right], 
//             nums[left] * nums[i] * nums[right] + dp[left][i] + dp[i][right]);
//     }
// ...


// 。。leetcodeccpp 的 编码格式。。好像只有0312和0313？  ISO-8859 ... 其他的格式 是 ASCII  UTF-8..

// tle。。  [8,2,6,8,9,8,1,4,1,5,3,0,7,7,0,4,2,2,5]    还有0。。0肯定先扎的。
func maxCoins(nums []int) int {
    map2 := map[string]int{}
    ans := dfsa312(nums, map2)
    fmt.Println(map2)
    return ans
}

func dfsa312(nums []int, map2 map[string]int) int {
    if len(nums) == 0 {
        return 0
    }
    k := fmt.Sprint(nums)
    // fmt.Println(k)
    if _, ok := map2[k]; ok {
        return map2[k]
    }
    mx := -1
    for i := 0; i < len(nums); i++ {
        t2 := nums[i]
        if i > 0 { t2 *= nums[i - 1] }
        if i < len(nums) - 1 { t2 *= nums[i + 1] }
        // fmt.Println(nums, i, nums[0 : i], nums[i + 1 : ])
        // arr := append(nums[0 : i], nums[i + 1 : ]...)           // 切片指向了相同的数组，所以导致这里nums被修改了。 应该是的。。
        // fmt.Println(arr, nums, i)
        arr := make([]int, len(nums))
        copy(arr, nums)
        arr = append(arr[ : i], arr[i + 1 : ]...)
        t2 += dfsa312(arr, map2)
        if t2 > mx {
            mx = t2
        }
    }
    map2[k] = mx
    return mx
}



// ...dfs + memo ? 把array转成string，memo起来？
// 长度1 -> sz1 ?
func maxCoins2(nums []int) int {

    // discuss: O(n^3) ... 那就是3个for 。。 算了， 还是 dfs+memo, 来个tle吧。

    return -1


    // // [9,76,64,21,97,60]
    // // sz1 := len(nums)
    // if len(nums) == 1 {
    //     return nums[0]
    // }
    // ans := 0
    // for len(nums) > 2 {
    //     mni := 1
    //     for i := 2; i < len(nums) - 1; i++ {
    //         if nums[i] < nums[mni] {
    //             mni = i
    //         }
    //     }
    //     ans += nums[mni - 1] * nums[mni] * nums[mni + 1]
    //     // nums = nums[ : mni] +
    //     nums = append(nums[ : mni], nums[mni + 1 : ]...)
    // }
    // ans += nums[0] * nums[1]
    // if nums[1] > nums[0] {
    //     ans += nums[1]
    // } else {
    //     ans += nums[0]
    // }
    // return ans


    // sz1 := len(nums)
    // // arr := [][]int{}            // 以i为起始 长度(j+1..算了) 的 array 的 maxCoin    只是一个 类似 上三角矩阵，只有一半。map 不管了
    // arr := make([][]int, sz1)
    // for i, _ := range arr {
    //     arr[i] = make([]int, sz1 + 1)
    // }
    // for i := 0; i < sz1; i++ {
    //     arr[i][1] = nums[i]
    // }
    // for l := 2; l <= sz1; l++ {
    //     for i := 0; i + l <= sz1; i++ {
    //         // 不行，删除一个后，合并后的数组是新的，不能dp啊。
    //         // 3,1,5,8
    //         // 3 1 5 8
    //         // (31) (15) (58)
    //         // (315) (158)            358

    //         // 2侧不动， 移除最小的？
    //     }
    // }
}


// func main_LT0312_20211210() {
func main() {

    fmt.Println("ans:")

    // arr := []int{3,1,5,8}
    arr := []int{1,5}

    // s := fmt.Sprint(arr)

    // fmt.Println(s)


    fmt.Println(maxCoins(arr))

}

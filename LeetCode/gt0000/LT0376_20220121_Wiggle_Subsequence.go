// package sdq
package main

import (
    "fmt"
)




// int up = 1;
// int down = 1;
// for(int i =1;i<n;i++){
//     if(nums[i] < nums[i-1]){
//         down = up + 1;
//     }else if(nums[i] > nums[i-1]){
//         up = down + 1;
//     }
// }
// return max(up,down);




// up position, it means nums[i] > nums[i-1]
// down position, it means nums[i] < nums[i-1]
// equals to position, nums[i] == nums[i-1]
//
// for(int i = 1 ; i < nums.length; i++){
//     if( nums[i] > nums[i-1] ){
//         up[i] = down[i-1]+1;
//         down[i] = down[i-1];
//     }else if( nums[i] < nums[i-1]){
//         down[i] = up[i-1]+1;
//         up[i] = up[i-1];
//     }else{
//         down[i] = down[i-1];
//         up[i] = up[i-1];
//     }
// }



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Wiggle Subsequence.
// Memory Usage: 2.3 MB, less than 32.88% of Go online submissions for Wiggle Subsequence.
// 前面 最近的 一个 小于自己的数。  或者放弃。。。。不，前面的 小于自己的 但是 长度最长的 + 1 就是 自己的长度
// 小于 或 大于， 所以2次。
// 扫描线 也可以。  主要是 小于自己 长度最长 。 用什么结构。 感觉只能 按值排序， 然而二分搜索，然后向前遍历。
// 扫描线 应该更简单，[0]作为第一个元素， 往后找 比它 大的， 如果 一直大，那么 就一直找， 直到 一个 下降趋势。 那么长度就是 3了。 找不到下降，那么 是2。
// 下降趋势 一直降，直到 上升。 然后上升趋势。
func wiggleMaxLength(nums []int) int {
    ans := 0
    t2 := 0

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i - 1] {
            nums = append(nums[0 : i], nums[i + 1 : ]...)
            i--
        }
    }

    // map2 := map[int]bool{}
    // map2[nums[0]] = true
    // map2[nums[len(nums) - 1]] = true
    for i := 1; i < len(nums) - 1; i++ {
        if nums[i] > nums[i - 1] && nums[i] > nums[i + 1] {
            t2++
        } else if nums[i] < nums[i - 1] && nums[i] < nums[i + 1] {
            t2++
        }
        // map2[nums[i]] = true
    }
    ans = t2 + 2
    if ans > len(nums) {
        ans = len(nums)
    }
    // if ans > len(map2) {            // 1,1,1,1,1,1,1..error... 1,2,1,2,1,2,1,2  len(map) == 2....
    //     ans = len(map2)
    // }
    return ans

    // // up
    // len := 1
    // up := true
    // val := nums[0]
    // for i := 1; i < len(nums); i++ {
    //     t2 := nums[i]
    //     if t2 > val {
    //         if up {     // 上升趋势，且本次大于 上一个。。 不需要val。  而且 好像是 直接 计算 几个 拐点就可以了。。。
    //             val = t2
    //         } else {
    //             len++
    //             val = t2
    //             up = false
    //         }
    //     } else if t2 < val {
    //         if up {

    //         } else {

    //         }
    //     }
    // }
    // ans = len

    // // down
    // len := 1
    // up = false
    // val = nums[0]
    // for i := 1; i < len(nums); i++ {

    // }
    // if len > ans { ans = len }
    // return ans
}


func main_LT0376_20220121() {
// func main() {

    fmt.Println("ans:")


    // 67
    arr := []int{33,53,12,64,50,41,45,21,97,35,47,92,39,0,93,55,40,
        46,69,42,6,95,51,68,72,9,32,84,34,64,6,2,26,98,3,43,30,60,3,
        68,82,9,97,19,27,98,99,4,30,96,37,9,78,43,64,4,65,30,84,90,87,
        64,18,50,60,1,40,32,48,50,76,100,57,29,63,53,46,57,93,98,42,80,
        82,9,41,55,69,84,82,79,30,79,18,97,67,23,52,38,74,15}

    ans := wiggleMaxLength(arr)

    fmt.Println(ans, len(arr))

}

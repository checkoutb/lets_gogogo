// package sdq
package main

import (
    "fmt"
)


// 不同语言的差别好奇怪。这道题目，java 基本 2-4ms，  go要 120ms。。。 但是 其他题目 go 和java 是差不多的啊。
// 难道go 对于这种 类型的 特别慢？  还是 这道题目 提交到的 那个go 执行器  比其他题目的 执行器 性能低？  不同的题目 应该是同一个go 执行器吧？

// for(int i=0; i<nums.length;) {
//     if(nums[i] > 0 && nums[i] <= nums.length && nums[nums[i]-1] != nums[i]) {
//         int temp = nums[nums[i]-1];
//         nums[nums[i]-1] = nums[i];
//         nums[i] = temp;
//     }
//     else {
//         i++;
//     }
// }
// 。。。。 直接交换 到 合适的位置上。  for 保证 每个元素都被计算。




// Runtime: 188 ms, faster than 51.79% of Go online submissions for First Missing Positive.
// Memory Usage: 27.4 MB, less than 25.73% of Go online submissions for First Missing Positive.

// nums 最长5万， 所以 肯定是5万以内的。。  或者说，肯定是 长度以内的。 不过是 O(n) 空间了。。
// hint: Can you apply that logic to the existing space?  。。。 就是 要原地 放入slot中。  ok。
// 小于0, >= sz1  放弃。 需要特殊值，5000000代表存在，500000001代表不存在。
// 如果值小于idx，设置下，[值]存在 且 当前下标为不存在
// 如果 大于idx。小于 sz1,  和 [值]的值 互换。 i--。  设置[zhi]为存在。
// 如果大于sz1，值放弃，当前下标不存在。
// 特殊值 不行的。
// 应该用v 和 i 比较，不然太慢了。
func firstMissingPositive(nums []int) int {
    idx := 0
    sz1 := len(nums)
    // b2 := false
    for i := 0; i < sz1; i++ {
        v := nums[i] - 1
        if v >= sz1 || v < 0 {
            continue
        } else if v <= i {
            for idx <= v {
                nums[idx] = 0
                idx++
            }
            nums[v] = 1
        } else {
            if nums[i] == nums[v] {
                continue
            } else {
                nums[i], nums[v] = nums[v], nums[i]
                i--
            }
            // for i2 := v; i2 < sz1; i2++ {

            // }
            // nums[i], nums[v] = nums[v], nums[i]        // [i] == [v]
            // i--
        }

        // v := nums[i]
        // if v > sz1 || v <= 0 {
        //     // if v == sz1 {
        //     //     b2 = true
        //     // }
        //     continue
        // } else if v <= i + 1 {
        //     for idx < v {
        //         nums[idx] = 0
        //         idx++
        //     }
        //     nums[v - 1] = 1
        //     // idx++
        // } else {
        //     nums[i], nums[v] = nums[v], nums[i]
        //     i--
        // }
    }
    fmt.Println(nums, ", ", idx)
    if idx == 0 {
        return 1
    }
    for i := 0; i < idx; i++ {          // i < idx
        if nums[i] == 0 {           // == 0
            return i + 1
        }
    }
    // if b2 {
    //     return sz1 + 1
    // }
    return idx + 1
}


func main_LT0041_20211019() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{1,2,0}
    // arr := []int{3,4,-1,1}
    // arr := []int{7,8,9,11,12}
    // arr := []int{1}
    arr := []int{2,2}

    ans := firstMissingPositive(arr)

    fmt.Println(ans)

}

// package sdq
package main

import (
    "fmt"
)



// 快慢指针。。。 不太理解， 不过： [0] 是不会被访问的。
// 这里不是找环， 这里是找 环的交点。  怎么证明 交点就是 重复值？ 。 环的交点是一个 3叉口， 可以认为 是2进1出， 那么2进 就意味者 有2个值是相同的，这样[值]才会碰到，才会形成3叉。
// [3] = 3 ,  就等另一个指针碰到 另一个3或者这个3了。  无法证明啊， 为什么另一个不会因为 [2]=2 也死循环了。 快慢指针的问题， 只可能有一个死循环， 意味 一个死循环 就足够把 2个指针全坑进去了。

// https://leetcode.com/problems/find-the-duplicate-number/discuss/72846/My-easy-understood-solution-with-O(n)-time-and-O(1)-space-without-modifying-the-array.-With-clear-explanation.



// 还有个问题， 不修改数组，固定空间的 非O(n) 解法是什么。。。   
// 二分。。 二分的是 抽象的 1-n，  如果< n/2 的 个数大于 n/2 那么 重复就在 1-n/2中。。
// 找数字 就能二分。


// 鸽巢原理。。。 follow up 的 第一问



// gg


// tle 。。。
// forfor 硬算，满足 不修改nums，使用固定空间。。。  follow up...
// 数字是[1,n], 下标是[0,n]， 直接把数组当作下标。  [x] += 1000000   >50000   %1000000
// follow up 的第一个是必然的， 数字一共n个， 位置有n+1个， 必然有重复的。
// 不修改数组，固定空间 的 O(n) 不知道。 
func findDuplicate(nums []int) int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                return nums[i]
            }
        }
    }
    return -123
}



// .... error
func findDuplicate2(nums []int) int {
    sum2 := 0
    for _, v := range nums { sum2 += v }
    return sum2 - len(nums) * (len(nums) - 1) / 2
}


func main_LT0287_20211203() {
// func main() {

    fmt.Println("ans:")


}

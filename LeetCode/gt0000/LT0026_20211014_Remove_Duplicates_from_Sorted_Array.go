// package sdq
package main

import (
    "fmt"
)


// int lst = nums[0] - 1;
// for (int i = 0; i < nums.size(); ++i)
// {
//     int t2 = nums[i];
//     if (t2 != lst)
//         nums[idx++] = t2;
//     lst = t2;
// }
// 。。。。





// Runtime: 8 ms, faster than 86.39% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 4.5 MB, less than 46.27% of Go online submissions for Remove Duplicates from Sorted Array.
func removeDuplicates(nums []int) int {
    ans := dfsa4(nums, 0, 0)
    return ans
}

// return deepest
func dfsa4(nums []int, depth int, idx int) (int) {
    if idx >= len(nums) {
        return depth
    }
    t2 := nums[idx]
    for idx++; idx < len(nums) && nums[idx] == t2; idx++ {}
    t3 := dfsa4(nums, depth + 1, idx)
    nums[depth] = t2
    return t3
}

// un-finish
func removeDuplicates2(nums []int) int {
    lst := len(nums) - 1
    for i := 0; i <= lst; {
        // if nums[i] == nums[i - 1] {
        //     // .... 这个移动量也太大了吧。。。ok。一次性把这个值 移光。
        // }
        en := i
        for en <= lst && nums[i] == nums[en] {
            en++
        }
        t2 := en - i - 1
        if t2 != 0 {
            for j := en; j <= lst; j++ {
                // ... 感觉直接dfs 一步到位就完事了。
            }
        }
        i = en          // [en] != [i]
    }
    return lst + 1
}

func main_LT0026_20211014() {
// func main() {

    fmt.Println("ans:")

    for i := 0; i < 10; i++ {
        fmt.Println(i)
        i = i + 3
    }

}

// package sdq
package main

import (
    "fmt"
)



// dp := []int{-1e5}
// for _, v := range nums {
//     i := sort.Search(len(dp), func(i int) bool { return dp[i] >= v })
//     if i < len(dp) { dp[i] = v } else { dp = append(dp, v) }
//     // fmt.Println(i, dp, v)
// }
// // fmt.Println()
// return len(dp) - 1


// Patience Sort

// (1) if x is larger than all tails, append it, increase the size by 1
// (2) if tails[i-1] < x <= tails[i], update tails[i]


// Arrays.fill(dp, 1);
// for (int i = 0; i < nums.length; i++) {
//     for (int j = 0; j < i; j++) {
//         if (nums[i] > nums[j]) 
//             dp[i] = Math.max(dp[i], dp[j] + 1);
//     }
// }



// public int lengthOfLIS(int[] nums) {
//     List<Integer> piles = new ArrayList<>(nums.length);
//     for (int num : nums) {
//         int pile = Collections.binarySearch(piles, num);
//         if (pile < 0) pile = ~pile;
//         if (pile == piles.size()) {
//             piles.add(num);
//         } else {
//             piles.set(pile, num);
//         }
//     }
//     return piles.size();
// }







// 从头到尾， 新队列里 二分找到位置， 如果 比所有的都大，那么 seq 尾巴增加。 否则 就更新， 二分需要返回 大于等于值 的位置。
// 感觉对，但是根本无法证明啊。  
// 得出的是 LIS，并且 每个位置 都是 这个位置 的 最小的可能值。
// 应该就是 最小可能值，所以能 append。

// 无涯






// Runtime: 304 ms, faster than 5.46% of Go online submissions for Longest Increasing Subsequence.
// Memory Usage: 5 MB, less than 5.46% of Go online submissions for Longest Increasing Subsequence.
// map + search
func lengthOfLIS(nums []int) int {
    ans, sz1 := 1, len(nums)
    map2 := map[int]int{ nums[sz1 - 1] : 1 }      // >=.count
    for i := sz1 - 2; i >= 0; i-- {
        t2, cnt := nums[i], 0
        for k, v := range map2 {
            if k > t2 && cnt < v {
                cnt = v
            }
        }
        if _, ok := map2[t2]; ok {
            if map2[t2] < cnt + 1 {
                map2[t2] = cnt + 1
            }
        } else {
            map2[t2] = cnt + 1
        }
        if cnt + 1 > ans {
            ans = cnt + 1
        }
    }
    return ans
}


// error...
// // 从后往前，插入排序。 就是 O(nlogn)。 不是。 看插入排序，内层还是for。。。 应该是二分啊。不不不。即使二分，还是需要数据移动。。 所以 确实不是 nlogn... bububu, 移动也是 logn的移动啊。 所以差不多是 nlogn..不稳定，所以确实是 O(n) - O(n^2)
// // 需要一个 map  upper_bound
// // ... 重复值。。。
// func lengthOfLIS(nums []int) int {
//     ans, sz1 := 1, len(nums)
//     st := sz1 - 1
//     map2 := map[int]bool{ nums[sz1 - 1] : true}
//     for i := len(nums) - 2; i >= 0; i-- {
//         t2 := nums[i]
//         // move := true
//         if _, ok := map2[t2]; ok {
//             // move = false
//             l, r := st, sz1
//             for l < r {
//                 mid := (l + r) / 2
//                 if nums[mid] == t2 {
//                     l, r = mid, mid
//                 } else if nums[mid] > t2 {
//                     r = t2 - 1
//                 } else {
//                     l = t2 + 1
//                 }
//             }
//             if (sz1 - l + 1) > ans {
//                 ans = sz1 - l + 1
//                 fmt.Println(".... ", ans)
//             }
//         } else {
//             // nums[st - 1] = t2
//             for j := st; j < sz1; j++ {
//                 if nums[j] < t2 {
//                     nums[j - 1] = nums[j]
//                 } else {
//                     nums[j - 1] = t2
//                     if sz1 - j + 2 > ans {
//                         ans = sz1 - j + 2
//                         fmt.Println(" - - ", ans, i, j, nums)
//                     }
//                     break
//                 }
//                 // if nums[j] < nums[j - 1] {
//                 //     nums[j - 1], nums[j] = nums[j], nums[j - 1]
//                 // } else {

//                 // }
//             }
//             if nums[sz1 - 1] < t2 {
//                 nums[sz1 - 1] = t2
//             }
//             map2[t2] = true
//             st--
//         }
//     }
//     return ans

//     // ans := 1
//     // for i := len(nums) - 2; i >= 0; i-- {
//     //     for j := i + 1; j < len(num); j++ {
//     //         if nums[j] < nums[j - 1] {

//     //         }
//     //     }
//     // }
//     // return ans
// }


// func main_LT0300_20211205() {
func main() {

    fmt.Println("ans:")

    // arr := []int{10,9,2,5,3,7,101,18}
    // arr := []int{0,1,0,3,2,3}
    arr := []int{7,7,7,7,7,6,8}

    fmt.Println(lengthOfLIS(arr))

}

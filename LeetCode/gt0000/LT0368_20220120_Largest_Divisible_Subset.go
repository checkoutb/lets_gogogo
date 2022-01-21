// package sdq
package main

import (
    "fmt"
    "sort"
)



// vector<int>dp(nums.size(),1);
// sort(nums.begin(),nums.end());
// for(int i=1;i<nums.size();i++){
//     for(int j=0;j<i;j++){
//         if(nums[i]%nums[j]==0) dp[i]=max(dp[i],1+dp[j]);
//     }
// }
// int val=0;
// for(int i=0;i<dp.size();i++) if(val<dp[i]) val=dp[i];
// vector<int>v;
// for(int i=dp.size()-1;i>-1;i--){
//     if(val==dp[i] and (!v.size() or v[v.size()-1]%nums[i]==0)){
//         v.push_back(nums[i]);
//         val--;
//     }
// }



// for (int i = 0; i < nums.size(); ++i) {
//     for (int j = i - 1; j >= 0; --j) {
//         if (nums[i] % nums[j] == 0 && count[i] < count[j] + 1) {
//             count[i] = count[j] + 1;
//             parent[i] = j;
//             if (count[i] > max_count) {
//                 max_count = count[i];
//                 idx = i;
//             }
//         }
//     }
// }


// S = {-1: set()}
// for x in sorted(nums):
//     S[x] = max((S[d] for d in S if x % d == 0), key=len) | {x}
// return list(max(S.values(), key=len))



// 之前的cpp  hash  tle了。
// Runtime: 72 ms, faster than 16.67% of Go online submissions for Largest Divisible Subset.
// Memory Usage: 3.2 MB, less than 42.86% of Go online submissions for Largest Divisible Subset.
// 最小公倍数。 a*b/gcd(a,b)   22可约，必然是倍数。
// sort后， 从小的 开始 *2 *3 找到下一个 然后 下一个再 *2 *3 ... .. bu, 遍历时判断 % == 0
// 必须是 已有子集 中最大数 的 倍数
// 越界，但int64 不管了
// hash 应该更快 ... hash 需要 主动*2 *3 ..
func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    // ans := 1
    ans := []int{}
    ansi := 0
    arr := make([]int, len(nums))
    nxt := make([]int, len(nums))

    // for i, _ := range arr {
    //     arr[i] = 1
    // }

    for i := 0; i < len(nums); i++ {
        t2 := dfsa368(nums, arr, nxt, i)
        if t2 > arr[ansi] {
            ansi = i
        }
    }
    ans = append(ans, nums[ansi])
    for nxt[ansi] != 0 {
        ansi = nxt[ansi]
        ans = append(ans, nums[ansi])
    }
    return ans

    // for i := 0; i < len(nums); i++ {
    //     if arr[i] != 0 { continue }
    //     cnt := 1
    //     t2 := nums[i]
    //     for j := i + 1; j < len(nums); j++ {
    //         if nums[i] % t2 == 0 {
    //             if arr[i] != 0 {
    //                 t2 += arr[i]
    //                 break
    //             }
    //             // 无法回写到 arr[i]         /// ... 如果从后往前的话， 就可以了。
    //         }
    //     }
    // }

    // map2 := map[int]int{}
    // for i, v := range nums {
    //     map2[v] = i
    // }
    // for i := 0; i < len(nums); i++ {
    //     if nums[i] == -1 {
    //         continue
    //     }
    // }

    // for i := 0; i < len(nums); i++ {
    //     if nums[i] == -1 {
    //         continue
    //     }
    //     t2 := 1
    //     for j := i + 1; j < len(nums);
    // }
    // return ans
}

func dfsa368(nums []int, arr []int, nxt []int, idx int) int {
    // if idx >= len(nums) {
    //     return 0
    // }
    if (arr)[idx] != 0 {
        return (arr)[idx]
    }
    t2 := nums[idx]
    for i := idx + 1; i < len(nums); i++ {
        if nums[i] % t2 == 0 {
            cnt := dfsa368(nums, arr, nxt, i)
            if cnt + 1 > arr[idx] {
                arr[idx] = cnt + 1
                nxt[idx] = i
            }
        }
    }
    if arr[idx] == 0 {
        arr[idx] = 1
    }
    return arr[idx]
}

// func dfsa368(arr []int,)


func main_LT0368_20220120() {
// func main() {

    fmt.Println("ans:")


}

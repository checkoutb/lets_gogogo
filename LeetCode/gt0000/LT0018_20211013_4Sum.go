// package sdq
package main

import (
    "fmt"
    "sort"
)


// 左右逼近。。。初始值。
// int l = pos, r = size-1;
// while(l < r)
// {
//     int t = nums[l]+nums[r];
//     if(t > target) {do{r--;}while(l<r and nums[r]==nums[r+1]);}
//     else if(t < target) {do{l++;}while(l<r and nums[l]==nums[l-1]);}
//     else
//     {
//         v.push_back(nums[l++]);
//         v.push_back(nums[r--]);
//         vv.push_back(v);
//         v.pop_back(), v.pop_back();
//         while(l<r && nums[l]==nums[l-1]) l++;
//         while(l<r && nums[r]==nums[r+1]) r--;
//     }
// }

// 跳过相同值。
// for(int i = start;i<nums.length;i++){
//     if(i>start && nums[i-1]==nums[i]) continue;


// for i := 0; i < l - 3; i++ {
//     if i > 0 && nums[i] == nums[i-1]{
//         continue
//     }
//     for j := i + 1; j < l - 2; j++ {
//         if j > i + 1 && nums[j] == nums[j-1] {
//             continue
//         }
//         lo := j + 1
//         hi := l - 1
//         for lo < hi {
//             sum := nums[i] + nums[j] + nums[lo] + nums[hi]
//             if(target > sum) {
//                 lo++
//             } else if (target < sum) {
//                 hi--
//             } else {
//                 ans = append(ans, []int{nums[i],nums[j],nums[lo],nums[hi]})
//                 lo++
//                 hi--
//                 for nums[lo] == nums[lo - 1] && lo < hi {
//                     lo++
//                 }
//                 for nums[hi] == nums[hi + 1] && lo < hi {
//                     hi--
//                 }
//             }
//         }
//     }
// }



// sz1 < 200
// 10^9 会溢出。   bu, int == int64 。。
// [][]int 能不能作为一个map的key？ 试下吧。。  no, invalid.

// quchong...数字重复。保存下标。只有下标小于i的，才使用。

type Tuple struct {
    a int
    b int
    c int
    d int
}


// Runtime: 1584 ms, faster than 5.33% of Go online submissions for 4Sum.
// Memory Usage: 8.6 MB, less than 6.56% of Go online submissions for 4Sum.
func fourSum(nums []int, target int) [][]int {

    ans := [][]int{}
    sort.Ints(nums)
    // sz1 := len(nums)
    map2 := make(map[int][]Tuple)       // index
    map3 := make(map[Tuple]int)         // ans     value
    for i, v := range nums {
        // for j := i + 1; j < sz1; j++ {
        for j := 0; j < i; j++ {                // j < i
            sum2 := v + nums[j]     
            tar := target - sum2
            tupArr, ok := map2[tar]
            if ok {
                for _, tup := range tupArr {
                    if tup.b < j {
                        tup2 := Tuple{nums[tup.a], nums[tup.b], nums[j], v}
                        map3[tup2] = 1
                    }
                }
            }
            tup3 := Tuple{j, i, -1, -1}
            map2[sum2] = append(map2[sum2], tup3)
        }
    }
    for tup, _ := range map3 {
        ans = append(ans, []int{tup.a, tup.b, tup.c, tup.d})
    }
    return ans

    // ans := [][]int{}
    // sz1 := len(nums)
    // sort.Ints(nums)
    // map2 := make(map[int]map[[]int]int)
    // map3 := make(map[[]int]int)         // ans
    // for i, v := range nums {
    //     for j := i + 1; j < sz1; j++ {
    //         sum2 := v + nums[j]
    //         tar := target - sum2
    //         mapa, ok := map2[tar]
    //         if ok {
    //             for arrk, _ := range mapa {
    //                 map3[[]int{arrk[0], arrk[1], v, nums[j]}] = 1
    //             }
    //         }
    //     }
    // }
    // return ans

    // sz1 := len(nums)
    // ans := [][]int{}
    // sort.Ints(nums)
    // // map3 := map[int]map[[]int]int{}
    // map2 := make(map[int][][]int)
    // for idx, val := range nums {
    //     for j := idx + 1; j < sz1; j++ {
    //         sum2 := val + nums[j]
    //         tar := target - sum2
    //         arr2, ok := map2[tar]
    //         if ok {
    //             for _, arr3 := range arr2 {
    //                 ans = append(ans, []int{arr3[0], arr3[1], val, tar})
    //             }
    //         }
    //     }
    // }
    // return ans
}



func main_LT0018_20211013() {
// func main() {

    fmt.Println("ans:")

    // var arr []int = []int{1,0,-1,0,-2,2}
    arr := []int{2,2,2,2,2,2,2,2}
    tar := 8

    var ans [][]int = fourSum(arr, tar)

    for _, arr2 := range ans {
        for _, val := range arr2 {
            fmt.Println(val)
        }
        fmt.Println()
    }

}

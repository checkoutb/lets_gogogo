// package sdq
package main

import (
    "fmt"
    "sort"
)



// var (
// 	fact = [...]int{2, 3, 5}
// 	memo = &cache{0}
// 	edge = &IntHeap{1}
// )


// class Solution:
//     ugly = sorted(2**a * 3**b * 5**c
//                   for a in range(32) for b in range(20) for c in range(14))
//     def nthUglyNumber(self, n):
//         return self.ugly[n-1]


// def nthUglyNumber(self, n):
//     ugly = [1]
//     i2 = i3 = i5 = 0
//     while len(ugly) < n:
//         while ugly[i2] * 2 <= ugly[-1]: i2 += 1
//         while ugly[i3] * 3 <= ugly[-1]: i3 += 1
//         while ugly[i5] * 5 <= ugly[-1]: i5 += 1
//         ugly.append(min(ugly[i2] * 2, ugly[i3] * 3, ugly[i5] * 5))
//     return ugly[-1]


// vector <int> results (1,1);
// int i = 0, j = 0, k = 0;
// while (results.size() < n)
// {
//     results.push_back(min(results[i] * 2, min(results[j] * 3, results[k] * 5)));
//     if (results.back() == results[i] * 2) ++i;
//     if (results.back() == results[j] * 3) ++j;
//     if (results.back() == results[k] * 5) ++k;
// }
// return results.back();


// (1) 1×2, 2×2, 3×2, 4×2, 5×2, …
// (2) 1×3, 2×3, 3×3, 4×3, 5×3, …
// (3) 1×5, 2×5, 3×5, 4×5, 5×5, …



// Runtime: 648 ms, faster than 5.21% of Go online submissions for Ugly Number II.
// Memory Usage: 7.3 MB, less than 9.37% of Go online submissions for Ugly Number II.
// 数组筛选？ 但是第1690个 是多少？ stream倒是可以。
// 第144个是5000 .....
// 265 是50000。。。。

// 需要一个pri_queue啊。  把 top*[2,3,5] 再放进去。
func nthUglyNumber(n int) int {
    arr := []int{1}
    map2 := map[int]bool{}
    for n > 1 {
        for _, num := range [3]int{2,3,5} {
            if _, ok := map2[arr[0] * num]; !ok {
                arr = append(arr, arr[0] * num)
                map2[arr[0] * num] = true
            }
        }
        arr = arr[1 : ]
        sort.Ints(arr)
        // fmt.Println(arr)
        n--
    }
    // sort.Ints(arr)
    return arr[0]


    // arr := [5000]bool{}
    // arr[1] = true
    // i := 1
    // for ; i < 5000 && n > 0; i++ {
    //     if arr[i] {
    //         n--
    //         for _, num := range [3]int{2,3,5} {
    //             if i * num < 5000 { arr[i * num] = true }
    //         }
    //         // if i * 2 < 5000 { arr[i * 2] = true }
    //         // if i * 3 < 5000 { arr[i * 3] = true }
    //         // if i * 5 
    //     }
    // }
    // return i

    // arr[1], arr[2], arr[3], arr[5] = true, true, true, true
    // // 每个true 乘一个 2，3，5
    // for i := 4
}


// func main_LT0264_20211129() {
func main() {

    fmt.Println("ans:")

    n := 265

    fmt.Println(nthUglyNumber(n))

}

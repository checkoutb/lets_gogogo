// package sdq
package main

import (
    "fmt"
)

// D D

// int res = -1, max_num[10] = {};
// for (auto num : nums) {
//     int max_d = 0;
//     for (int v = num; v; v /= 10)
//         max_d = max(max_d, v % 10);
//     if (max_num[max_d])
//         res = max(res, max_num[max_d] + num);
//     max_num[max_d] = max(max_num[max_d], num);
// }
// one array is enough.


// Runtime14 ms
// Beats
// 70.81%
// Memory5.1 MB
// Beats
// 94.74%
func maxSum(nums []int) int {
    var vi []int = make([]int, 10);
    var v2 []int = make([]int, 10);
    var t2, t3 = 0, 0;
    for _, val := range nums {
        t2 = 0;
        t3 = val;
        for val > 0 {
            if val % 10 > t2 {
                t2 = val % 10;
            }
            val /= 10;
        }
        if t3 >= vi[t2] {
            v2[t2] = vi[t2];
            vi[t2] = t3;
        } else if t3 >= v2[t2] {
            v2[t2] = t3;
        }
    }
    var ans = -1;
    for i := 1; i < 10; i++ {
        if v2[i] != 0 {
            if v2[i] + vi[i] > ans {
                ans = v2[i] + vi[i];
            }
        }
    }
    return ans;
}


func main_LT2815_20231010() {
// func main() {

    fmt.Println("ans:")


}

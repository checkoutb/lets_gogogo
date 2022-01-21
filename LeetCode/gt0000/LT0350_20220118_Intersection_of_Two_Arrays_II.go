// package sdq
package main

import (
    "fmt"
)

// for (int i : nums1)
//     map2[i]++;
// for (int i : nums2)
// {
//     if (map2[i])        // > 0
//     {
//         map2[i]--;
//         result.push_back(i);
//     }
// }


// Runtime: 7 ms, faster than 19.63% of Go online submissions for Intersection of Two Arrays II.
// Memory Usage: 3.3 MB, less than 27.10% of Go online submissions for Intersection of Two Arrays II.
func intersect(nums1 []int, nums2 []int) []int {
    map2 := map[int]int{}
    for _, v := range nums1 {
        map2[v]++
    }
    for _, v := range nums2 {
        if _, ok := map2[v]; ok {
            map2[v] += 10000
        }
    }
    ans := []int{}
    for k, v := range map2 {
        a, b := v / 10000, v % 10000
        t2 := a
        if a > b {
            t2 = b
        }
        for i := 0; i < t2; i++ {
            ans = append(ans, k)
        }
    }
    return ans
}

func main_LT0350_20220118() {
// func main() {

    fmt.Println("ans:")


}

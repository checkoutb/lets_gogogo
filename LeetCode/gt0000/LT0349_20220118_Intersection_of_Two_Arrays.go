// package sdq
package main

import (
    "fmt"
)


// for _, val := range nums2 {
//     if potentialMatches[val] {
//         result = append(result, val)
//         potentialMatches[val] = false
//     }
// }

// set<int> s(nums1.begin(), nums1.end());
// vector<int> out;
// for (int x : nums2)
//     if (s.erase(x))
//         out.push_back(x);


// Runtime: 2 ms, faster than 77.82% of Go online submissions for Intersection of Two Arrays.
// Memory Usage: 3.4 MB, less than 7.75% of Go online submissions for Intersection of Two Arrays.
func intersection(nums1 []int, nums2 []int) []int {
    map2 := map[int]int{}
    for _, v := range nums1 {
        map2[v] = 1
    }
    for _, v := range nums2 {
        map2[v] = map2[v] | 2
    }
    ans := []int{}
    for k, v := range map2 {
        if v > 2 {
            ans = append(ans, k)
        }
    }
    return ans
}


func main_LT0349_20220118() {
// func main() {

    fmt.Println("ans:")


}

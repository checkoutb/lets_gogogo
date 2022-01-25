// package sdq
package main

import (
    "fmt"
    "math/rand"
)



// int pick(int target) {
//     int count = 0, pickedIdx=-1;
//     for(int i=0;i<nums.size();i++) {
//         if(nums[i] == target) {
//             count++;
//             if(rand()%count == 0)
//                 pickedIdx = i;
//         }
//     }
//     return pickedIdx;
// }
// 应该是 蓄水池采样。
// 但是 这种竟然 比 map 快。。。


// r := rand.Intn(count+1)
// if r == 0 {
//     index = i
// }



// Runtime: 260 ms, faster than 10.31% of Go online submissions for Random Pick Index.
// Memory Usage: 10.2 MB, less than 29.90% of Go online submissions for Random Pick Index.
type Solution struct {
    map2 map[int][]int
}

func Constructor(nums []int) Solution {
    map2 := map[int][]int{}
    for i, v := range nums {
        map2[v] = append(map2[v], i)
    }
    return Solution{ map2 }
}

func (this *Solution) Pick(target int) int {
    return this.map2[target][rand.Intn(len(this.map2[target]))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

func main_LT0398_20220125() {
// func main() {

    fmt.Println("ans:")


}

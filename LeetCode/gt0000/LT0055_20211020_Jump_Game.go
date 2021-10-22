// package sdq
package main

import (
    "fmt"
)





// int leftMost = nums.length-1;
// for(int i = nums.length-2; i>=0; i--){
//     if(i+nums[i]>=leftMost){
//         leftMost = i;
//     }
// }
// return leftMost==0;
// 倒退走，如果 前面 有个点+长度 能到达leftMost， 就 修改leftMost。


// int currMax=nums[0];
// int ptr=0;
// while (ptr < nums.length&&ptr<=currMax){
//     currMax=Math.max(currMax,ptr+nums[ptr]);
//     ptr++;
// }
// return ptr>=nums.length;


// int limit = 0, i;
// for (i = 0; i <= limit && limit < nums.length - 1; i++) {
//     if (limit < i + nums[i]) {
//         limit = i + nums[i];
//     }
// }
// return limit >= nums.length - 1;


// for i := 0; i <= max; i++ {
//     if i+nums[i] > max {
//         max = i + nums[i]
//     }
//     if max >= len(nums)-1 {
//         return true
//     }
// }


// Runtime: 78 ms, faster than 53.60% of Go online submissions for Jump Game.
// Memory Usage: 7.9 MB, less than 27.62% of Go online submissions for Jump Game.
func canJump(nums []int) bool {
    i, idx, sz1 := 0, nums[0] + 0, len(nums)
    for idx < sz1 - 1 {
        idx2 := idx
        for i <= idx {
            t2 := nums[i] + i
            if t2 > idx2 {
                idx2 = t2
            }
            i++
        }
        if idx2 == idx {
            return false
        }
        idx = idx2
    }
    return true
}


func main_LT0055_20211020() {
// func main() {

    fmt.Println("ans:")


}

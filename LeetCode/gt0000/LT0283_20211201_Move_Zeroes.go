// package sdq
package main

import (
    "fmt"
)

// arr[insPos++] = arr[i];


// var zerosToJumpOver int
// for i, v := range nums {
//     if v == 0 {
//         zerosToJumpOver++
//     } else if zerosToJumpOver > 0 {
//         nums[i - zerosToJumpOver], nums[i] = nums[i], 0
//     }
// }


// for (int i=0;i<nums.length;i++){
//     if (nums[i]==0){
//         snowBallSize++; 
//     }
//     else if (snowBallSize > 0) {
//         int t = nums[i];
//         nums[i]=0;
//         nums[i-snowBallSize]=t;
//     }
// }
// 把 0000视为一个 雪球， 会越滚越大， 但每次最多只移动一位。 把 第一个0 和 后续的 非0 swap， 或者 雪球+后续的0 合并。


// fill(remove(nums.begin(), nums.end(),0), nums.end(), 0);
// std::stable_partition(nums.begin(), nums.end(), [](int n){return n!=0;});


// Runtime: 24 ms, faster than 72.50% of Go online submissions for Move Zeroes.
// Memory Usage: 6.8 MB, less than 92.38% of Go online submissions for Move Zeroes.
func moveZeroes(nums []int) {
    idx := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[idx], idx = nums[i], idx + 1
        }
    }
    for idx < len(nums) {
        nums[idx], idx = 0, idx + 1
    }
}


// error， need maintain the relative order
func moveZeroes2(nums []int)  {
    lst := len(nums)
    for i := 0; i < lst; i++ {
        if nums[i] == 0 {
            nums[i], i, nums[lst - 1], lst = nums[lst - 1], i - 1, nums[i], lst - 1
        }
    }
}


func main_LT0283_20211201() {
// func main() {

    fmt.Println("ans:")


}

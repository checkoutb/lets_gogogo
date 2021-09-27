
package main

import (
    "fmt"
)

// D D

// int rsum = accumulate(begin(nums), end(nums), 0);
// for (int i = 0, lsum = 0; i < nums.size(); lsum += nums[i++])
//     if (lsum * 2 == rsum - nums[i])
//         return i;
// return -1;

// int leftSum = 0;
// Map<Integer, Integer> map = new HashMap<>();
// for (int i=0; i<nums.length; leftSum+=nums[i++])
//     map.putIfAbsent(leftSum*2+nums[i], i);
// return map.getOrDefault(leftSum, -1);


// left := 0
// right := 0
// for _, v := range nums {
//     right += v
// }
// for i,v := range nums {
//     right -= v
//     if left == right {
//         return i
//     }
//     left += v
// }   
// return -1


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Find the Middle Index in Array.
// Memory Usage: 2.5 MB, less than 57.02% of Go online submissions for Find the Middle Index in Array.
func findMiddleIndex(nums []int) int {
    sz1 := len(nums)
    arr := make([]int, sz1, sz1)
    ans := -1
    // arr[sz1 - 1] = nums[sz1 - 1]
    for i := sz1 - 2; i >= 0; i-- {
        arr[i] = nums[i + 1] + arr[i + 1]
    }
    sum2 := 0
    for i := 0; i < sz1; i++ {
        if arr[i] == sum2 {
            ans = i
            break
        }
        sum2 += nums[i]
    }
    return ans;
}


func main_LT1991_20210927() {
// func main() {

    // arr := []int{2,3,-1,8,4}
    // arr := []int{1,-1,4}
    // arr := []int{2,5}
    arr := []int{1}

    fmt.Println("ans:")

    fmt.Println(findMiddleIndex(arr))
}

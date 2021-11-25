// package sdq
package main

import (
    "fmt"
)


// reverse(nums, 0, nums.length - k - 1);
// reverse(nums, nums.length - k, nums.length - 1);
// reverse(nums, 0, nums.length - 1);
// nums = [1,2,3,4,5,6,7] and k = 3, first we reverse [1,2,3,4], it becomes[4,3,2,1]; 
// then we reverse[5,6,7], it becomes[7,6,5], 
// finally we reverse the array as a whole, it becomes[4,3,2,1,7,6,5] ---> [5,6,7,1,2,3,4].


// void rotate(int nums[], int n, int k) {
//     for (; k %= n; n -= k)
//         for (int i = 0; i < k; i++)
//             swap(*nums++, nums[n - k]);
// }



// Runtime: 28 ms, faster than 97.11% of Go online submissions for Rotate Array.
// Memory Usage: 8.1 MB, less than 64.14% of Go online submissions for Rotate Array.
func rotate(nums []int, k int)  {
    sz1 := len(nums)
    k = k % sz1
    if k == 0 {
        return
    }
    cnt := gcda189(sz1, k)
    for i := 0; i < cnt; i++ {
        t2 := nums[i]
        for j := (i + k) % sz1; j > i; j = (j + k) % sz1 {
            t2, nums[j] = nums[j], t2
        }
        nums[i] = t2
        // fmt.Println(nums, ",,,,, ", k)
        // if sz1 % k != 0 {
        //     break
        // }
    }
}

func gcda189(a, b int) int {
    c := a % b
    if c == 0 {
        return b
    }
    return gcda189(b, c)
}


func main_LT0189_20211125() {
// func main() {

    fmt.Println("ans:")

    // k := 3
    // arr := []int{1,2,3,4,5,6,7}

    k := 4
    arr := []int{1,2,3,4,5,6}
    

    rotate(arr, k)

    fmt.Println(arr)

}

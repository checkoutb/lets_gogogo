// package sdq
package main

import (
    "fmt"
)

// cnt := 1
// cur := 0
// for i := 1; i < length; i++ {
//     if nums[cur] != nums[i] {
//         cnt = 1
//         cur++
//         nums[cur] = nums[i]
//     } else if cnt < 2 {
//         cnt++
//         cur++
//         nums[cur] = nums[i]
//     }        
// }
// return cur + 1


// while(i<nums.length&&j<nums.length){
//     if(nums[i-1]==nums[i]){
//         counter++;
//     }else{
//         counter=1;
//     }
//     if(counter<=2){
//         nums[j]=nums[i];
//         j++;
//     }
//     i++;
// }


// int i = 0;
// for (int n : nums)
//     if (i < 2 || n > nums[i-2])
//         nums[i++] = n;
// return i;
// 1 1 1 2 2 3
// 1 1 2 2 2 3
// ... 这里使用  重构后的 i-2 , 而不是 直接 i-2..  .. 这里的i 是 新数组的 下标.     我的go的i 是 旧数组的下标..



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted Array II.
// Memory Usage: 3.2 MB, less than 55.08% of Go online submissions for Remove Duplicates from Sorted Array II.
func removeDuplicates80(nums []int) int {
    if len(nums) <= 2 {
        return len(nums)
    }
    idx := 2
    pre, ppre := nums[1], nums[0]
    for i := 2; i < len(nums); i++ {
        if nums[i] == ppre {
            ;
        } else {
            nums[idx] = nums[i]
            idx++
        }
        ppre, pre = pre, nums[i]
    }
    return idx
}

func main_LT0080_20211025() {
// func main() {

    fmt.Println("ans:")

    arr := []int{1,1,1,2,2,3}

    ans := removeDuplicates80(arr)

    fmt.Println(ans, ", ", arr)

}

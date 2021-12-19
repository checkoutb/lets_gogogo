// package sdq
package main

import (
    "fmt"
    "sort"
)


// it's possible to find the median within O(n)-time and O(1)-space.
// We can use nth_element to find the median, but it's not O(n)-time and O(1)-space. 

// (1) Elements that are larger than the median: we can put them in the first few odd slots;
// (2) Elements that are smaller than the median: we can put them in the last few even slots;
// (3) Elements that equal the median: we can put them in the remaining slots.

// int map_index(int idx, int n) {
//     return (2 * idx + 1) % (n | 1);
// }


// Runtime: 30 ms, faster than 30.43% of Go online submissions for Wiggle Sort II.
// Memory Usage: 6.9 MB, less than 8.70% of Go online submissions for Wiggle Sort II.
// 
func wiggleSort(nums []int)  {
    arr := make([]int, len(nums))
    copy(arr, nums)
    sort.Ints(arr)
    // i1, i2, ia := 0, (len(nums) + 1) / 2, 0
    i1, i2, ia := (len(nums) - 1) / 2, len(nums) - 1, 0
    for ia < len(arr) {
        nums[ia] = arr[i1]
        ia++
        i1--
        if ia < len(arr) {
            nums[ia] = arr[i2]
            ia++
            i2--
        }
    }
}


// error
// sort 对分，merge
func wiggleSort2(nums []int)  {
    arr := make([]int, len(nums))
    copy(arr, nums)
    sort.Ints(arr)
    i1, i2, ia := 0, (len(nums) + 1) / 2, 0
    for ia < len(arr) {
        nums[ia] = arr[i1]
        ia, i1 = ia + 1, i1 + 1
        if ia < len(arr) {
            nums[ia] = arr[i2]
            ia, i2 = ia + 1, i2 + 1
        }
    }
}


func main_LT0324_20211219() {
// func main() {

    fmt.Println("ans:")

    arr := []int{4,5,5,6}  // 5 6 4 5
    // 1 1 1 2 3 
    // 1 2 2 2 3  wu jie.

}

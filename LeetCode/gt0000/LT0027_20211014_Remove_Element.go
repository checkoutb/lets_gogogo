// package sdq
package main

import (
    "fmt"
)


// k := -1
// for i := 0; i < len(nums); i++ {
//     if nums[i] != val {
//         k++
//         nums[k] = nums[i]
//     }
// }
// return k + 1


// nums.erase(remove(nums.begin(), nums.end(), val), nums.end());
// return nums.size();

// "It doesn't matter what you leave beyond the new length." So there is no need to call nums.erase.
// const auto it = std::remove(nums.begin(), nums.end(), val);
// return std::distance(nums.begin(), it);



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
// Memory Usage: 2.2 MB, less than 99.75% of Go online submissions for Remove Element.
func removeElement(nums []int, val int) int {
    lst := len(nums) - 1
    for i := 0; i <= lst; i++ {
        if nums[i] == val {
            nums[i] = nums[lst]
            lst--
            i--
        }
    }
    return lst + 1
}


func main_LT0027_20211014() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
    "fmt"
)




// Runtime: 28 ms, faster than 91.00% of Go online submissions for Range Sum Query - Immutable.
// Memory Usage: 9.6 MB, less than 62.00% of Go online submissions for Range Sum Query - Immutable.
type NumArray struct {
    Arr []int
}
func Constructor(nums []int) NumArray {
    sz1 := len(nums)
    arr := make([]int, sz1)
    arr[0] = nums[0]
    for i := 1; i < sz1; i++ {
        arr[i] = arr[i - 1] + nums[i]
    }
    return NumArray{ Arr : arr }
}
func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
        return this.Arr[right]
    } else {
        return this.Arr[right] - this.Arr[left - 1]
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */


func main_LT0303_20211206() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
    "fmt"
)

// Fenwick Tree

// Segment Tree

// Binary Index Tree




// Runtime: 3052 ms, faster than 13.89% of Go online submissions for Range Sum Query - Mutable.
// Memory Usage: 21 MB, less than 66.67% of Go online submissions for Range Sum Query - Mutable.
// 更新多，还是搜索多？
type NumArray struct {
    Arr []int
}
func Constructor(nums []int) NumArray {
    arr := make([]int, len(nums))
    arr[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        arr[i] = arr[i - 1] + nums[i]
    }
    return NumArray{ Arr : arr }
}
func (this *NumArray) Update(index int, val int)  {
    t2 := this.Arr[index]
    if index > 0 { t2 -= this.Arr[index - 1] }
    t2 = val - t2
    for i := index; i < len(this.Arr); i++ {
        this.Arr[i] += t2
    }
}
func (this *NumArray) SumRange(left int, right int) int {
    t2 := this.Arr[right]
    if left > 0 {
        t2 -= this.Arr[left - 1]
    }
    return t2
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func main_LT0307_20211206() {
// func main() {

    fmt.Println("ans:")


}

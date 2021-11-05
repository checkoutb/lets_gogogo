// package sdq
package main

import (
    "fmt"
)








// Runtime: 0 ms, faster than 100.00% of Go online submissions for Convert Sorted Array to Binary Search Tree.
// Memory Usage: 3.5 MB, less than 100.00% of Go online submissions for Convert Sorted Array to Binary Search Tree.
func sortedArrayToBST(nums []int) *TreeNode {
    return dfsa108(nums)
}

func dfsa108(arr []int) *TreeNode {
    if len(arr) == 0 {
        return nil
    }
    t2 := len(arr) / 2
    return &TreeNode{arr[t2], dfsa108(arr[0 : t2]), dfsa108(arr[t2 + 1 : ])}
}


func main_LT0108_20211104() {
// func main() {

    fmt.Println("ans:")


}

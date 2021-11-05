// package sdq
package main

import (
    "fmt"
)






// Runtime: 4 ms, faster than 90.56% of Go online submissions for Path Sum II.
// Memory Usage: 4.5 MB, less than 87.55% of Go online submissions for Path Sum II.
func pathSum(root *TreeNode, targetSum int) [][]int {
    if root == nil {
        return [][]int{}
    }
    ans := [][]int{}
    dfsa113(root, targetSum, []int{}, &ans)
    return ans
}

func dfsa113(node *TreeNode, remain int, arr []int, ans *[][]int) {
    if node.Left == nil && node.Right == nil {
        if remain - node.Val == 0 {             // remain == node.Val
            arr = append(arr, node.Val)
            arr2 := make([]int, len(arr))
            copy(arr2, arr)
            *ans = append(*ans, arr2)
        }
        return
    }
    remain -= node.Val
    arr = append(arr, node.Val)
    if node.Left != nil {
        dfsa113(node.Left, remain, arr, ans)
    }
    if node.Right != nil {
        dfsa113(node.Right, remain, arr, ans)
    }
    arr = arr[0 : len(arr) - 1]
}


func main_LT0113_20211104() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
    "fmt"
    "math"
)







/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    ans := dfsa98(root, math.MinInt32, math.MaxInt32)
    return ans
}

func dfsa98(node *TreeNode, st, en int) bool {
    if node == nil {
        return true
    }
    t2 := node.Val
    if t2 < st || t2 > en {
        return false
    }
    // if !dfsa98(node.Left, st, t2 - 1) {
    //     return false
    // }
    return dfsa98(node.Left, st, t2 - 1) && dfsa98(node.Right, t2 + 1, en)
}


func main_LT0098_20211102() {
// func main() {

    fmt.Println("ans:")


}

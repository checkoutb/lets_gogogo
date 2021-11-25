// package sdq
package main

import (
    "fmt"
)



// dfs, 用 len(ans) 和 深度对比 来决定是否 append。 先遍历右子树





// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Right Side View.
// Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Binary Tree Right Side View.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    arr := []int{}
    if root == nil {
        return arr
    }
    que := []*TreeNode{root}
    for len(que) != 0 {
        arr = append(arr, que[len(que) - 1].Val)
        que2 := []*TreeNode{}
        for _, v := range que {
            if v.Left != nil {
                que2 = append(que2, v.Left)
            }
            if v.Right != nil {
                que2 = append(que2, v.Right)
            }
        }
        que = que2
    }
    return arr
}


func main_LT0199_20211125() {
// func main() {

    fmt.Println("ans:")


}

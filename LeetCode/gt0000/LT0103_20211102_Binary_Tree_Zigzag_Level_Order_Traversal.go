// package sdq
package main

import (
    "fmt"
)


// int index = (leftToRight) ? i : (size - 1 - i);
// row[index] = node->val;


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
// Memory Usage: 2.6 MB, less than 81.34% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
// level + reverse.
// bfs
func zigzagLevelOrder(root *TreeNode) [][]int {
    ans := [][]int{}
    if root == nil {
        return ans
    }
    que := []*TreeNode{root}
    b2 := true      // l -> r
    for len(que) != 0 {
        que2 := []*TreeNode{}
        arr := []int{}
        if b2 {
            for i := 0; i < len(que); i++ {
                arr = append(arr, que[i].Val)
                if que[i].Left != nil {
                    que2 = append(que2, que[i].Left)
                }
                if que[i].Right != nil {
                    que2 = append(que2, que[i].Right)
                }
            }
        } else {
            for i := len(que) - 1; i >= 0; i-- {
            // for i := 0; i < len(que); i++ {
                arr = append(arr, que[i].Val)
                if que[i].Right != nil {
                    que2 = append(que2, que[i].Right)
                }
                if que[i].Left != nil {
                    que2 = append(que2, que[i].Left)
                }
            }
        }
        ans = append(ans, arr)
        // if len(que2) == 0 {
        //     break
        // }
        // que = make([]*TreeNode, len(que2))
        // for i := 0; i < len(que2); i++ {
        //     que[i] = que2[len(que2) - i - 1]
        // }
        que = make([]*TreeNode, len(que2))
        if b2 {
            copy(que, que2)
        } else {
            for i := 0; i < len(que2); i++ {
                que[i] = que2[len(que2) - i - 1]
            }
        }
        b2 = !b2
    }
    return ans
}


func main_LT0103_20211102() {
// func main() {

    fmt.Println("ans:")


}

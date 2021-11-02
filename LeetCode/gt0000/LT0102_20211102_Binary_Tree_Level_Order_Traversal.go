// package sdq
package main

import (
    "fmt"
)



// if len <= depth


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
// Memory Usage: 3.1 MB, less than 45.85% of Go online submissions for Binary Tree Level Order Traversal.
func levelOrder(root *TreeNode) [][]int {
    ans := [][]int{}
    dfsa102(root, 0, &ans)
    return ans
}

func dfsa102(node *TreeNode, depth int, ans *[][]int) {
    if node == nil {
        return
    }
    for len(*ans) <= depth {                // "if" is enough
        (*ans) = append((*ans), []int{})
    }
    (*ans)[depth] = append((*ans)[depth], node.Val)
    dfsa102(node.Left, depth + 1, ans)
    dfsa102(node.Right, depth + 1, ans)
}

func main_LT0102_20211102() {
// func main() {

    fmt.Println("ans:")


}

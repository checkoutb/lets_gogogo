// package sdq
package main

import (
    "fmt"
)





// java.list.add(0, arr)




// 使用 ans = append(ans, arr)  +  reverse
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
// Memory Usage: 2.8 MB, less than 84.62% of Go online submissions for Binary Tree Level Order Traversal II.

// 使用  ans = append([][]int{arr}, ans...)
// Runtime: 4 ms, faster than 34.19% of Go online submissions for Binary Tree Level Order Traversal II.
// Memory Usage: 6.8 MB, less than 11.11% of Go online submissions for Binary Tree Level Order Traversal II.
// ans.reverse
// append([]int{x}, ans...)
func levelOrderBottom(root *TreeNode) [][]int {
    ans := [][]int{}
    if root == nil {
        return ans
    }
    que := []*TreeNode{root}

    for len(que) > 0 {
        que2 := []*TreeNode{}
        arr := []int{}
        for _, v := range que {
            arr = append(arr, v.Val)
            if v.Left != nil {
                que2 = append(que2, v.Left)
            }
            if v.Right != nil {
                que2 = append(que2, v.Right)
            }
        }
        // ans = append([][]int{arr}, ans...)          // .
        ans = append(ans, arr)
        que = que2
    }

    for i, j := 0, len(ans) - 1; i < j; i++ {           // 每个part (应该)只能有 一条语句.
        ans[i], ans[j] = ans[j], ans[i]
        j--
    }
    return ans
}


func main_LT0107_20211104() {
// func main() {

    fmt.Println("ans:")


}

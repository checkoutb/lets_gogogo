// package sdq
package main

import (
    "fmt"
)




// 如果结点不为nil 就一直压栈,并且指向 左结点,  然后出栈一个, ++ans, 结点指向右结点,    然后 从本行的开头 继续.

// Morris


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
// Memory Usage: 2.1 MB, less than 99.88% of Go online submissions for Binary Tree Inorder Traversal.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ans := []int{}
    dfsa94(root, &ans)
    return ans
}

func dfsa94(node *TreeNode, ans *[]int) {               // 必须 * , 切片是指针的,但是 指向切片的 变量 不是指针的
    if node == nil {
        return
    }
    dfsa94(node.Left, ans)
    *ans = append(*ans, node.Val)
    dfsa94(node.Right, ans)
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main_LT0094_20211102() {
// func main() {

    fmt.Println("ans:")


}

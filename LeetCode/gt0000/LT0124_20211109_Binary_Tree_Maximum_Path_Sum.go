// package sdq
package main

import (
    "fmt"
)


// int temp=Math.max(Math.max(l,r)+root.val,root.val);
// int ans=Math.max(temp,l+r+root.val);





// Runtime: 12 ms, faster than 98.62% of Go online submissions for Binary Tree Maximum Path Sum.
// Memory Usage: 8 MB, less than 61.03% of Go online submissions for Binary Tree Maximum Path Sum.
// any non-empty path.  
// 没有说 要到 叶子结点。。
func maxPathSum(root *TreeNode) int {
    _, ans := dfsa124(root)
    return ans
}

func dfsa124(node *TreeNode) (int, int) {     // path_to_leaf, max_Path....  max-half-path, max-path
    if node == nil {
        // return 0, 0
        return -111111, -111111
    }
    l1, l2 := dfsa124(node.Left)        // ..
    r1, r2 := dfsa124(node.Right)
    t2, t3 := node.Val, node.Val           // t3 必须做一次链接 ... 不  [-3]
    t2 = mx124(t2, mx124(t2 + l1, t2 + r1))
    t3 = mx124(t3 + l1 + r1, mx124(t3 + l1, mx124(t3, t3 + r1)))
    t3 = mx124(t3, mx124(l2, r2))
    return t2, t3

    // l1, l2, r1, r2 := dfsa124(node.Left), dfsa124(node.Right)
    // t2, t3 := r1 + node.Val, node.Val + l1 + r1
    // if l1 > r1 {
    //     t2 = l1 + node.Val
    // }
    // if l2 > t3 {
    //     t3 = l2
    // }
    // if r2 > t3 {
    //     t3 = r2
    // }
    // return t2, t3
}

func mx124(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func main_LT0124_20211109() {
// func main() {

    fmt.Println("ans:")


}

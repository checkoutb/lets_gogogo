// package sdq
package main

import (
    "fmt"
)




// Runtime82 ms
// Beats
// 43.69%
// Memory7.1 MB
// Beats
// 99.66%

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    
    // dfs := func(node *TreeNode, low int, high int) int {
    //     if node == nil || node.Val < low || node.Val > high{
    //         return 0;
    //     }
    //     return node.Val + dfs(node.Left, low, high) + dfs(node.Right, low, high);        // undefined: dfs
    // }

    // return dfs(root, low, high);
    return dfsa1(root, low, high);
}

func dfsa1(node *TreeNode, low int, high int) int {
    if node == nil {
        return 0
    }
    if node.Val >= low && node.Val <= high {
        return dfsa1(node.Left, low, high) + node.Val + dfsa1(node.Right, low, high);
    } else if node.Val < low {
        return dfsa1(node.Right, low, high);
    } else {
        return dfsa1(node.Left, low, high);
    }
    // if node == nil || node.Val < low || node.Val > high {
    //     return 0;
    // }
    // return node.Val + dfsa1(node.Left, low, high) + dfsa1(node.Right, low, high);
}

func main_LT0938_20240108() {
// func main() {

    fmt.Println("ans:")


}

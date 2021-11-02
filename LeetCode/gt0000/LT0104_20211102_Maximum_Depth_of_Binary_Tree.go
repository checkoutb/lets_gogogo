// package sdq
package main

import (
    "fmt"
)


// if root == nil {
//     return 0
// }
// left := 1 + maxDepth(root.Left)
// right := 1 + maxDepth(root.Right)
// return int(math.Max(float64(left), float64(right)))

// return root == NULL ? 0 : max(maxDepth(root -> left), maxDepth(root -> right)) + 1;


// Runtime: 4 ms, faster than 87.58% of Go online submissions for Maximum Depth of Binary Tree.
// Memory Usage: 4.4 MB, less than 60.22% of Go online submissions for Maximum Depth of Binary Tree.
func maxDepth(root *TreeNode) int {
    return dfsa104(root, 0)
}

func dfsa104(node *TreeNode, depth int) int {
    if node == nil {
        return depth
    }
    t1, t2 := dfsa104(node.Left, depth + 1), dfsa104(node.Right, depth + 1)
    if t1 > t2 {
        return t1
    } else {
        return t2
    }
}

func main_LT0104_20211102() {
// func main() {

    fmt.Println("ans:")


}

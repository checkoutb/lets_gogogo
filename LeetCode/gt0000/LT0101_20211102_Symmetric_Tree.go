// package sdq
package main

import (
    "fmt"
)



// deque  bfs
// que * 2  bfs



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Symmetric Tree.
// Memory Usage: 2.9 MB, less than 57.83% of Go online submissions for Symmetric Tree.
// 2棵树, 你往左 我往右.
func isSymmetric(root *TreeNode) bool {
    return dfsa101(root.Left, root.Right)
}

func dfsa101(p, q *TreeNode) bool {
    if p == nil || q == nil {
        return q == p
    }
    return p.Val == q.Val && dfsa101(p.Left, q.Right) && dfsa101(p.Right, q.Left)
}

func main_LT0101_20211102() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
    "fmt"
)





// Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
// Memory Usage: 2.1 MB, less than 54.80% of Go online submissions for Same Tree.
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return dfsa100(p, q)    
}

func dfsa100(p, q *TreeNode) bool {
    if p == nil || q == nil {
        return p == q
    }
    return p.Val == q.Val && dfsa100(p.Left, q.Left) && dfsa100(p.Right, q.Right)
}

func main_LT0100_20211102() {
// func main() {

    fmt.Println("ans:")


}

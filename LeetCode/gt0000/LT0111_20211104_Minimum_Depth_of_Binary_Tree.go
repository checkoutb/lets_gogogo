// package sdq
package main

import (
    "fmt"
)





// if (root == nil) {
//     return 0
// }
// if (root.Left == nil) {
//     return minDepth(root.Right) + 1
// } else if (root.Right == nil) {
//     return minDepth(root.Left) + 1
// } else {
//     return min(minDepth(root.Left) + 1, minDepth(root.Right) + 1)
// }


// queue   bfs


// int minDepth(TreeNode* root) {
//     if (!root) return 0;
//     int L = minDepth(root->left), R = minDepth(root->right);
//     return 1 + (!L-!R ? max(L, R) : min(L, R));
// }
// int minDepth(TreeNode* root) {
//     if (!root) return 0;
//     int L = minDepth(root->left), R = minDepth(root->right);
//     return L<R && L || !R ? 1+L : 1+R;
// }
// ! 一个int， 是 把int转为bool ？



// Runtime: 184 ms, faster than 97.04% of Go online submissions for Minimum Depth of Binary Tree.
// Memory Usage: 19.4 MB, less than 60.53% of Go online submissions for Minimum Depth of Binary Tree.
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return dfsa111(root)
}

func dfsa111(node *TreeNode) int {
    if node.Left == nil && node.Right == nil {
        return 1
    }
    t2 := 1000000
    if node.Left != nil {
        t2 = dfsa111(node.Left)
    }
    if node.Right != nil {
        t3 := dfsa111(node.Right)
        if t3 < t2 {
            t2 = t3
        }
    }
    return t2 + 1
}


func main_LT0111_20211104() {
// func main() {

    fmt.Println("ans:")


}

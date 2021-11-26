// package sdq
package main

import (
    "fmt"
)


// 先序遍历
// stack 如果有left 就压left，没有就 pop，并且走到 Right，然后继续push left


// int kthSmallest(TreeNode* root, int& k) {
//     if (root) {
//         int x = kthSmallest(root->left, k);
//         return !k ? x : !--k ? root->val : kthSmallest(root->right, k);
//     }
// }


// t1,t2 混了。
// Runtime: 6 ms, faster than 94.87% of Go online submissions for Kth Smallest Element in a BST.
// Memory Usage: 6.6 MB, less than 74.68% of Go online submissions for Kth Smallest Element in a BST.
func kthSmallest(root *TreeNode, k int) int {
    _, t2 := dfsa230(root, k)
    return t2
}

// 子树结点数， k个的值
func dfsa230(node *TreeNode, k int) (int, int) {
    if node == nil {
        return 0, -1
    }
    t1, t2 := dfsa230(node.Left, k)
    if t2 != -1 {
        return 0, t2
    }
    if t1 + 1 == k {
        return 0, node.Val
    }
    t3, t4 := dfsa230(node.Right, k - t1 - 1)
    return t1 + t3 + 1, t4
}

func main_LT0230_20211126() {
// func main() {

    fmt.Println("ans:")


}

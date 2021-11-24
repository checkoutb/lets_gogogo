// package sdq
package main

import (
    "fmt"
)



// if root == nil{
//     return nil
// }
// var result []int
// result = append(result, postorderTraversal(root.Left)...)
// result = append(result, postorderTraversal(root.Right)...)
// result = append(result, root.Val)
// return result


// while (root || !todo.empty()) {
//     if (root) {
//         todo.push(root);
//         root = root -> left;
//     } else {
//         TreeNode* node = todo.top();
//         if (node -> right && last != node -> right) {
//             root = node -> right;
//         } else {
//             nodes.push_back(node -> val);
//             last = node;
//             todo.pop();
//         }
//     }
// }



// Runtime: 3 ms, faster than 11.14% of Go online submissions for Binary Tree Postorder Traversal.
// Memory Usage: 2.2 MB, less than 27.72% of Go online submissions for Binary Tree Postorder Traversal.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    ans := []int{}
    dfsa145(root, &ans)
    return ans
}

func dfsa145(node *TreeNode, ans *[]int) {
    if node == nil {
        return
    }
    dfsa145(node.Left, ans)
    dfsa145(node.Right, ans)
    *ans = append(*ans, node.Val)
}


func main_LT0145_20211122() {
// func main() {

    fmt.Println("ans:")


}

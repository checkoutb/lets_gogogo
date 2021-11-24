// package sdq
package main

import (
    "fmt"
)


// var res[]int
// var stack = []*TreeNode{root}
// for len(stack) != 0 {
//     cur := stack[len(stack)-1]
//     stack = stack[:len(stack)-1]
//     if cur != nil {
//         res = append(res, cur.Val)
//         stack = append(stack, cur.Right, cur.Left) // note how right node is pushed into stack before left node because we want to pop left node first and process
//     }
// }
// return res


// Stack<TreeNode> rights = new Stack<TreeNode>();
// while(node != null) {
//     list.add(node.val);
//     if (node.right != null) {
//         rights.push(node.right);
//     }
//     node = node.left;
//     if (node == null && !rights.isEmpty()) {
//         node = rights.pop();
//     }
// }


// while (root) {
//     if (root -> left) {
//         TreeNode* pre = root -> left;
//         while (pre -> right && pre -> right != root) {
//             pre = pre -> right;
//         }
//         if (!pre -> right) {
//             pre -> right = root;
//             nodes.push_back(root -> val);
//             root = root -> left;
//         } else {
//             pre -> right = NULL;
//             root = root -> right;
//         }
//     } else {
//         nodes.push_back(root -> val);
//         root = root -> right;
//     }
// }
// Morris


// Runtime: 4 ms, faster than 5.87% of Go online submissions for Binary Tree Preorder Traversal.
// Memory Usage: 2.2 MB, less than 24.40% of Go online submissions for Binary Tree Preorder Traversal.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    ans := []int{}
    dfsa144(root, &ans)
    return ans
}

func dfsa144(node *TreeNode, ans *[]int) {
    if node == nil {
        return
    }
    *ans = append(*ans, node.Val)
    dfsa144(node.Left, ans)
    dfsa144(node.Right, ans)
}


func main_LT0144_20211122() {
// func main() {

    fmt.Println("ans:")


}

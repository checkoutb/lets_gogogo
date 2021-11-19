// package sdq
package main

import (
    "fmt"
)


// public int totalSum(TreeNode root, int sum){
//     if(root == null) return 0;
//     if(root.left == null && root.right == null) return sum*10+root.val;
//     return totalSum(root.left,sum*10+root.val) + totalSum(root.right,sum*10+root.val);
// }

// int sumNumbers(TreeNode* root, int sum = 0) {
//     if(!root) return 0;
//     sum = sum*10 + root->val;
//     if(!root->left && !root->right) return sum;
//     return sumNumbers(root->left, sum) + sumNumbers(root->right, sum);
// }


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum Root to Leaf Numbers.
// Memory Usage: 2.2 MB, less than 15.08% of Go online submissions for Sum Root to Leaf Numbers.
func sumNumbers(root *TreeNode) int {
    ans := dfsa129(root, 0)
    return ans
}

func dfsa129(node *TreeNode, val int) int {
    val = val * 10 + node.Val
    if node.Left == nil && node.Right == nil {
        return val
    }
    ans := 0
    if node.Left != nil {
        ans += dfsa129(node.Left, val)
    }
    if node.Right != nil {
        ans += dfsa129(node.Right, val)
    }
    return ans
}

func main_LT0129_20211111() {
// func main() {

    fmt.Println("ans:")


}

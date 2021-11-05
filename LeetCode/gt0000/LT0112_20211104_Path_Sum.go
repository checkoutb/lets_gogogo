// package sdq
package main

import (
    "fmt"
)



// if (root == NULL) return false;
// if (root->val == sum && root->left ==  NULL && root->right == NULL) return true;
// return hasPathSum(root->left, sum-root->val) || hasPathSum(root->right, sum-root->val);

// if(root == NULL)  return false;
// if(root->right == root->left)  return sum == root->val;
// return hasPathSum(root->left, sum - root->val) || hasPathSum(root->right, sum - root->val);



// Runtime: 4 ms, faster than 90.11% of Go online submissions for Path Sum.
// Memory Usage: 4.7 MB, less than 44.95% of Go online submissions for Path Sum.
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    return dfsa112(root, 0, targetSum)
}

func dfsa112(node *TreeNode, now, tar int) bool {
    if node.Left == nil && node.Right == nil {
        return now + node.Val == tar
    }
    return (node.Left != nil && dfsa112(node.Left, now + node.Val, tar)) || (node.Right != nil && dfsa112(node.Right, now + node.Val, tar))
}


func main_LT0112_20211104() {
// func main() {

    fmt.Println("ans:")


}

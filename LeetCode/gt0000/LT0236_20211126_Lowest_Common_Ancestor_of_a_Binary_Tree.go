// package sdq
package main

import (
    "fmt"
)



// public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
//     if(root==null || root==p || root==q) return root;
//     TreeNode left=lowestCommonAncestor(root.left,p,q);
//     TreeNode right=lowestCommonAncestor(root.right,p,q);
//     if(left==null) return right;
//     if(right==null) return left;
//     return root;
// }


// if (root == null){
//     return root;
// }
// if (root.val == p.val || root.val == q.val){
//     return root;
// }
// TreeNode leftLCA = lowestCommonAncestor(root.left, p, q);
// TreeNode rightLCA = lowestCommonAncestor(root.right, p, q);
// if (leftLCA!=null && rightLCA!=null){
//     return root;
// }
// return (leftLCA!=null)?leftLCA:rightLCA;


// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//     if root == nil || root == p || root == q {
//         return root
//     }
//     left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
//     switch {
//     case left == nil && right == nil:
//         return nil
//     case left == nil && right != nil:
//         return right
//     case left != nil && right == nil:
//         return left
//     default:
//         return root
//     }
// }


// TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
//     if (!root || root == p || root == q) return root;
//     TreeNode* left = lowestCommonAncestor(root->left, p, q);
//     TreeNode* right = lowestCommonAncestor(root->right, p, q);
//     return !left ? right : !right ? left : root;
// }




// Runtime: 8 ms, faster than 95.26% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
// Memory Usage: 7.8 MB, less than 39.40% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 一个在 左子树，一个在右子树       或者就是这个结点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    ans, _ := dfsa236(root, p, q)
    return ans
}

func dfsa236(node *TreeNode, p, q *TreeNode) (*TreeNode, int) {
    if node == nil { return nil, 0 }
    t1, n1 := dfsa236(node.Left, p, q)
    if t1 != nil {
        return t1, 0
    }
    t2, n2 := dfsa236(node.Right, p, q)
    if t2 != nil { return t2, 0 }
    n3 := 0
    if node == q || node == p { n3 = 1 }
    if n1 + n2 + n3 > 1 {
        return node, 0
    }
    return nil, n1 + n2 + n3
}


func main_LT0236_20211126() {
// func main() {

    fmt.Println("ans:")


}

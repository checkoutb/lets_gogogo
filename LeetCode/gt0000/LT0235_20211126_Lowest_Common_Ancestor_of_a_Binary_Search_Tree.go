// package sdq
package main

import (
    "fmt"
)



// public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
//     while ((root.val - p.val) * (root.val - q.val) > 0)
//         root = p.val < root.val ? root.left : root.right;
//     return root;
// }



// public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
//     return (root.val - p.val) * (root.val - q.val) < 1 ? root :
//            lowestCommonAncestor(p.val < root.val ? root.left : root.right, p, q);
// }







// TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
//     if (root->val < p->val && root->val < q->val)
//         return lowestCommonAncestor(root->right, p, q);
//     if (root->val > p->val && root->val > q->val)
//         return lowestCommonAncestor(root->left, p, q);
//     return root;
// }


// if p.Val > root.Val && q.Val > root.Val {
//     return lowestCommonAncestor(root.Right, p, q)
// } else if p.Val < root.Val && q.Val < root.Val {
//     return lowestCommonAncestor(root.Left, p, q)
// } else {
//     return root
// }




// Runtime: 28 ms, faster than 13.88% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.
// Memory Usage: 7.7 MB, less than 8.20% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.
// ... 和 236 一样 ..
// 不 这个可以区间， divide， 如果这个区间里 包含 pq 继续divide。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    pv, qv := p.Val, q.Val
    if pv > qv {
        pv, qv = qv, pv
    }
	ans := dfsa235(root, pv, qv)
    return ans
}

// pv < qv
func dfsa235(node *TreeNode, pv, qv int) *TreeNode {
    // if node == nil || node == q || node == p { return node }
    if node == nil { return node }
    t2 := node.Val
    if pv <= t2 && qv >= t2 {
        return node
    }
    if qv < t2 {
        return dfsa235(node.Left, pv, qv)
    } else {
        return dfsa235(node.Right, pv, qv)
    }
}


func main_LT0235_20211126() {
// func main() {

    fmt.Println("ans:")


}

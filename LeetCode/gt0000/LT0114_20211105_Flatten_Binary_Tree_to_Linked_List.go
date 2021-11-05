// package sdq
package main

import (
    "fmt"
)


// for cur := root; cur != nil; cur = cur.Right {
//     if cur.Left != nil {
//         pre := cur.Left
//         for pre.Right != nil { pre = pre.Right }
//         pre.Right = cur.Right
//         cur.Right = cur.Left
//         cur.Left = nil
//     }
// }


// public void flatten(TreeNode root) {
//     if (root == null) return;    
//     TreeNode left = root.left;
//     TreeNode right = root.right;
//     root.left = null;
//     flatten(left);
//     flatten(right);
//     root.right = left;           // 现在 右子树 实际上是 左子树
//     TreeNode cur = root;
//     while (cur.right != null) cur = cur.right;           // 就是找 原先的 左子树的 最右结点。  但是 如果 碰到一个结点没有 右结点，有左结点。 那么 应该拿左结点吧?
//     cur.right = right;
// }



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Flatten Binary Tree to Linked List.
// Memory Usage: 2.9 MB, less than 53.21% of Go online submissions for Flatten Binary Tree to Linked List.
// pre-order
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    prev := &root
    l, r := root.Left, root.Right
    dfsa114(l, prev)
    dfsa114(r, prev)
}

func dfsa114(node *TreeNode, prev **TreeNode) {
    if node == nil {
        return
    }
    l, r := node.Left, node.Right
    (*prev).Right = node
    (*prev).Left = nil
    *prev = node

    dfsa114(l, prev)
    dfsa114(r, prev)
}


func main_LT0114_20211105() {
// func main() {

    fmt.Println("ans:")


}

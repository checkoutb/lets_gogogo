// package sdq
package main

import (
    "fmt"
)






// Runtime: 27 ms, faster than 60.87% of Go online submissions for Binary Search Tree Iterator.
// Memory Usage: 10.5 MB, less than 18.63% of Go online submissions for Binary Search Tree Iterator.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序
// BST，。 BST， 左子树均大于， 右子树均小于， 所以没有重复
type BSTIterator struct {
    Stk []*TreeNode
    // Lst int
}
func Constructor(root *TreeNode) BSTIterator {
    bi := BSTIterator{ []*TreeNode{} }
    for root != nil {
        bi.Stk = append(bi.Stk, root)
        root = root.Left
    }
    return bi
}
func (this *BSTIterator) Next() int {
    n2 := this.Stk[len(this.Stk) - 1]
    this.Stk = this.Stk[0 : len(this.Stk) - 1]
    n3 := n2.Right
    for n3 != nil {
        this.Stk = append(this.Stk, n3)
        n3 = n3.Left
    }
    return n2.Val
}
func (this *BSTIterator) HasNext() bool {
    return len(this.Stk) > 0
}
/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main_LT0173_20211124() {
// func main() {

    fmt.Println("ans:")


}

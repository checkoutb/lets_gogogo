// package sdq
package main

import (
    "fmt"
)



// root:=preorder[0] 
// var leftLen int
// for i, v :=range inorder{
//     if v==root{
//         leftLen = i
//         break
//     }
// }    
// pLeft, pRight := preorder[1:1+leftLen], preorder[1+leftLen:]
// inLeft,  inRight := inorder[:leftLen],inorder[1+leftLen:]
// 先序是  root + 左子树 + 右子树
// 中序是  左子树 + root + 右子树
// 所以能这样切分.







// Runtime: 140 ms, faster than 5.31% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
// Memory Usage: 4.7 MB, less than 10.03% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
// 先序, 中序.
// 先序的第一个 就是 root,  在 中序中 划分成  左右子树.
// 左右子树 中 第一个出现在 先序中的 就是 子树的root.  然后 在 中序中 继续划分 左右子树.
// 怎么快速知道 第一个出现呢.. preorder转map,  不能sort.  可以,增加一个 数组,sort下,但是 好费力,需要手动分组.而且... 或者最大堆,pop直到 出现在 现在的中序中的.
func buildTree(preorder []int, inorder []int) *TreeNode {
    map2 := map[int]int{}
    for i, v := range preorder {
        map2[v] = i
    }
    ans := dfsa105(map2, inorder)
    return ans
}

func dfsa105(map2 map[int]int, inorder []int) *TreeNode {
    if len(inorder) <= 0 {          // inorder[ 2 : 1 ]  ?
        return nil
    }
    vInMap := 100000        // min
    vIdx := -1

    for i, v := range inorder {
        if map2[v] < vInMap {
            vInMap = map2[v]
            vIdx = i
        }
    }
    var n TreeNode = TreeNode{inorder[vIdx], dfsa105(map2, inorder[0 : vIdx]), dfsa105(map2, inorder[vIdx + 1 : ])}
    return &n
}

func main_LT0105_20211103() {
// func main() {

    fmt.Println("ans:")



}

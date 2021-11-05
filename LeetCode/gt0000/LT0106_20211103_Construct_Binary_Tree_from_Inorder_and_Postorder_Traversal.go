// package sdq
package main

import (
    "fmt"
)



// 用map 记录 inorder的  <值:idx>.  这样就不需要for了

// private TreeNode generateTree(int[] postorder, Map<Integer, Integer> value2InIndex, int postStart, int postEnd, int inStart, int inEnd)
// 不需要 inorder


// root := &TreeNode{Val: postorder[lenPost - 1]}
// postorder = postorder[:lenPost - 1]
// for inPos, node := range inorder {
//     if root.Val == node {
//         root.Left = buildTree(inorder[:inPos],    postorder[:len(inorder[:inPos])])
//         root.Right = buildTree(inorder[inPos+1:], postorder[len(inorder[:inPos]):])
//     }
// }


// Runtime: 4 ms, faster than 93.06% of Go online submissions for Construct Binary Tree from Inorder and Postorder Traversal.
// Memory Usage: 4.3 MB, less than 63.19% of Go online submissions for Construct Binary Tree from Inorder and Postorder Traversal.
// 后序 的最后一个 是root..
func buildTree(inorder []int, postorder []int) *TreeNode {
    ans := dfsa106(inorder, postorder)
    return ans
}

func dfsa106(in []int, post []int) *TreeNode {
    if len(in) <= 0 {
        return nil
    }
    t2 := post[len(post) - 1]
    t3 := -1
    for i, v := range in {
        if v == t2 {
            t3 = i
            break
        }
    }
    return &TreeNode{in[t3], dfsa106(in[0 : t3], post[0 : t3]), dfsa106(in[t3 + 1 : ], post[t3 : len(post) - 1])}
}


func main_LT0106_20211103() {
// func main() {

    fmt.Println("ans:")


}

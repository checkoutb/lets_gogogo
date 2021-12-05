// package sdq
package main

import (
    "fmt"
    "strconv"
    "leetcode/utils"       // 得放 /usr/local/go/src/leetcode/utils，这个是GOROOT    放gopath里无效。。 怎么设置 project 的 gopath？ .. setting 搜索gopath。。。
)



// "1 2 3 n n " 






// Runtime: 14 ms, faster than 50.25% of Go online submissions for Serialize and Deserialize Binary Tree.
// Memory Usage: 7.9 MB, less than 35.18% of Go online submissions for Serialize and Deserialize Binary Tree.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//
type Codec struct {
    
}
func Constructor() Codec {
    return Codec{}
}
// Serializes a tree to a single string.
func (this *Codec) serialize(root *utils.TreeNode) string {
    ans := this.dfsa297(root)
    return ans
}
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *utils.TreeNode {    
    ans := this.dfsa297aaa(data[1 : len(data) - 1])
    return ans
}
// root.Val|left.cnt|right.cnt
// in-order
func (this *Codec) dfsa297(node *utils.TreeNode) string {
    if node == nil {
        return "()"
    }
    ans := "(" + this.dfsa297(node.Left)
    ans += strconv.Itoa(node.Val) 
    ans += this.dfsa297(node.Right) + ")"
    // fmt.Println(ans)
    return ans
}
func (this *Codec) dfsa297aaa(s string) *utils.TreeNode {
    if len(s) <= 2 {            // "()"
        return nil
    }
    cnt := 1
    ans := utils.TreeNode{}
    for i := 1; i < len(s); i++ {
        if s[i] == '(' {
            cnt++
        } else if s[i] == ')' {
            cnt--
        }
        if cnt == 0 {
            ans.Left = this.dfsa297aaa(s[1 : i])
            val := 0
            fff := 1
            i++
            for i < len(s) {
                if s[i] >= '0' && s[i] <= '9' {
                    val *= 10
                    val += int(s[i] - '0')
                } else if s[i] == '-' {
                    fff = -1                // ...
                } else {
                    ans.Val = val * fff
                    ans.Right = this.dfsa297aaa(s[i + 1 : len(s) - 1])
                    break
                }
                i++
            }
            break
        }
    }
    return &ans
}

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func main_LT0297_20211204() {
// func main() {

    fmt.Println("ans:")

    // n1 := &TreeNode{1,nil,nil}
    // n2 := &TreeNode{2,nil,nil}
    // n3 := &TreeNode{3,nil,nil}
    // n4 := &TreeNode{4,nil,nil}
    // n5 := &TreeNode{5,nil,nil}
    // // n1 := TreeNode{1,nil,nil}
    // n1.Left, n1.Right = n2, n3
    // n3.Left, n3.Right = n4, n5

    // sdq := Constructor()
    // s := sdq.serialize(nil)
    // fmt.Println(s)
    // nxx := sdq.deserialize(s)
    // fmt.Println(nxx)

    // [4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]
    arr := []int{4,-7,-3,utils.Zro,utils.Zro,-9,-3,9,-7,-4,utils.Zro,6,utils.Zro,-6,-6,utils.Zro,utils.Zro,0,6,5,utils.Zro,9,utils.Zro,utils.Zro,-1,-4,utils.Zro,utils.Zro,utils.Zro,-2}
    r := utils.Convert2TreeNode(arr)
    // utils.ShowTreeNode(r)

    sdq := Constructor()
    s := sdq.serialize(r)
    fmt.Println(s)
    nxx := sdq.deserialize(s)
    utils.ShowTreeNode(nxx)

}

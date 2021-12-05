// package sdq
package main

import (
    "fmt"
    "strings"
    "strconv"
)



// StringBuilder.setLength

// if root.Left == nil && root.Right == nil {
//     return []string{strconv.Itoa(root.Val)}
// }
// tmpLeft := binaryTreePaths(root.Left)
// for i := 0; i < len(tmpLeft); i++ {
//     res = append(res, strconv.Itoa(root.Val)+"->"+tmpLeft[i])
// }


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Paths.
// Memory Usage: 2.4 MB, less than 51.22% of Go online submissions for Binary Tree Paths.
func binaryTreePaths(root *TreeNode) []string {
    ans := []string{}
    dfsa257(root, []string{}, &ans)
    return ans
}

func dfsa257(node *TreeNode, arr []string, ans *[]string) {
    arr = append(arr, strconv.Itoa(node.Val))
    if node.Left == nil && node.Right == nil {
        str := strings.Join(arr, "->")
        *ans = append(*ans, str)
        arr = arr[0 : len(arr) - 1]
        return
    }

    if node.Left != nil {
        dfsa257(node.Left, arr, ans)
    }
    if node.Right != nil {
        dfsa257(node.Right, arr, ans)
    }
    arr = arr[0 : len(arr) - 1]
}


func main_LT0257_20211129() {
// func main() {

    fmt.Println("ans:")


}

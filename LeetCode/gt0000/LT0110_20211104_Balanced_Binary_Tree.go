// package sdq
package main

import (
    "fmt"
)





// bool isBalanced (TreeNode *root) {
//     if (root == NULL) return true;
//     int left=depth(root->left);
//     int right=depth(root->right);
//     return abs(left - right) <= 1 && isBalanced(root->left) && isBalanced(root->right);
// }


// Runtime: 8 ms, faster than 66.81% of Go online submissions for Balanced Binary Tree.
// Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Balanced Binary Tree.
func isBalanced(root *TreeNode) bool {
    ans, _ := dfsa110b(root)
    return ans
}

func dfsa110b(node *TreeNode) (bool, int) {
    if node == nil {
        return true, 1
    }
    l1, l2 := dfsa110b(node.Left)
    if !l1 {
        return false, l2 + 1
    }
    r1, r2 := dfsa110b(node.Right)
    if !r1 {
        return false, r2 + 1
    }
    if l2 > r2 {
        return l2 - r2 <= 1, l2 + 1
    } else {
        return r2 - l2 <= 1, r2 + 1
    }
}


// error
// [1,2,3,4,5,6,null,8]
//     1
//   2    3
// 4  5 6 nil
// 8
func isBalanced2(root *TreeNode) bool {
    map2 := map[int]int{}
    return dfsa110(root, 0, map2)
}

func dfsa110(node *TreeNode, depth int, map2 map[int]int) bool {
    if len(map2) > 2 {
        return false
    }
    if node == nil {
        map2[depth] = 1
        return len(map2) <= 2
    }
    return dfsa110(node.Left, depth + 1, map2) && dfsa110(node.Right, depth + 1, map2)
}


func main_LT0110_20211104() {
// func main() {

    fmt.Println("ans:")


}

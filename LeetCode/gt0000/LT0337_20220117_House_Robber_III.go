// package sdq
package main

import (
    "fmt"
)



// int rob(TreeNode* root, int& l, int& r) {
//     if (root == nullptr) return 0;
//     int ll = 0, lr = 0, rl = 0, rr = 0;
//     l = rob(root->left, ll, lr);
//     r = rob(root->right, rl, rr);
//     return max(root->val+ll+lr+rl+rr, l+r);
// }
// show. &



// private int[] robSub(TreeNode root) {
//     if (root == null) return new int[2];
    
//     int[] left = robSub(root.left);
//     int[] right = robSub(root.right);
//     int[] res = new int[2];

//     res[0] = Math.max(left[0], left[1]) + Math.max(right[0], right[1]);
//     res[1] = root.val + left[0] + right[0];
    
//     return res;
// }





// result = new int[2] {arrl[1] + arrr[1] + node->val, max(arrl[0], arrl[1]) + max(arrr[0], arrr[1])};

// Runtime: 4 ms, faster than 94.10% of Go online submissions for House Robber III.
// Memory Usage: 5.1 MB, less than 99.67% of Go online submissions for House Robber III.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 抢了一个，那么 它的子结点 就不能抢了。
// 所以 就2层。 抢子结点的 最大值 和 不抢子结点的最大值。
func rob(root *TreeNode) int {
    a1, _ := dfsa337(root)
    return a1
}

// 抢， 不抢
func dfsa337(node *TreeNode) (int, int) {
    if node == nil {
        return 0, 0
    }
    l1, l2 := dfsa337(node.Left)
    r1, r2 := dfsa337(node.Right)
    rob, notRob := l2 + r2 + node.Val, l1 + r1
    if notRob > rob {               // 好奇怪，感觉对的，又感觉错的。。。试试吧。。。      因为 这个，所以能保证 rob 是 最大值，实际上 不一定rob了这个结点。 。所以 父结点的 notRob 直接等于这2个 相加， 就不需要 再比较大小了 (父结点不rob，那么可能是 rob left + rob right or  rob_left + rob-right-left + rob-right-right or ...)。
        rob = notRob
    }
    return rob, notRob
}

// 3
// 23
// 31
func main_LT0337_20220117() {
// func main() {

    fmt.Println("ans:")


}

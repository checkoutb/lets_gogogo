// package sdq
package main

import (
    "fmt"
    "leetcode/utils"
)



// func dfs(node *ListNode, curr **ListNode, start, end int) *TreeNode {
//     if end < start {
//         return nil
//     }
//     mid := start + (end - start)/2 
//     left := dfs(node, curr, start, mid-1)
//     val := (*curr).Val 
//     *curr = (*curr).Next                         // dfs，从第一个开始处理。
//     right := dfs(node, curr, mid+1, end)
//     return &TreeNode{
//         Val: val,
//         Left: left,
//         Right: right,
//     }
// }


// if(head==tail) return null;
// while(fast!=tail&&fast.next!=tail){
//     fast = fast.next.next;
//     slow = slow.next;
// }
// TreeNode thead = new TreeNode(slow.val);
// thead.left = toBST(head,slow);
// thead.right = toBST(slow.next,tail);



// Runtime: 4 ms, faster than 95.40% of Go online submissions for Convert Sorted List to Binary Search Tree.
// Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Convert Sorted List to Binary Search Tree.

// ... remove utils.

func sortedListToBST(head *utils.ListNode) *utils.TreeNode {
    if head == nil {
        return nil
    }
    t2 := head
    cnt := 0
    for t2 != nil {         // convert to array...
        cnt++
        t2 = t2.Next
    }
    return dfsa109(head, cnt)
}

func dfsa109(head *utils.ListNode, sz1 int) *utils.TreeNode {
    // fmt.Println(sz1, ".sz")
    if sz1 <= 0 {
        return nil
    }
    // if sz1 == 1 {
    //     return &TreeNode{head.Val, nil, nil}
    // }
    mid := sz1 / 2
    t2 := head
    // fmt.Println(mid, " . . ")
    // for mid > 0 {
    for i := 0; i < mid; i++ {
        t2 = t2.Next
        // mid--
    }
    // fmt.Println(t2, ", ", mid)
    return &utils.TreeNode{t2.Val, dfsa109(head, mid), dfsa109(t2.Next, sz1 - mid - 1)}
}

func main_LT0109_20211104() {
// func main() {

    fmt.Println("ans:")

    arr := []int{-10,-3,0,5,9}

    head := utils.ConvertToListNode(arr)
    utils.ShowListNode(head)

    root := sortedListToBST(head)

    utils.ShowTreeNode(root)

}

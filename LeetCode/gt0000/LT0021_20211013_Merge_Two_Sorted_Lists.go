// package sdq
package main

import (
    "fmt"
)


    // if l1.Val > l2.Val {
    //     l1, l2 = l2, l1
    // }
    // curr.Next = l1
    // curr, l1 = curr.Next, l1.Next
// if l2 != nil {
//     curr.Next = l2
// }

// dummy := &ListNode{
//     Val: 0,
//     Next: ln,
// }


// if(l1 == null) return l2;
// if(l2 == null) return l1;
// if(l1.val < l2.val){
//     l1.next = mergeTwoLists(l1.next, l2);
//     return l1;
// } else{
//     l2.next = mergeTwoLists(l1, l2.next);
//     return l2;
// }


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
// Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Merge Two Sorted Lists.
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var dummy *ListNode = &ListNode{-1, nil}
    p := dummy
    for l1 != nil || l2 != nil {
        t1, t2 := getVal(l1), getVal(l2)
        if t1 < t2 {
            p.Next, l1 = l1, l1.Next
            p = p.Next
        } else {
            p.Next, l2 = l2, l2.Next         // vscode怎么知道 ListNode有Next属性的？。。 go run 就是说 undefined
            p = p.Next
        }
    }
    return dummy.Next
}
func getVal(node *ListNode) int {
    if node == nil {
        return 101
    }
    return (*node).Val
}


func main_LT0021_20211013() {
// func main() {

    fmt.Println("ans:")

    l1 := &ListNode{-1, nil}

    ans := mergeTwoLists(nil, l1)

    fmt.Println(ans)

}

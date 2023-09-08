// package sdq
package main

import (
    "fmt"
)

// D D

// slow, fast := head, head
// for fast != nil && fast.Next != nil {
//     fast = fast.Next.Next
//     if fast == slow {
//         return true
//     }
//     slow = slow.Next
// }



// Runtime8 ms
// Beats
// 69.8%
// Memory4.4 MB
// Beats
// 97.97%

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    // 忘记怎么声明变量了。。。。

    var fst *ListNode = head;
    var slw *ListNode = head;
    
    // while (fst.Next != nil && fst.Next.Next != nil) {
    for fst != nil && fst.Next != nil && fst.Next.Next != nil {
        fst = fst.Next.Next;
        slw = slw.Next;
        if (slw == fst) {
            return true;
        }
    }
    return false;
 }


func main_LT0141_20230904() {
// func main() {

    fmt.Println("ans:")


}

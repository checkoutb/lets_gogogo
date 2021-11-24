// package sdq
package main

import (
    "fmt"
)


// ListNode pA = headA, pB = headB;
// while (pA != pB) {
//     pA = pA == null? headB: pA.next;
//     pB = pB == null? headA: pB.next;
// }
// return pA;
// ... 一个指针从 先遍历A，然后遍历B， 一个指针先遍历B，再遍历A。   相等 或 都是nil 的时候 就是交点。。。


// 一个list 转map



// Runtime: 46 ms, faster than 26.79% of Go online submissions for Intersection of Two Linked Lists.
// Memory Usage: 7.7 MB, less than 26.32% of Go online submissions for Intersection of Two Linked Lists.
// 没有环就给他一个环 b.tail.Next = a
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    t2 := headA
    for t2.Next != nil {
        t2 = t2.Next
    }
    t2.Next = headB
    f, s := headA, headA
    for f != nil {
        s = s.Next
        f = f.Next
        if f != nil {
            f = f.Next
        }
        if s == f {
            break
        }
    }
    if f == nil {
        t2.Next = nil
        return nil
    }
    f = headA
    for f != s {
        f = f.Next
        s = s.Next
    }
    t2.Next = nil
    return f
}

func main_LT0160_20211124() {
// func main() {

    fmt.Println("ans:")


}

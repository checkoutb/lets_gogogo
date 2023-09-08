// package sdq
package main

import (
    "fmt"
)





// Runtime1 ms
// Beats
// 77.48%
// Memory2.1 MB
// Beats
// 96.70%

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 1 2 3 4

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left == right {
        return head;
    }
    dum := ListNode{-1, head};
    
    right -= left;

    ptr := &dum;        // left's previsou

    for left > 1 {
        ptr = ptr.Next;
        left -= 1;
    }

    ptr2 := ptr.Next;        // after reverse, sublist's tail
    ptr3 := ptr;       // previous of ptr4
    ptr4 := ptr.Next;       // now,   first element before reverse, last ele after reverse
    ptr5 := ptr.Next.Next;       // next of ptr4
    ptr6 := ptr;        // temp
    // ptr3 <- ptr4  ptr5->
    for right > 0 {

        // fmt.Println(dum);
        // fmt.Println(ptr3.Val, ptr4.Val);
        // if ptr5 != nil {
        //     fmt.Println(ptr5.Val);
        // }

        ptr6 = nil;
        if ptr5 != nil {
            ptr6 = ptr5.Next;
        }
        ptr3 = ptr4;
        ptr4 = ptr5;
        ptr5 = ptr6;
        ptr4.Next = ptr3;

        // ptr4 = ptr4.Next;
        // ptr5 = ptr4.Next;
        // ptr4.Next = ptr3;
        // ptr3 = ptr4;
        // ptr4 = ptr5;
        right -= 1;
    }

    // fmt.Println(ptr.Val, ptr2.Val, ptr4.Val);

    ptr2.Next = ptr5;
    ptr.Next = ptr4;

    return dum.Next;
}


func main_LT0092_20230907() {
// func main() {

    fmt.Println("ans:")


}

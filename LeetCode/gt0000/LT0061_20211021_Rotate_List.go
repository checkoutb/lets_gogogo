// package sdq
package main

import (
    "fmt"
)




// for (i=0;fast.next!=null;i++)//Get the total length 
// fast=fast.next;
// for (int j=i-n%i;j>0;j--) //Get the i-n%i th node
// slow=slow.next;
// slow.next 就是 结果。



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate List.
// Memory Usage: 2.6 MB, less than 60.31% of Go online submissions for Rotate List.
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    sz1 := 0
    np := head
    tail := head
    for np != nil {
        sz1++
        tail = np
        np = np.Next
    }
    k = k % sz1
    if k == 0 {
        return head
    }
    k = sz1 - k                 // ... 向右了。。。。
    pre, node := np, head           // np is nil
    for k > 0 {
        k--
        pre, node = node, node.Next
    }
    tail.Next, pre.Next = head, nil

    return node
}



// 这个是 往左的。。。 不清楚 是不是正确的 往左。。。
func rotateRight2(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    sz1 := 0
    np := head
    tail := head
    for np != nil {
        sz1++
        tail = np
        np = np.Next
    }
    k = k % sz1
    if k == 0 {
        return head
    }
    pre, node := np, head           // np is nil
    for k > 0 {
        k--
        pre, node = node, node.Next
    }
    tail.Next, pre.Next = head, nil

    return node
}



func main_LT0061_20211021() {
// func main() {

    fmt.Println("ans:")


}

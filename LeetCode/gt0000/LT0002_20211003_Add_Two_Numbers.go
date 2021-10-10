
package main

import (
    "fmt"
    // "./utils"
    // "./utils/utils"
    // "./LeetCode/gt0000/utils/utils"
)
// cao.. 这是什么包管理。反正就是导不进去。


type ListNode struct {
    Val int
    Next *ListNode
}



// ... HOW TO IMPORT MY UTILS.....





// Runtime: 12 ms, faster than 59.91% of Go online submissions for Add Two Numbers.
// Memory Usage: 4.8 MB, less than 93.16% of Go online submissions for Add Two Numbers.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // var node *ListNode = nil
    var node *ListNode = &ListNode{-1, nil}
    ans := node
    carry := 0
    for l1 != nil || l2 != nil {
        // t2 := node
        // t3 = carry + (0 if l1 == nil else l1.Val) + (0 if l2 == nil else l2.Val)
        t3 := carry
        if l1 != nil {
            t3 += l1.Val
            l1 = l1.Next
        }
        if (l2 != nil) {
            t3 += l2.Val
            l2 = l2.Next
        }
        // node = make(ListNode)
        // node = &ListNode{t3 % 10, node}
        node.Next = &ListNode{t3 % 10, nil}
        node = node.Next
        // node.Val = t3 % 10
        carry = t3 / 10
        // node.Next = t2
    }
    if carry != 0 {
        // node = ListNode(carry, node)
        // node = &ListNode{carry, node}
        node.Next = &ListNode{carry, nil}
    }
    return ans.Next
}


func main_LT0002_20211003() {
// func main() {

    // arr1 := []int{2,4,3}
    // arr2 := []int{5,6,4}

    // arr1 := []int{0}
    // arr2 := []int{0}

    arr1 := []int{9,9,9,9,9,9,9}
    arr2 := []int{9,9,9,9}

    var l1 *ListNode = nil
    var l2 *ListNode = nil

    // for _, val := range arr1 {
    for i := len(arr1) - 1; i >= 0; i-- {
    // for i := 0; i < len(arr1); i++ {
        val := arr1[i]
        l1 = &ListNode{val, l1}
    }

    // for _, val := range arr2 {          // no in.
    for i := len(arr2) - 1; i >= 0; i-- {
    // for i := 0; i < len(arr2); i++ {
        val := arr2[i]
        l2 = &ListNode{val, l2}
    }

    // l1 = &ListNode(1, ListNode(1, nil))
    // l2 = LsitNode(2, ListNode(1, nil))

    fmt.Println("ans:")

    l3 := addTwoNumbers(l1, l2)

    for l3 != nil {
        fmt.Println(l3.Val)
        l3 = l3.Next
    }

}

package utils

import (
    "fmt"
    // "strconv"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func convertToListNode(arr []int) *ListNode {
    dummy := &ListNode{-1, nil}
    np := dummy
    for _, val := range arr {
        np.Next = &ListNode{val, nil}
        np = np.Next
    }
    return dummy.Next
}

func showListNode(np *ListNode) {           // no void
    for np != nil {
        // fmt.Printf(np.Val)           // prinf 只能string， format。 println 可以int
        // fmt.Printf(", ")
        fmt.Printf("%d, ", np.Val)
        np = np.Next
    }
    fmt.Println()
}

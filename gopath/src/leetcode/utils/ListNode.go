package utils

import (
    "fmt"
    // "strconv"
)

// 首字母 大写公开,小写私有.

type ListNode struct {          // 外部使用的时候 好像必须带 包名前缀...等于没用.. 而且下面的 类型都是 这个 类型..
    Val int
    Next *ListNode
}

func ConvertToListNode(arr []int) *ListNode {
    dummy := &ListNode{-1, nil}
    np := dummy
    for _, val := range arr {
        np.Next = &ListNode{val, nil}
        np = np.Next
    }
    return dummy.Next
}

func ShowListNode(np *ListNode) {           // no void
    for np != nil {
        // fmt.Printf(np.Val)           // prinf 只能string， format。 println 可以int
        // fmt.Printf(", ")
        fmt.Printf("%d, ", np.Val)
        np = np.Next
    }
    fmt.Println()
}

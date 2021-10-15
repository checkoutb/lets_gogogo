// package sdq
package main

import (
    "fmt"
)




// for currentPosition := head; currentPosition != nil; currentPosition = currentPosition.Next {   
//     // fmt.Printf("visiting: %v\n", currentPosition)
//     if i++; i % k == 0 {
//
// for cursor != tailNext {
//     // fmt.Printf("inside visite n=%v\n", cursor)
//     nextNode = cursor.Next
//     cursor.Next = lastNode
//     lastNode = cursor
//     cursor = nextNode
// }
// 先取到k个，然后原地反转k个，

// if i++; i % k == 0 {
    // 那么 for i++; i < 10 {}  应该也可以。


// for _ in range(k):
//     cur.next, cur, pre = pre, cur.next, cur  # standard reversing
// jump.next, jump, l = pre, l, r  # connect two k-groups




// Runtime: 4 ms, faster than 94.37% of Go online submissions for Reverse Nodes in k-Group.
// Memory Usage: 3.8 MB, less than 33.33% of Go online submissions for Reverse Nodes in k-Group.
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{-1, head}
    b2 := true
    np := dummy
    for b2 {
        t2, t3, b1 := reverse_k_a1(np.Next, k)          // 不能直接  b2 :=  会报错 b2定义了，但是没有使用， 应该是 := 创建了新的b2.
        // fmt.Println(t2, t3, np)
        b2 = b1
        np.Next = t2
        np = t3
        // showListNode(dummy)
    }
    return dummy.Next
}

// head, tail, bool
// 或者直接外面比较 返回值 是否等于 xx.Next， 等于就是结束。
func reverse_k_a1(node *ListNode, k int) (*ListNode, *ListNode, bool) {
    if node == nil {
        return node, node, false
    }
    np := node
    orik := k
    for k > 1 {
        k--
        np = np.Next
        if np == nil {
            return node, np, false
        }
    }
    tmp := np.Next
    // fmt.Println(np.Val, "1123")
    h, t := dfsa3(node, orik)
    // t.Next = np.Next                    // .. np 所指向的对象 的 next 在 dfsa3 里被修改了。。 。。。。
                                // 无法确保 属性/指针 不被修改， 所以 尽量提前保存。
    t.Next = tmp
    // fmt.Println(t, np)
    return h, t, true
}

// head, tail
func dfsa3(node *ListNode, k int) (*ListNode, *ListNode) {
    if k == 1 {
        return node, node
    }
    h, t := dfsa3(node.Next, k - 1)
    // t.Next, node.Next = node, nil
    // fmt.Println(t, node, k)
    t.Next = node
    // fmt.Println(t, node)
    node.Next = nil
    // fmt.Println(t, node)                // kao,  node 和 t 是同一个，导致 node.Next 把 t.next 也变成nil了， 然后外面 空指针了。。
                                // 但是 返回的t
                                // 可能是外面 循环了，导致这里 node == t 了。。  所以不是这里的原因。。。。。
    // if node == nil {
    //     fmt.Println("asdasd")
    // }
    // if t.Next == nil {
    //     fmt.Println("zxczxc", t.Val, node.Val, t, node)
    // }
    return h, t.Next
}



func main_LT0025_20211014() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{1,2,3,4,5,6,7}
    // arr := []int{1,2,3}
    // k := 3
    arr := []int{1}
    k := 1

    np := convertToListNode(arr)

    ans := reverseKGroup(np, k)

    showListNode(ans)
}



type ListNode struct {              // 这个在vscode中 是 redeclared， 但是 go run 需要这个
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

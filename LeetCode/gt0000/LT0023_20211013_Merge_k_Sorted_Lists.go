// package sdq
package main

import (
    "fmt"
)

// merge
// divide and conquer
// 不停切片，直到 切片的 len <= 2,  0返回nil， 1返回[0], 2就mergeTwoList， 然后返回。


// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 { return nil}
// 	headMap := make(map[int]*ListNode,0)
// 	tailMap := make(map[int]*ListNode,0)
// 	for _,list :=  range lists {
// 		cur := list
// 		for cur != nil {
// 			temp := cur.Next

// 			head := headMap[cur.Val]
// 			tail := tailMap[cur.Val]
// 			if head == nil {
// 				headMap[cur.Val] = cur              // 这个Val 的 第 一 次 出现的 node。 并且 node 没有后继。
// 				headMap[cur.Val].Next = nil
// 			}
// 			if tail != nil {
// 				tailMap[cur.Val].Next = cur         // 应该是 所有的 Val 这个值 的 所有 node。  不不不，这里的 被下面 刷新了。
// 			}
// 			tailMap[cur.Val] = cur
// 			tailMap[cur.Val].Next = nil             // 所以 tailMap 里 也是 最后一次出现的， 且 没有后继。 
// 			cur = temp
// 		}
// 	}
// 	if len(headMap) == 0 { return nil}
	
// 	slice := []int{}
// 	for key, _ := range headMap {
// 		slice = append(slice, key)
// 	}
// 	sort.Ints(slice)

// 	head := headMap[slice[0]]
// 	tail := tailMap[slice[0]]
// 	for i := 1 ; i < len(headMap) ; i++ {
// 		tail.Next = headMap[slice[i]]
// 		tail = tailMap[slice[i]]
// 	}
// 	return head
// }
// 这个是 桶排序， 桶的key是 Val， 里面保存了 链表。
// headMap 是 这个 桶里的 链表的头   为了在后续 进行拼接。
// tailMap 是 这个 桶里的 链表的尾  这样 往桶里 增加Node 时 比较简单。
// 然后 后续就是 把 Val sort， 然后遍历， 然后 取出 链表， 并且 想链接。



// Runtime: 84 ms, faster than 31.96% of Go online submissions for Merge k Sorted Lists.
// Memory Usage: 5.4 MB, less than 81.01% of Go online submissions for Merge k Sorted Lists.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 最好是一个 pri_que, 保存pair<int, ListNode*>  。。。 有个heap，也给了一个 heap实现的pri_queue的代码。。
//                                                  https://golang.google.cn/pkg/container/heap/
// struct Pair {
//     val int
//     next *ListNode
// }
func mergeKLists(lists []*ListNode) *ListNode {
    dummy := &ListNode{-1, nil}
    arr := lists
    p := dummy

    AAA:
    for len(arr) > 0 {
        mnidx := -1
        mn := 10001
        cnt := 0
        sz1 := len(arr)
        for i := 0; i < (sz1 - cnt); i++ {
            for arr[i] == nil && i < (sz1 - cnt) {
                // fmt.Printf("1")
                cnt++
                arr[i] = arr[sz1 - cnt]
            }
            if arr[i] == nil {          // 
                // fmt.Printf("2")
                // cnt++
                break;
            }
            val := arr[i].Val
            if val < mn {
                mn = val
                mnidx = i
            }
        }
        // fmt.Println(cnt)
        // fmt.Println(len(arr))
        // fmt.Println(mnidx)
        // fmt.Println(arr[0] == nil)           // 最后一次  true
        if mnidx == -1 {
            break AAA           // 所以 AAA这个for 是死循环， 必须得手动
        }
        p.Next, arr[mnidx] = arr[mnidx], arr[mnidx].Next            // 修改了 arr[mnidx]，所以导致 最后一次的时候 上面for的方法体不会执行。
        p = p.Next

        if cnt != 0 {
            arr = arr[0 : sz1 - cnt]
            // fmt.Println(len(arr))
        }
    }

    return dummy.Next
}
// func getVal2(node *ListNode) int {
//     if node == nil {
//         return 10001
//     } else {
//         return (*node).Val
//     }
// }

func main_LT0023_20211013() {
// func main() {

    fmt.Println("ans:")

    arr := [][]int{{1,4,5},{1,3,4},{2,6}}

    // [[],[-1,5,11],[],[6,10]]
    // arr := [][]int{{},{-1,5,11},{},{6,10}}

    narr := []*ListNode{}

    for _, v := range arr {
        narr = append(narr, convertToListNode(v))
    }

    for _, v := range narr {
        showListNode(v)
    }

    ans := mergeKLists(narr)

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


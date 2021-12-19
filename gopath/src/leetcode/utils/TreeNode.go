
package utils

import "fmt"
import "container/list"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var (
	Zro = -987897897
)

// 1，2，3,-1,-1,4,5
func Convert2TreeNode(nodes []int) *TreeNode {
	ans := &TreeNode{nodes[0],nil,nil}
	idx := 1
	arr := []*TreeNode{ ans }
	for len(arr) > 0 {
		sz1 := len(arr)
		for n := 0; n < sz1; n++ {
			if idx < len(nodes) {
				t2 := nodes[idx]
				if t2 != Zro {
					arr[n].Left = &TreeNode{ nodes[idx], nil, nil }
					arr = append(arr, arr[n].Left)
				}
				idx++
				if idx < len(nodes) {
					t2 = nodes[idx]
					if t2 != Zro {
						arr[n].Right = &TreeNode{ nodes[idx], nil, nil }
						arr = append(arr, arr[n].Right)
					}
					idx++
				} else {
					break
				}
			} else {
				break
			}
		}
		arr = arr[sz1 : ]
	}
	return ans
}

func ShowTreeNode(node *TreeNode) {
	que := list.New()
	que.PushBack(node)
	emp := &TreeNode{-1, nil, nil}
	allEmp := false
	for que.Len() > 0 && !allEmp {
		sz1 := que.Len()
		allEmp = true
		for sz1 > 0 {
			// var e *Element = que.Front()				// ????? 根本没有说哪里导入啊.
			// e := que.Front()
			var e *list.Element = que.Front()		// 既然能直接 Element,那么说明在一个包中.
			
			n2 := e.Value.(*TreeNode)			// ...    .(Type)
			if n2 == emp {			// 可以直接用nil吧
				que.PushBack(emp)
				que.PushBack(emp)
				fmt.Print("nil, ")
			} else {
				if n2.Left != nil {
					que.PushBack(n2.Left)
				} else {
					que.PushBack(emp)
				}
				if n2.Right != nil {
					que.PushBack(n2.Right)
				} else {
					que.PushBack(emp)
				}
                fmt.Print(n2.Val, ", ")
				allEmp = false
			}
			sz1--
            que.Remove(e)
		}
        fmt.Println()
	}
}

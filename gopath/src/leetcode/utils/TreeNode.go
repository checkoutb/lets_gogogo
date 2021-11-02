
package utils

import "fmt"
import "container/list"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
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

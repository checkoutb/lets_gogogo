// package sdq
package main

import (
    "fmt"
    "container/list"
)






// vector<TreeNode*> ans;
// for (int i = st; i <= en; ++i)
// {
//     vector<TreeNode*> vl = dfsa1(st, i - 1);
//     vector<TreeNode*> vr = dfsa1(i + 1, en);
//     for (TreeNode* l : vl)
//     {
//         for (TreeNode* r : vr)
//         {
//             TreeNode* n = new TreeNode(i);
//             n->left = l;
//             n->right = r;
//             ans.push_back(n);
//         }
//     }
// }


// gg


// func generateTrees(n int) []*TreeNode {
//     ans := &[]*TreeNode{}
//     dfsa95(1, n, n, []*TreeNode{}, ans, nil)
//     return *ans
// }

// func dfsa95a(st, en, n int, )




// error 

// insert 还是 遍历可能的值 ?
// 怎么复制?
func generateTrees2(n int) []*TreeNode {
    ans := &[]*TreeNode{}
    dfsa95(1, n, n, []*TreeNode{}, ans, nil)
    return *ans
}

func dfsa95(st, en, n int, arr []*TreeNode, ans *[]*TreeNode, parent *TreeNode) {
    // fmt.Println("., ", len(arr))
    // if parent != nil {
    //     fmt.Println("..pp... ", parent.Val)
    // }

    if len(arr) == n {
        t2 := make([]*TreeNode, n)
        copy(t2, arr)                       // 这个,t2的子结点 是指向 arr的节点的吧?

        // tx := t2[0].Left
        // if tx != nil {
        //     fmt.Println(" == -- 123 = ")
        //     for _, n := range arr {
        //         if &tx == &n {                        // 同时在 arr  和 t2 中.. 那就是 == 的 判断方式的问题了.   
        //                                             // 但是 &tx == &n   永远不成立...     而且 上面已经是 指针了.     想起来了, 好像自动转换的,
        //                                             //      使用指针 可以像 普通变量 那样 直接调用方法...
        //                                             // 无法理解这种设计.  强类型, 还不带数值类型的自动升级, 结果 指针和对象 混在一起..
        //                                             // baidu 不到 如何判断 2个变量 是否是同一个对象.....
        //                                             // 这里是复制了指针.  所以 指针的值 确实可以 都存在.
        //                                             // 但是为什么 & 就都不存在了呢
        //             fmt.Println(" in arr")
        //         }
        //     }
        //     for _, n := range t2 {
        //         // fmt.Println(n)
        //         if &tx == &n {
        //             fmt.Println("  in        t2")
        //         }
        //     }
        //     fmt.Println(" == --  = ")
        // }

        *ans = append(*ans, t2[0])
        // fmt.Println(" . ..  ", len(*ans))
        return
    }
    if st > en {
        return
    }
    for i := st; i <= en; i++ {
        node := TreeNode{i, nil, nil}
        if parent != nil {
            if i > parent.Val {
                parent.Right = &node
            } else {
                parent.Left = &node
            }
        }
        arr = append(arr, &node)
        t4 := make([]*TreeNode, len(arr))           // 问题就是在这种 复制, 指针.
        copy(t4, arr)
        dfsa95(st, i - 1, n, t4, ans, &node)
        dfsa95(i + 1, en, n, t4, ans, &node)
        arr = arr[0 : len(arr) - 1]
    }
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func main_LT0095_20211102() {
// func main() {

    fmt.Println("ans:")

    n := 3

    ans := generateTrees(n)

    fmt.Println(len(ans))

    for _, n := range ans {
        ShowTreeNode(n)
    }

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

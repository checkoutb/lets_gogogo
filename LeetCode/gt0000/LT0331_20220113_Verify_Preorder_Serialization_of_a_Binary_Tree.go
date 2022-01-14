// package sdq
package main

import (
    "fmt"
)



// int sum = 0;
// for (char c: preorder) {
//     if (sum <= -2)
//         return false;
//     if (c == ',')
//         sum += 1;
//     if (c == '#')
//         sum -= 2;
// }
// return sum == -2;

// vector<int> stk;
// for (int i = 0, j = 0; j <= preorder.size(); ++j) {
//     if (j == preorder.size() || preorder[j] == ',') {
//         if (j == i + 1 && preorder[i] == '#') {
//             stk.emplace_back(2);
//         } else {
//             stk.emplace_back(0);
//         }
//         while (stk.size() >= 2 && stk.back() == 2) {
//             stk.pop_back();
//             ++stk.back();
//         }
//         i = j + 1;
//     }
// }
// return stk.size() == 1 && stk.back() == 2;

// 好像 并不需要 中途判断 cnt <= 0...


// public boolean isValidSerialization(String preorder) {
//     String s = preorder.replaceAll("\\d+,#,#", "#");
//     return s.equals("#") || !s.equals(preorder) && isValidSerialization(s);
// }


// int need = 1;
// for (String val : preorder.split(",")) {
//     if (need == 0)
//         return false;
//     need -= " #".indexOf(val);
// }
// return need == 0;


// We just need to remember how many empty slots we have during the process.
// Initially we have one ( for the root ).
// for each node we check if we still have empty slots to put it in.
//     a null node occupies one slot.
//     a non-null node occupies one slot before he creates two more. the net gain is one.
// null 占一个， 非空先占一个，然后创建2个。



// all non-null node provides 2 outdegree and 1 indegree (2 children and 1 parent), except root
// all null node provides 0 outdegree and 1 indegree (0 child and 1 parent).
//
// String[] nodes = preorder.split(",");
// int diff = 1;
// for (String node: nodes) {
//     if (--diff < 0) return false;
//     if (!node.equals("#")) diff += 2;
// }
// return diff == 0;


// Runtime: 4 ms, faster than 18.18% of Go online submissions for Verify Preorder Serialization of a Binary Tree.
// Memory Usage: 2.4 MB, less than 81.82% of Go online submissions for Verify Preorder Serialization of a Binary Tree.
// 先序
// 紧跟后面的，肯定是 左子结点。 如果 紧跟着后面的(左子结点) 为#，那么 再后面一个 就是 右子结点。
// 叶子结点 肯定是后跟 #,#    再后面一个 就是 上N级的 右子结点。  stack？
// 9,3,4,#,#,1,#,#,2,#,6,#,#
// 9 push， 等待 右子结点
// 3 push，
// 4 push
// 4 pop  第二个#
// 3 pop  1
// 1 push
// 1 pop  第四个#
// 2 push
// 2 pop  6
// 6 push
// 6 pop 
// 剩 个 9...  2的时候 需要 pop 9.  每个结点 进来的时候 需要 pop掉。 不不不，3呢。  所以 初始化 应该是 无限 left 压栈。 
// 不不不 是每次push的时候 都无限left。 然后 忽视一个#(这个是 左子结点) 。 其他的 每个 #，正常结点 都 pop一个。 当然正常结点 还是 无限left的。
func isValidSerialization(preorder string) bool {
    arr := []bool{}      // isNull ?
    for i, sz1 := 0, len(preorder); i < sz1; {
        if preorder[i] == '#' {
            i += 2
            arr = append(arr, true)
        } else {
            for i < sz1 && preorder[i] != ',' {
                i++
            }
            i++
            arr = append(arr, false)
        }
    }
    if len(arr) == 1 && arr[0] {
        return true
    }
    // fmt.Println(arr)
    // stk := []int{}
    // idx, sz1 := 0, len(preorder)
    cnt := 0
    for i, sz1 := 0, len(arr); i < sz1; i++ {
        if arr[i] {     // is null
            // if len(stk) == 0 {
            //     return false
            // }
            // stk =                .. just a count...
            if cnt <= 0 {
                // fmt.Println("... ", i)
                return false
            }
            cnt--           // 我怕 森林。 会不会 0 以后 ，又可以加，好像不行。
        } else {    // not null
            if i != 0 {
                if cnt <= 0 {
                    // fmt.Println("==== ", i)
                    return false
                }
                cnt--
            }
            for i < sz1 && !arr[i] {
                // stk = append(stk, 1)
                i++
                cnt++
            }
            // for ;;i++ 跳过 左子结点。
        }
    }
    // fmt.Println(" - - - ", cnt)
    return cnt == 0
}


func main_LT0331_20220113() {
// func main() {

    fmt.Println("ans:")

    // s := "9,3,4,#,#,1,#,#,2,#,6,#,#"
    // s := "1,#"
    // s := "9,#,#,1"
    s := "#"

    ans := isValidSerialization(s)

    fmt.Println(ans)

}

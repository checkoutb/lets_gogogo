// package sdq
package main

import (
    "fmt"
    "container/list"
)


// .同层,不需要任何操作,  ..上级,pop,    其他 下级, push....  但是.
// fallthrough  默认break
func simplifyPath(path string) string {
    stk := list.New()
    for i := 0; i < len(path); i++ {
        switch path[i] {
        case '/':
            for i < len(path) && path[i] == '/' {
                i++
            }
            i--
        case '.':
            t2 := 1
            orii := i
            for i < len(path) && path[i] == '.' {
                t2++
                i++
            }
            i--
            if t2 > 2 {
                i = orii
                // fallthrough         // 不能这样用的...要直接在case 内层.
                // goto AAA         // 不能goto
            } else if t2 == 2 {
                stk.Remove(stk.Back())
            }
        default:
            // AAA:
            orii := i
            for i < len(path) && path[i] != '/' {
                i++
            }
            s2 := path[orii : i]
            i--
            stk.PushBack(s2)
        }
    }
    ans := ""
    for e := stk.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
        // ans += e.Value + "/"
    }
    return ans[0 : len(ans) - 1]
}


// func main_LT0071_20211022() {
func main() {

    fmt.Println("ans:")

    s := "/a/./b/../../c/"

    ans := simplifyPath(s)

    fmt.Println(ans)

}

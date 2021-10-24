// package sdq
package main

import (
    "fmt"
    "container/list"
)




// dirs := strings.Split(path, "/")
// stack := []string{}
// stack = stack[:len(stack)-1]
// stack = append(stack, dir)
// return "/" + strings.Join(stack, "/")

// return fmt.Sprintf("/%s", strings.Join(files, "/"))


// Runtime: 6 ms, faster than 17.36% of Go online submissions for Simplify Path.
// Memory Usage: 5 MB, less than 6.61% of Go online submissions for Simplify Path.
func simplifyPath(path string) string {
    stk := list.New()
    str := ""
    path += "/"
    for i := 0; i < len(path); i++ {
        if path[i] == '/' {
            if len(str) != 0 {
                switch str {
                case ".":
                    ;
                case "..":
                    if stk.Len() > 0 {
                        stk.Remove(stk.Back())
                    }
                default:
                    stk.PushBack(str)
                }
            }
            for i < len(path) && path[i] == '/' {
                i++
            }
            i--
            str = ""
        } else {
            str += string(path[i])
        }
    }
    ans := ""
    for e := stk.Front(); e != nil; e = e.Next() {
        // ans += "/" + string(e.Value)     // need type assert
        ans += "/" + e.Value.(string)
    }
    if len(ans) == 0 {
        ans = "/"
    }
    return ans
}


// error ..  这里 不好处理 ..... 的情况，跨case了。  而且 不如 直接 split by /
// .同层,不需要任何操作,  ..上级,pop,    其他 下级, push....  但是.
// fallthrough  默认break
func simplifyPath2(path string) string {
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


func main_LT0071_20211022() {
// func main() {

    fmt.Println("ans:")

    // s := "/a/./b/../../c/"
    // s := "/asd///wdf/"
    // s := "/asd/////"
    // s := "/../"
    s := "/qwe///grg/.../asde/../sdfe/"

    ans := simplifyPath(s)

    fmt.Println(ans)

}

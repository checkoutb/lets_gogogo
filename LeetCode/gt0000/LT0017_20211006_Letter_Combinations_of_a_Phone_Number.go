// package sdq
package main

import (
    "fmt"
    // "strconv"
)


// digitToLetters := map[rune][]string{
//     '1': []string{},
//     '2': []string{"a", "b", "c"},
// ''


// queue

// 新切片保存 当前结果，然后swap


// 想起来，不知道可以defer 吗？



// *res = append(*res, curStr)
/// 。。。。。。。。。。。。。。




// Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
// Memory Usage: 2.2 MB, less than 38.99% of Go online submissions for Letter Combinations of a Phone Number.
// 。。这个怎么把切片传递进去 作为ans的容器？
func letterCombinations(digits string) []string {
    chs := []rune (digits)
    map2 := map[rune][]string{
        rune("2"[0]):[]string{"a","b","c"},
        rune("3"[0]):[]string{"d","e","f"},
        rune("4"[0]):[]string{"g","h","i"},
        rune("5"[0]):[]string{"j","k","l"},
        rune("6"[0]):[]string{"m","n","o"},
        rune("7"[0]):[]string{"p","q","r","s"},
        rune("8"[0]):[]string{"t","u","v"},
        rune("9"[0]):[]string{"w","x","y","z"}}

    // arr := [16*16]string{}
    // ans := &[]string{}
    // ans := arr[0 : 0]
    ans := []string{}
    if (len(chs) == 0) {
        return ans
    }
    ans = dfsa1("", chs, 0, map2, ans)
    return ans
}

func dfsa1(s string, chs []rune, idx int, map2 map[rune][]string, ans []string) []string {
    // fmt.Println(strconv.Itoa(idx) + ", " + strconv.Itoa(len(chs)))
    // fmt.Println(len(ans))
    // fmt.Println(ans)
    if (idx == len(chs)) {
        // t2 := append(ans, s)
        // fmt.Println(t2)
        // ans = &t2           // 感觉return 以后 就析构了。。     // 指向了不同的东西了。。。 。。 这个切片只能这样啊。
        // fmt.Println(len(*ans))
        // fmt.Println(ans)
        return []string{s}
    }
    ch := chs[idx]
    arr := map2[ch]
    ans2 := []string{}
    for i := 0; i < len(arr); i++ {
        t2 := dfsa1(s + arr[i], chs, idx + 1, map2, ans);
        ans2 = append(ans2, t2...)
    }
    return ans2
}


//func main_LT0017_20211006() {
func main() {

    fmt.Println("ans:")

    // s := "23"
    // s := ""
    s := "2"

    arr := letterCombinations(s)

    fmt.Println(len(arr))
    for _, v := range arr {
        fmt.Println(v)
    }

}

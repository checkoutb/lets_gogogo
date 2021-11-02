// package sdq
package main

import (
    "fmt"
)





// gg....



// tle...
// s1 = "great", s2 = "rgeat"
// 类似 divide conquer, 
// 那么 s1[0] 必须在 前半部分,  s1[last]  必须在 后半部分.   所以 上面的 划分 就是   gr    eat       gre at
//           前半部分 必须在 前半部分,  后半部分 必须在 后半部分......  非常富有哲理.....    就是 最大下标要 小于等于现在的, 那么 都在前面了.
// 可以swap的, 所以 不是 前半 必须在 前半,  而是 半部分 是一个整体.
// dfs算啊.  感觉可以dp的. 但...
func isScramble(s1 string, s2 string) bool {
    arr1 := [26]int{}
    arr2 := [26]int{}
    for i, ch := range s1 {
        ch2 := s2[i]
        t1, t2 := ch - 'a', ch2 - 'a'
        arr1[t1]++
        arr2[t2]++
    }
    for i := 0; i < 26; i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    ans := dfsa87(&s1, &s2, 0, len(s1) - 1, 0, len(s2) - 1)
    return ans
}

// [st, en]
// sz < 30 go
func dfsa87(s1, s2 *string, st1, en1 int, st2, en2 int) bool {
    // fmt.Println(st1, ", ", en1, " - ", st2, ", ", en2)
    if st1 > en1 {
        return true
    }
    if st1 == en1 {
        return (*s1)[st1] == (*s2)[st2] && st2 == en2
    }

    // 加上 arr11, arr12  本地感觉更慢..
    arr11 := [26]int{}
    arr12 := [26]int{}
    for i := st1; i <= en1; i++ {
        arr11[(*s1)[i] - 'a']++
        arr12[(*s2)[st2 + i - st1] - 'a']++
    }
    for i := 0; i < 26; i ++ {
        if arr11[i] != arr12[i] {
            // fmt.Println("asd")
            return false
        }
    }
    
    arr1 := [26]int{}
    arr2 := [26]int{}
    arr3 := [26]int{}
    for i := st1; i < en1; i++ {
        t2 := i - st1
        a1, a2, a3 := (*s1)[t2 + st1] - 'a', (*s2)[st2 + t2] - 'a', (*s2)[en2 - t2] - 'a'
        arr1[a1]++
        arr2[a2]++
        arr3[a3]++
        b2, b3 := true, true
        for j := 0; j < 26; j++ {
            if b2 && arr1[j] != arr2[j] {
                b2 = false
            }
            if b3 && arr1[j] != arr3[j] {
                b3 = false
            }
        }
        if b2 {
            b2 = dfsa87(s1, s2, st1, t2 + st1, st2, st2 + t2)
            if b2 {
                b2 = dfsa87(s1, s2, st1 + t2 + 1, en1, st2 + t2 + 1, en2)
                if b2 {
                    return true
                }
            }
        }
        if b3 {
            b3 = dfsa87(s1, s2, st1, st1 + t2, en2 - t2, en2)
            if b3 {
                b3 = dfsa87(s1, s2, st1 + t2 + 1, en1, st2, en2 - t2 - 1)
                if b3 {
                    return true
                }
            }
        }
    }
    return false
}

func main_LT0087_20211027() {
// func main() {

    fmt.Println("ans:")

    // s1, s2 := "great", "rgeat"
    // s1, s2 := "a", "a"
    // s1, s2 := "abcde", "caebd"
    s1, s2 := "eebaacbcbcadaaedceaaacadccd", "eadcaacabaddaceacbceaabeccd"
    

    ans := isScramble(s1, s2)

    fmt.Println(ans)

}

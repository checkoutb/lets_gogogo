// package sdq
package main

import (
    "fmt"
)


// sPattern, tPattern := map[uint8]int{}, map[uint8]int{}
// for index := range s {
//     if sPattern[s[index]] != tPattern[t[index]] {
//         return false
//     } else {
//         sPattern[s[index]] = index + 1
//         tPattern[t[index]] = index + 1
//     }
// }
// 比较这个值的上一次出现。。



// bool isIsomorphic(string s, string t) {
//     int m1[256] = {0}, m2[256] = {0}, n = s.size();
//     for (int i = 0; i < n; ++i) {
//         if (m1[s[i]] != m2[t[i]]) return false;
//         m1[s[i]] = i + 1;
//         m2[t[i]] = i + 1;
//     }
//     return true;
// }




// Runtime: 7 ms, faster than 27.01% of Go online submissions for Isomorphic Strings.
// Memory Usage: 2.6 MB, less than 56.32% of Go online submissions for Isomorphic Strings.
func isIsomorphic(s string, t string) bool {
    for i := 0; i < 2; i++ {
        map2 := map[byte]byte{}
        for i := 0; i < len(s); i++ {
            c1, c2 := s[i], t[i]
            ori1 := c1
            if _, ok := map2[c1]; ok {
                c1 = map2[c1]
            }
            if c1 != c2 {
                if _, ok := map2[ori1]; ok {
                    return false
                }
            }
            if c1 == ori1 {
                map2[c1] = c2
            }
        }
        s, t = t, s
    }
    return true
}

func isIsomorphic2(s string, t string) bool {
    map2 := map[byte]byte{}
    for i := 0; i < len(s); i++ {
        c1, c2 := s[i], t[i]
        ori1 := c1
        if _, ok := map2[c1]; ok {
            c1 = map2[c1]
        }
        if c1 != c2 {
            if _, ok := map2[ori1]; ok {
                return false
            }
        }
        if c1 == ori1 {
            map2[c1] = c2
        }
        // c1, c2 := s[i], t[i]
        // ori1 := c1
        // if _, ok := map2[ori1]; ok {
        //     c1 = map2[ori1]
        // }
        // if c1 != c2 {
        //     if _, ok := map2[ori1]; ok {
        //         return false
        //     }
        //     map2[ori1] = c2
        // }
    }
    // map3 := map[byte]bool{}
    // for _, v := range map2 {
    //     if _, ok := map3[v]; ok {
    //         return false
    //     } else {
    //         map3[v] = true
    //     }
    // }
    return true
}


func main_LT0205_20211127() {
// func main() {

    fmt.Println("ans:")


}

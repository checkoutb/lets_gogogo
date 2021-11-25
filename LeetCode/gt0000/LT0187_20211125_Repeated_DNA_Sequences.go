// package sdq
package main

import (
    "fmt"
)


// table := [1048576]byte{}
// for i := 0; i < 10; i++ {
//     hash = (hash << 2) | (int(s[i]-'A'+1) % 5)
// }
// table[hash] = 1
// for i := 10; i < len(s); i++ {
//     hash = ((hash << 2) ^ (int(s[i-10]-'A'+1)%5)<<20) | (int(s[i]-'A'+1) % 5)
//     if table[hash] == 0 {
//         table[hash] = 1
//     } else if table[hash] == 1 {
//         table[hash] = 2
//         ret = append(ret, s[i-9:i+1])
//     }
// }
// 4个数字，0,1,2,3  需要2bit，  所以<<2


// var numMap = map[byte]uint32{
// 	'A': 0,
// 	'C': 1,
// 	'G': 2,
// 	'T': 3,
// }
// var upperMask uint32 = 0xFFF00000
//
// sMap := map[uint32]bool{}
// res := []string{}
// var v uint32
// for i := 0; i < len(s); i++ {
//     num := numMap[s[i]]
//     v <<= 2
//     v |= num
//     v = v &^ upperMask
//     if i < 9 {
//         continue
//     }
//     if count, ok := sMap[v]; !ok {
//         sMap[v] = false
//     } else if !count {
//         sMap[v] = true
//         res = append(res, s[i-9:i+1])
//     }
// }



// Runtime: 12 ms, faster than 84.09% of Go online submissions for Repeated DNA Sequences.
// Memory Usage: 9.1 MB, less than 75.00% of Go online submissions for Repeated DNA Sequences.
// 这个例子有点。。
// Example2 可以看到 它可以 复用自己的一部分
// 但是 Example1 没有 AAACCCCCAA 作为结果啊。 。。。。 第一次是5个c， 第二次是6个c。。。。
// 应该是那个 hash的那个字符串搜索算法。   40yi  long        go.int/go.int64
func findRepeatedDnaSequences(s string) []string {
    sz1 := len(s)
    ans := []string{}
    if sz1 <= 10 {
        return ans
    }
    var hsh int64 = 0
    map2 := map[int64]int{}     // last char's index,,, 不需要。。。bool   false:exist,  true:got
    fun2 := func(ch byte) int64 {
        switch (ch) {
        case 'A':
            return 1
        case 'C':
            return 2
        case 'G':
            return 3
        case 'T':
            return 4
        }
        return 0
    }
    var mask int64 = 1000_0000_000
    for i := 0; i < 9; i++ {
        hsh *= 10
        hsh += fun2(s[i])
    }
    for i := 9; i < sz1; i++ {
        hsh *= 10
        hsh += fun2(s[i])
        hsh = hsh % mask
        if _, ok := map2[hsh]; ok {
            t2 := map2[hsh]
            if t2 != -1 {
                ans = append(ans, s[i - 9 : i + 1])
                map2[hsh] = -1
            }
        } else {
            map2[hsh] = i
        }
    }
    return ans
}


func main_LT0187_20211125() {
// func main() {

    fmt.Println("ans:")

    // s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
    s := "AAAAAAAAAAAAAAAAAAAAAAAAAA"

    ans := findRepeatedDnaSequences(s)

    fmt.Println(ans)

}

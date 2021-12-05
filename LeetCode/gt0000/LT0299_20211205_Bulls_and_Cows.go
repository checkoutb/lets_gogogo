// package sdq
package main

import (
    "fmt"
    "strconv"
)



// return fmt.Sprintf ("%dA%dB", bulls, cows)


// int[] numbers = new int[10];
// for (int i = 0; i<secret.length(); i++) {
//     if (secret.charAt(i) == guess.charAt(i)) bulls++;
//     else {
//         if (numbers[secret.charAt(i)-'0']++ < 0) cows++;
//         if (numbers[guess.charAt(i)-'0']-- > 0) cows++;
//     }
// }
// secret 的数字是 ++， guess的是--
// 所以当 你增加了一个secret 数字，且 count 是小于0 的，那就说明  guess前面 有这个数字， 所以 cows++


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Bulls and Cows.
// Memory Usage: 2.3 MB, less than 62.00% of Go online submissions for Bulls and Cows.
// 正确的值在正确的位置 A
// 存在的值在错误的位置 B
func getHint(secret string, guess string) string {
    map2 := map[byte]int{}
    map3 := map[byte]int{}
    cnta, cntb := 0, 0
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {
            cnta++
        } else {
            map2[secret[i]]++
            map3[guess[i]]++
        }
    }
    for k, v := range map2 {
        if _, ok := map3[k]; ok {
            if v > map3[k] {
                cntb += map3[k]
            } else {
                cntb += v
            }
        }
    }
    return strconv.Itoa(cnta) + "A" + strconv.Itoa(cntb) + "B"
}


func main_LT0299_20211205() {
// func main() {

    fmt.Println("ans:")


}

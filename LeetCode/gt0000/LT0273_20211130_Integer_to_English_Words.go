// package sdq
package main

import (
    "fmt"
    "strings"
)





// 0
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Integer to English Words.
// Memory Usage: 2.4 MB, less than 72.45% of Go online submissions for Integer to English Words.
// "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
// 1234567891
// Billion  Million  Thousand
// One, Two, Three Four
func numberToWords(num int) string {
    ans := []string{}

    if num >= 1_000_000_000 {
        ans = append(ans, converta273(num / 1_000_000_000)...)
        ans = append(ans, arr3[0])
        num %= 1_000_000_000
    }
    if num >= 1_000_000 {
        ans = append(ans, converta273(num / 1_000_000)...)
        ans = append(ans, arr3[1])
        num %= 1_000_000
    }
    if num >= 1_000 {
        ans = append(ans, converta273(num / 1_000)...)
        ans = append(ans, arr3[2])
        num %= 1_000
    }
    ans = append(ans, converta273(num)...)
    if len(ans) == 0 {
        ans = append(ans, arr[0])
    }

    return strings.Join(ans, " ")


    // arr := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", 
    //     "Ten", "Eleven", "Twelve", "Thirteen", "Forteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"};
    // arr2 := []string{"","","Twenty","Thirty","Forty","Fifty","Sixty","Seventy","Eighty","Ninety"};
    // arr3 := []string{"Billion", "Million", "Thousand"}
    // arr4 := []string{"Hundred"}

}

var (
    arr = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", 
    "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"};
    arr2 = []string{"","","Twenty","Thirty","Forty","Fifty","Sixty","Seventy","Eighty","Ninety"};
    arr3 = []string{"Billion", "Million", "Thousand"}
    arr4 = []string{"Hundred"}
)


// < 1000
func converta273(num int) []string {
    ans := []string{}
    if num / 100 > 0 {
        ans = append(ans, arr[num / 100])
        ans = append(ans, arr4[0])
        num %= 100
    }
    if num / 10 >= 2 {      // .. >= 20
        ans = append(ans, arr2[num / 10])
        num %= 10
    //     if num > 0 {
    //         ans = append(ans, arr[num])
    //     }
    // } else {
    //     ans = append(ans, )
    }
    if num > 0 {
        ans = append(ans, arr[num])
    }
    return ans
}

func main_LT0273_20211130() {
// func main() {

    fmt.Println("ans:")


}

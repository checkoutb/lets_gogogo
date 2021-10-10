// package sdq
package main

import (
    "fmt"
)



// int_vec := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1};
// rom_vec := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"};
// .....


// String M[] = {"", "M", "MM", "MMM"};
// String C[] = {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"};
// String X[] = {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"};
// String I[] = {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"};
// return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10];



// Runtime: 16 ms, faster than 25.00% of Go online submissions for Integer to Roman.
// Memory Usage: 3.7 MB, less than 36.76% of Go online submissions for Integer to Roman.
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// I can be placed before V (5) and X (10) to make 4 and 9. 
// X can be placed before L (50) and C (100) to make 40 and 90. 
// C can be placed before D (500) and M (1000) to make 400 and 900.
// 4-2 3-2 
// 6-4 5-4
// 2-0 1-0
func intToRoman(num int) string {
    arr1 := []int{1,5,10,50,100,500,1000}       // 0-6
    arr2 := []string{"I", "V", "X", "L", "C", "D", "M"}
    ans := ""
    for i := 6; i >= 0; i-- {
        for num >= arr1[i] {
            ans += arr2[i]
            num -= arr1[i]
        }
        if i > 0 && (num / arr1[(i - 1) / 2 * 2]) % 10 == (4 + 5 * ((i + 1) % 2)) {
            ans += arr2[(i - 1) / 2 * 2] + arr2[i]
            num -= (arr1[i] - arr1[(i - 1) / 2 * 2])
        }
    }
    return ans
}


func main_LT0012_20211006() {
// func main() {

    fmt.Println("ans:")

    // a := 3
    // a := 4
    // a := 9
    // a := 58
    a := 1994

    fmt.Println(intToRoman(a))

}

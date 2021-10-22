// package sdq
package main

import (
    "fmt"
    "strconv"
)


// var aInt big.Int
// var bInt big.Int
// aInt.SetString(a, 2)
// bInt.SetString(b, 2)
// var answer big.Int
// var carry big.Int
// var zero big.Int
// for bInt.Cmp(&zero) != 0 {
//     answer.Xor(&aInt, &bInt)
//     carry.And(&aInt, &bInt).Lsh(&carry, 1)
//     aInt.Set(&answer)
//     bInt.Set(&carry)
// }
// return aInt.Text(2)


// sum := x1+x2+carry
// carry = sum/2
// result = append(result, sum%2+'0')


// sum = (firstDigit ^ secondDigit ^ carry);
// carry = ((firstDigit | carry) & (secondDigit | carry) & (firstDigit | secondDigit));
// sb.append(sum);


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
// Memory Usage: 2.9 MB, less than 23.03% of Go online submissions for Add Binary.
func addBinary(a string, b string) string {
    i1, i2, carry := len(a) - 1, len(b) - 1, 0
    ans := ""
    for i1 >= 0 || i2 >= 0 || carry > 0 {
        t1, t2 := 0, 0
        if i1 >= 0 {
            t1 = int(a[i1] - '0')
            i1--
        }
        if i2 >= 0 {
            t2 = int(b[i2] - '0')
            i2--
        }
        t2 += t1 + carry
        carry = 0
        if t2 > 1 {
            t2 %= 2
            carry = 1
        }
        ans = strconv.Itoa(t2) + ans
    }
    return ans
}


func main_LT0067_20211022() {
// func main() {

    fmt.Println("ans:")


}

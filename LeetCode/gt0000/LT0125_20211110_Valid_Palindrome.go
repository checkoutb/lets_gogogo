// package sdq
package main

import (
    "fmt"
)



// runes := []rune{}
// for _, char := range s {
//     if unicode.IsLetter(char) || unicode.IsDigit(char) {
//         runes = append(runes, unicode.ToLower(char))
//     }
// }


// switch {
//     case s[r] >= 'a' && s[r] <= 'z': break
//     case s[r] >= 'A' && s[r] <= 'Z': break
//     case s[r] >= '0' && s[r] <= '9': break
//     default: r--; continue           // 。。。 continue。。。。
// }


// String actual = s.replaceAll("[^A-Za-z0-9]", "").toLowerCase();
// String rev = new StringBuffer(actual).reverse().toString();
// return actual.equals(rev);


// for (int i = 0, j = s.size() - 1; i < j; i++, j--) { // Move 2 pointers from each end until they collide
//     while (isalnum(s[i]) == false && i < j) i++; // Increment left pointer if not alphanumeric
//     while (isalnum(s[j]) == false && i < j) j--; // Decrement right pointer if no alphanumeric
//     if (toupper(s[i]) != toupper(s[j])) return false; // Exit and return error if not match
// }


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Palindrome.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Valid Palindrome.
// alpha number
func isPalindrome(s string) bool {
    l, r := 0, len(s) - 1
    for l < r {
        for l < r && !isAlphanum125(s[l]) {
            l++
        }
        for l < r && !isAlphanum125(s[r]) {
            r--
        }
        if l < r {
            if convert125(s[l]) != convert125(s[r]) {
                return false
            }
            l++
            r--
        }
    }
    return true
}

func isAlphanum125(ch byte) bool {
    return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func convert125(ch byte) byte {
    if ch >= 'A' && ch <= 'Z' {
        return ch + 'a' - 'A'
    }
    return ch
}

func main_LT0125_20211110() {
// func main() {

    fmt.Println("ans:")


}

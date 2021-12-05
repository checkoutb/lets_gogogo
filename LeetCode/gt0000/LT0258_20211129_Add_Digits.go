// package sdq
package main

import (
    "fmt"
)


// if(num == 0) return 0;
// if(num % 9 == 0) return 9;
// return num%9;

// return 1 + (num - 1) % 9;


// dr(n) = 0 if n == 0
// dr(n) = (b-1) if n != 0 and n % (b-1) == 0
// dr(n) = n mod (b-1) if n % (b-1) != 0
// or
// dr(n) = 1 + (n - 1) % 9

// ~input: 0 1 2 3 4 ...
// output: 0 1 2 3 4 5 6 7 8 9 1 2 3 4 5 6 7 8 9 1 2 3 ....
// ..后面没有0


// Runtime: 4 ms, faster than 33.87% of Go online submissions for Add Digits.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Add Digits.
// buzhidao
func addDigits(num int) int {
    for num > 9 {
        t2 := 0
        for num > 0 {
            t2 += num % 10
            num /= 10
        }
        num = t2
    }
    return num
}


func main_LT0258_20211129() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
    "fmt"
)



// while (n > 1)
// {
//     if ((n & 1) == 0)
//         n >>= 1;
//     else if (n == 3)
//         return ans + 2;
//     else if ((n & 3) == 3)
//         ++n;
//     else
//         --n;
//     ++ans;
// }
// 3 代表 最后2位 是 11, 此时 确实 + 1 ， 最后2位变成 00 更快。
// x11 y00 y0 y
// x11 x10 x1 x0 x


// public int integerReplacement(int n) {
//     if (n == 1) return 0;
//     if (n % 2 == 0) 
//         return (1 + integerReplacement(n / 2));
//     else 
//         return (2 + Math.min(integerReplacement(n/2), integerReplacement(n/2 + 1)));
// }
// 规避 INT_MAX



// while (n > 1){
//     if (n % 2 == 0) n  /= 2;
//     else{
//         if ( (n + 1) % 4 == 0 && (n - 1 != 2) ) n++;
//         else n--;
//     }
//     count++;
// }
// When n is odd it can be written into the form n = 2k+1 (k is a non-negative integer.). 
// That is, n+1 = 2k+2 and n-1 = 2k. Then, (n+1)/2 = k+1 and (n-1)/2 = k. 
// So one of (n+1)/2 and (n-1)/2 is even, the other is odd. 
// And the "best" case of this problem is to divide as much as possible. 
// Because of that, always pick n+1 or n-1 based on if it can be divided by 4. 
// The only special case of that is when n=3 you would like to pick n-1 rather than n+1.



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Integer Replacement.
// Memory Usage: 1.9 MB, less than 79.31% of Go online submissions for Integer Replacement.
func integerReplacement(n int) int {
    return dfsa0397(n)
}

func dfsa0397(n int) int {
    cnt := 0
    if n & (n - 1) == 0 {
        for n > 1 {
            cnt++
            n >>= 1
        }
        return cnt
    }
    if n % 2 == 0 {
        return dfsa0397(n / 2) + 1
    } else {
        t2 := n - 1
        if t2 & (t2 - 1) == 0 {
            return dfsa0397(n - 1) + 1
        }
        t2 = n + 1
        if t2 & (t2 - 1) == 0 {
            return dfsa0397(n + 1) + 1
        }
        c1 := dfsa0397(n - 1)
        c2 := dfsa0397(n + 1)
        if c1 > c2 {
            return c2 + 1
        } else {
            return c1 + 1
        }
    }
} 

// 是不是 奇数始终 -1 ？ 好像不是。    65535
// +1 是 2的幂 就 +1 ，否则 -1 ？ ....  3
// ...... 10000
// dfs了。 +1 -1 。。。 能匹配2^n 直接返回 n + 1
func integerReplacement2(n int) int {
    cnt := 0
    for n != 1 {
        if n % 2 == 1 {
            t2 := n - 1
            if t2 & (t2 - 1) == 0 {
                n--
            } else {
                t2 = n + 1             // 2^31 - 1
                if t2 & (t2 - 1) == 0 {
                    n++
                } else {
                    n--
                }
            }
        } else {
            n /= 2
        }
        cnt++
    }
    return cnt
}

// func main_LT0397_20220125() {
func main() {

    fmt.Println("ans:")

    n := 65535      // 17
    // n := 3
    // n := 10000          // 16

    fmt.Println(integerReplacement(n))

}

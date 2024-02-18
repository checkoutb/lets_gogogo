// package sdq
package main

import (
    "fmt"
    "sort"
)

// D D

// var pal = make([]int, 0, 109999)

// func init() {
//     for base:=1; base <= 10000; base*=10 {

// golang中有个神奇的函数init,该函数会在所有程序执行开始前被调用,
// 每个包可以包含多个init函数,所有被编辑器识别到的init函数都会在main函数执行前被调用。




// palindrome1 := median
// for !isPalindrome(palindrome1) {
//     palindrome1--
// }

// ...






// We have to take median first and then check for the nearest possible palindrom of the median
// To find the possible palindrom, nearest to the median there are five cases need to check :
//     case 1: one digit less nearest palindrom
//     case 2: one digit more nearest palindrom
//     case 3: simply take the mirror image of the left half
//     case 4: take mirror image of the (left half + 1)
//     case 5: take mirror image of the (left half - 1)






// Runtime101 ms
// Beats
// 36.92%
// Memory7.5 MB
// Beats
// 98.46%


// 如何求 小于 x 的最大的 回文数字。 ok，想起来了。  12344321, 只需要 遍历 前半部分， 就是 1234, 然后 就可以计算 回文：12344321,1234321
// 但是 1234321 还是 12344321 ， 要2次 BS？ 一次 1234321, 一次 12344321， 应该是的。
// 抛物线
func minimumCost(nums []int) int64 {
    sort.Ints(nums)

    min2 := func(a, b int64) int64 {
        if a < b {
            return a;
        }
        return b;
    }

    st, en := 1, 9999;
    var ans int64 = calAbsSum(nums, 10_0000_0001);          // not 1000000000 ...
    for st <= en {
        
        md := (st + en) / 2;
        pd1 := mkPalind(md - 1, true);
        t1 := calAbsSum(nums, pd1);

        pd2 := mkPalind(md, true);
        t2 := calAbsSum(nums, pd2);

        // fmt.Println(md, pd1, pd2, t1, t2);

        if t1 == t2 {
            ans = min2(ans, t1);
            break;
        } else if t1 > t2 {
            ans = min2(ans, t2);
            st = md + 1;
        } else {
            ans = min2(ans, t1);
            en = md - 2;
        }
    }

    st, en = 1, 99999;
    for st <= en {
        md := (st + en) / 2;
        pd1 := mkPalind(md - 1, false);
        pd2 := mkPalind(md, false);
        t1 := calAbsSum(nums, pd1);
        t2 := calAbsSum(nums, pd2);

        // fmt.Println(md, pd1, pd2, t1, t2);

        if t1 == t2 {
            ans = min2(ans, t1)
            break;
        } else if t1 > t2 {
            ans = min2(ans, t2);
            st = md + 1;
        } else {
            ans = min2(ans, t1);
            en = md - 2;
        }
    }

    return ans;
}

func calAbsSum(nums []int, palind int) int64 {
    var ans int64 = 0

    // prefix sum + BS is faster

    abs2 := func(a int) int {
        if a < 0 {
            a = -a;
        }
        return a;
    }

    for _, n := range nums {
        ans += int64(abs2(n - palind));
    }
    return ans
}

func mkPalind(pfx int, cp bool) int {
    ans := pfx
    if cp {     // 1234 4321
        ;
    } else {    // 1234 321
        pfx /= 10;
    }
    for pfx > 0 {
        ans *= 10;
        ans += (pfx % 10);
        pfx /= 10
    }
    return ans
}


func main_LT2967_20240108() {
// func main() {

    // vi := []int{1,2,3,4,5};
    // vi := []int{10,12,13,14,15};
    vi := []int{1,2};
   
    fmt.Println("ans:", minimumCost(vi))


}

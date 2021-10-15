// package sdq
package main

import (
    "fmt"
)



// StefanPochmann：
// I learned a lot by reading tutorials/docs and by looking at solutions of others at checkio.org (a Python programming site) 
// and by spending a month on Python topics at Stackoverflow. Also, practice practice practice :-)


// func generateParenthesis(n int) []string {
//     var result []string
//     curr := make([]byte, 2 * n)
//     helper(n, 0, 0, curr, &result)
//     return result
// }
// func helper(left, right, idx int, curr []byte, result *[]string) {
//     if left == 0 && right == 0 {
//         *result = append(*result, string(curr))
//     } else {
//         if left > 0 {
//             curr[idx] = '('
//             helper(left - 1, right + 1, idx + 1, curr, result)
//         }
//         if right > 0 {
//             curr[idx] = ')'
//             helper(left, right - 1, idx + 1, curr, result)
//         }
//     }
// }
// 这里的byte 更像 char。。。
// *result = append(*result, string(curr))



// if(open < max)
// backtrack(list, str+"(", open+1, close, max);
// if(close < open)
// backtrack(list, str+")", open, close+1, max);




// Runtime: 7 ms, faster than 10.19% of Go online submissions for Generate Parentheses.
// Memory Usage: 4 MB, less than 7.56% of Go online submissions for Generate Parentheses.
func generateParenthesis(n int) []string {
    ans := []string{}
    ans = dfsa2(n * 2, 0, ans, "")              // 这里vs code 识别了 LT0017的那个 dfsa1。。
    return ans
}

// 。。。 离大谱。。 方法签名都不一样。。是vs code的还是go的？
// dfsa1 redeclared in this block
// 	/home/ubuntu20/Documents/git/lets_gogogo/LeetCode/gt0000/LT0017_20211006_Letter_Combinations_of_a_Phone_Number.go:57:83: previous declaration
// diff == count('(') - count(')')
func dfsa2(mxsz int, diff int, ans []string, s string) []string {
    if len(s) == mxsz {
        // ans = append(a)             // 这个。。ans会直接替换，所以不能用。试试指针
        // ans =            // 指针也是，除非指针的指针。。  直接return算了。  估计得用list之类的。
        return []string{s};
    }
    var ans2 []string = []string{}
    if (len(s) + diff == mxsz) {
        s2 := s + ")"
        res := dfsa2(mxsz, diff - 1, ans, s2)           // 并不需要，因为后续全是))))) 了。。
        ans2 = append(ans2, res...)
        // for _, val := range res {
        //     ans2 = append(ans2, val)
        // }
    } else {
        s2 := s + "("
        res := dfsa2(mxsz, diff + 1, ans, s2)
        if diff > 0 {
            s2 = s + ")"            // 直接写 形参 的位置上。。
            res = append(res, dfsa2(mxsz, diff - 1, ans, s2)...)
        }
        ans2 = append(ans2, res...)         // 不需要res，，直接 dfsa1()...

// 如果把 diff的判断 放到方法开始，那么这里能一行， 在dfs开始的地方，判断 如果 diff<0  return []string{}
// ans2 = append(ans2, append(dfsa1(mxsz, diff + 1, ans, s + "(")...), dfsa1(mxsz, diff - 1, ans, s + ")"))

    }
    return ans2
}


func main_LT0022_20211013() {
// func main() {

    fmt.Println("ans:")

    n := 1

    ans := generateParenthesis(n)

    for _, s := range ans {
        fmt.Println(s)
    }

}

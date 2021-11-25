// package sdq
package main

import (
    "fmt"
    "strings"
    "strconv"
)



// for(int i = 0; i < len; i++){
//     char cur = expression.charAt(i);
//     if(cur == '+' || cur == '-' || cur == '*'){
//         List<Integer> left = diffWaysToCompute(expression.substring(0, i));
//         List<Integer> right = diffWaysToCompute(expression.substring(i+1));
//         for(int k : left){
//             for(int j : right){
// 本次是 按 i 切分。


// for i := st; i < end; i++ {
//     tm1, tm2 := solve(st, i, nums, operator), solve(i+1, end, nums, operator)
//     for _, left := range tm1 {
//         for _, right := range tm2 {
//             ans = append(ans, calc(left, right, operator[i]))
//         }
//     }
// }


// divide-and-conquer


// vector<int> diffWaysToCompute(string input) {
//     vector<int> output;
//     for (int i=0; i<input.size(); i++) {
//         char c = input[i];
//         if (ispunct(c))
//             for (int a : diffWaysToCompute(input.substr(0, i)))
//                 for (int b : diffWaysToCompute(input.substr(i+1)))
//                     output.push_back(c=='+' ? a+b : c=='-' ? a-b : a*b);
//     }
//     return output.size() ? output : vector<int>{stoi(input)};
// }

// def diffWaysToCompute(self, input):
//     return [eval(`a`+c+`b`)
//             for i, c in enumerate(input) if c in '+-*'
//             for a in self.diffWaysToCompute(input[:i])
//             for b in self.diffWaysToCompute(input[i+1:])] or [int(input)]

// def diffWaysToCompute(self, input):
// return [a+b if c == '+' else a-b if c == '-' else a*b
//         for i, c in enumerate(input) if c in '+-*'
//         for a in self.diffWaysToCompute(input[:i])
//         for b in self.diffWaysToCompute(input[i+1:])] or [int(input)]





// Runtime: 0 ms, faster than 100.00% of Go online submissions for Different Ways to Add Parentheses.
// Memory Usage: 2.3 MB, less than 44.44% of Go online submissions for Different Ways to Add Parentheses.
func diffWaysToCompute(expression string) []int {
    isDig := func (ch rune) bool {
        return ch <= '9' && ch >= '0'
    }
    notDig := func (ch rune) bool {
        return !isDig(ch)
    }
    nums2 := strings.FieldsFunc(expression, notDig)
    ops := strings.FieldsFunc(expression, isDig)
    sz1 := len(nums2)
    nums := make([]int, sz1)
    for i, _ := range nums {
        t5 , _ := strconv.Atoi(nums2[i])
        nums[i] = t5
    }
    arr := make([][][]int, sz1)
    for i, _ := range arr {
        arr[i] = make([][]int, sz1)
        for j, _ := range arr[i] {
            arr[i][j] = []int{}
        }
    }

    funop := func(op byte, a, b int) int {
        switch op {
        case '+':
            return a + b
        case '-':
            return a - b
        case '*':
            return a * b
        }
        return -1
    }

    for a := 0; a < sz1; a++ {
        if a == 0 {
            for i := 0; i < sz1; i++ {
                arr[i][i] = append(arr[i][i], nums[i])
            }
        } else {
            for i := 0; i + a < sz1; i++ {
                j := i + a
                for k := i; k < j; k++ {            // 在 k 后面切分
                    map2 := arr[i][k]           // arr...
                    map3 := arr[k + 1][j]
                    for _, v := range map2 {
                        for _, v2 := range map3 {
                            t2 := funop(ops[k][0], v, v2)
                            // arr[i][j][t2] = true
                            arr[i][j] = append(arr[i][j], t2)
                        }
                    }
                }
            }
        }
    }
    return arr[0][sz1 - 1]
}





// wrong ...  needn't distinct

// 只能dfs遍历吧？ 不， 还可以 dp。但是。。或许行，但是 得返回 一个集合， 因为 括号里面 也有很多种可能。
// 2-3-4-5
// 2-3
// 2-3-4  (2-3)-4  2-(3-4)
// 2-3-4-5  (2-3)-(4-5) ((2-3)-4)-5  2-(3-4)-5 
//      23-45 2-34-5 
// 应该是 长度
// 1-2-3-4-5-6
// 12,23,34,45,56   .. 遍历这个， 前后加
// 12+3,1+23,23+4,2+34,3+45,34+5
func diffWaysToCompute2(expression string) []int {
    isDig := func (ch rune) bool {
        return ch <= '9' && ch >= '0'
    }
    notDig := func (ch rune) bool {
        return !isDig(ch)
    }
    nums2 := strings.FieldsFunc(expression, notDig)
    ops := strings.FieldsFunc(expression, isDig)

    // fmt.Println(nums2)
    // fmt.Println(ops)

    sz1 := len(nums2)
    nums := make([]int, sz1)
    for i, _ := range nums {
        t5 , _ := strconv.Atoi(nums2[i])
        nums[i] = t5
    }

    arr := make([][]map[int]bool, sz1)
    for i, _ := range arr {
        arr[i] = make([]map[int]bool, sz1)
        for j, _ := range arr[i] {
            arr[i][j] = map[int]bool{}
        }
    }

    funop := func(op byte, a, b int) int {
        switch op {
        case '+':
            return a + b
        case '-':
            return a - b
        case '*':
            return a * b
        }
        return -1
    }

    for a := 0; a < sz1; a++ {
        if a == 0 {
            for i := 0; i < sz1; i++ {
                arr[i][i][nums[i]] = true
            }
        } else {
            for i := 0; i + a < sz1; i++ {
                j := i + a
                for k := i; k < j; k++ {            // 在 k 后面切分
                    map2 := arr[i][k]
                    map3 := arr[k + 1][j]
                    for v, _ := range map2 {
                        for v2, _ := range map3 {       // ... k,v..
                            t2 := funop(ops[k][0], v, v2)
                            arr[i][j][t2] = true
                        }
                    }
                }
            }
        }
    }
    // // fmt.Println(arr)
    // for i := 0; i < sz1; i++ {
    //     for j := i; j < sz1; j++ {
    //         fmt.Print(i, "+", j, ": ")
    //         for k, _ := range arr[i][j] {
    //             fmt.Print(k, ", ")
    //         }
    //         fmt.Println()
    //     }
    // }
    ans := []int{}
    for k, _ := range arr[0][sz1 - 1] {
        ans = append(ans, k)
    }
    return ans

    // arr := make([][]int, sz1)
    // for i, _ := range arr {
    //     arr[i] = make([]int, sz1)
    // }

    // for i = 1; i <= sz1; i++ {          // length
    //     if i == 1 {
    //         for a := 0; a < sz1; a++ {
    //             arr[a][a] = 
    //         }
    //     } else {

    //     }
    // }
}



func main_LT0241_20211125() {
// func main() {

    fmt.Println("ans:")

    // s := "2-1-1"
    s := "2*3-4*5"

    ans := diffWaysToCompute(s)

    fmt.Println(ans)

}

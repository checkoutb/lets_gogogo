// package sdq
package main

import (
    "fmt"
)


// for(string i :tokens){
//     if(i!="+" && i!="-" && i!="*" && i!="/" ){
//         st.push(stoi( i ));
//     }
//     else{
//         int second =st.top();
//         st.pop();
//         int first=st.top();
//         st.pop();
//         if(i=="+")
//            st.push(first+second);
//         else if(i=="-")
//             st.push(first-second);
//         else if(i == "*")
//             st.push(first*second);
//         else
//             st.push(first/second);
//     }
// }


// for _, token := range tokens {
//     var lop, rop int
//     switch token {
//         case "+":
//         rop, lop, stack = stack[len(stack)-1], stack[len(stack)-2], stack[:len(stack)-2]
//         stack = append(stack, lop+rop)
//         case "-":
//         rop, lop, stack = stack[len(stack)-1], stack[len(stack)-2], stack[:len(stack)-2]
//         stack = append(stack, lop-rop)
//         case "*":
//         rop, lop, stack = stack[len(stack)-1], stack[len(stack)-2], stack[:len(stack)-2]
//         stack = append(stack, lop*rop)
//         case "/":
//         rop, lop, stack = stack[len(stack)-1], stack[len(stack)-2], stack[:len(stack)-2]
//         stack = append(stack, lop/rop)
//         default:
//         val, _ := strconv.Atoi(token)
//         stack = append(stack, val )
//     }
// }


// unordered_map<string, function<int (int, int) > > map = {
//     { "+" , [] (int a, int b) { return a + b; } },
//     { "-" , [] (int a, int b) { return a - b; } },
//     { "*" , [] (int a, int b) { return a * b; } },
//     { "/" , [] (int a, int b) { return a / b; } }
// };
// std::stack<int> stack;
// for (string& s : tokens) {
//     if (!map.count(s)) {
//         stack.push(stoi(s));
//     } else {
//         int op1 = stack.top();
//         stack.pop();
//         int op2 = stack.top();
//         stack.pop();
//         stack.push(map[s](op2, op1));
//     }
// }
// return stack.top();



// Runtime: 4 ms, faster than 88.18% of Go online submissions for Evaluate Reverse Polish Notation.
// Memory Usage: 4.4 MB, less than 40.00% of Go online submissions for Evaluate Reverse Polish Notation.
// 后置
func evalRPN(tokens []string) int {
    map2 := map[string]func(int, int) int {
        "+" : func(a, b int) int {
            return a + b
        },
        "-" : func(a, b int) int {
            return a - b
        },
        "*" : func(a, b int) int {
            return a * b
        },
        "/" : func(a, b int) int {
            return a / b
        },
    }
    conv := func(s string) (num int, err error) {
        if len(s) == 1 && (s[0] < '0' || s[0] > '9') {
            err = errors.New("NaN")
            return
        } else {
            fun2 := func(s2 string) int {
                result := 0
                for _, ch := range s2 {
                    result *= 10
                    result += (int(ch) - '0')
                }
                return result
            }
            rr := 0
            if s[0] == '+' || s[0] == '-' {
                rr = fun2(s[1 : ])
                if s[0] == '-' {
                    rr *= -1
                }
            } else {
                rr = fun2(s)
            }
            num = rr
            return
        }
    }
    stk := []int{}
    for _, str := range tokens {
        num, err := conv(str)
        if err != nil {
            a, b := stk[len(stk) - 2], stk[len(stk) - 1]
            t2 := map2[str](a, b)
            stk = stk[0 : len(stk) - 1]
            stk[len(stk) - 1] = t2
        } else {
            stk = append(stk, num)
        }
    }
    return stk[0]
}


func main_LT0150_20211122() {
// func main() {

    fmt.Println("ans:")


}

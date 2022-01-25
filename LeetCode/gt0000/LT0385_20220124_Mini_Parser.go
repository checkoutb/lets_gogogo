// package sdq
package main

import (
    "fmt"
)







// func bracket(s string, i *int) *NestedInteger {
// 	ret := &NestedInteger{}
// 	out:
// 	for *i < len(s) {
// 		c := s[*i]
// 		*i++
// 		switch c {
// 		case '[':
// 			ret.Add(*bracket(s, i))
// 		case ']':
// 			break out
// 		case ',':
// 			continue
// 		default:
// 			ret.Add(*number(s, *i - 1, i))
// 		}
// 	}
// 	return ret
// }
//
// func number(s string, start int, i *int) *NestedInteger {
// 	for *i < len(s) && s[*i] >= '0' && s[*i] <= '9' {
// 		*i++
// 	}
// 	n, _ := strconv.ParseInt(s[start:*i], 10, 0)
// 	ret := &NestedInteger{}
// 	ret.SetInteger(int(n))
// 	return ret
// }



// stack<NestedInteger>st;
// st.push(NestedInteger());
// for(int i = 0;i<s.size();i++){
//     if(s[i] == '['){
//         st.push(NestedInteger());
//     }
//     else if(s[i] == ']'){
//         auto it = st.top();
//         st.pop();
//         st.top().add(it);
//     }
//     else if(s[i] == ',')
//         continue;
//     else{
//         int k = i;
//         while(k+1<s.size() && isdigit(s[k+1])) k++;
//         st.top().add(NestedInteger(stoi(s.substr(i, k-i+1))));
//         i = k;
//     }
// }




// NestedInteger parse(const string& s) {
//     if (s[p] == '[') {
//         ++p;
//         NestedInteger root;
//         while (s[p] != ']') {
//             root.add(parse(s));
//             if (s[p] == ',') {
//                 ++p;
//             }
//         }
//         ++p;
//         return root;
//     }
//
//     bool negative = false;
//     if (s[p] == '-') {
//         ++p;
//         negative = true;
//     }
//
//     int value = 0;
//     while (p < s.size() && s[p] >= '0' && s[p] <= '9') {
//         value = value * 10 + s[p] - '0';
//         ++p;
//     }
//     if (negative) {
//         value = 0 - value;
//     }
//     return NestedInteger{value};
// }





// Runtime: 7 ms, faster than 14.29% of Go online submissions for Mini Parser.
// Memory Usage: 5.3 MB, less than 28.57% of Go online submissions for Mini Parser.
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
//
// [ 代表list 的开始。
func deserialize(s string) *NestedInteger {
    idx := 0
    ans := dfsa385(s, &idx)
    return ans
}

func dfsa385(s string, idx *int) *NestedInteger {
    if *idx >= len(s) {
        return nil
    }
    if s[*idx] == ',' {
        *idx++
    }
    ans := &NestedInteger{}
    if s[*idx] == '[' {
        *idx++
        for s[*idx] != ']' {
            ans.Add(*dfsa385(s, idx))
        }
        *idx++
    } else {
        ans.SetInteger(toInta385(s, idx))
    }
    return ans
}

func toInta385(s string, idx *int) int {
    ans := 0
    isN := false
    if s[*idx] == '-' {
        isN = true
        *idx++
    }
    for *idx < len(s) && s[*idx] != ']' && s[*idx] != ',' {
        ans *= 10
        ans += (int(s[*idx]) - int('0'))
        *idx++
    }
    if isN {
        ans *= -1
    }
    return ans
}


func main_LT0385_20220124() {
// func main() {

    fmt.Println("ans:")


}

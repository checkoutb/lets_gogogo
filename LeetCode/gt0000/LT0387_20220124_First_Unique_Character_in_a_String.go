// package sdq
package main

import (
    "fmt"
)



// Map<Character, Integer> map = new LinkedHashMap<>();
// Set<Character> set = new HashSet<>();
// for (int i = 0; i < s.length(); i++) {
//     if (set.contains(s.charAt(i))) {
//         if (map.get(s.charAt(i)) != null) {                  // ..这个不需要判断，直接 remove 也可以
//             map.remove(s.charAt(i));
//         }
//     } else {
//         map.put(s.charAt(i), i);
//         set.add(s.charAt(i));
//     }
// }
// return map.size() == 0 ? -1 : map.entrySet().iterator().next().getValue();


// var firstUniqChar = function(s) {
//     for(i=0;i<s.length;i++){
//         if (s.indexOf(s[i])===s.lastIndexOf(s[i])){
//            return i;
//        } 
//     }
//     return -1;
//  };
// 时间O(n^2)


// Runtime: 15 ms, faster than 72.11% of Go online submissions for First Unique Character in a String.
// Memory Usage: 5.6 MB, less than 50.09% of Go online submissions for First Unique Character in a String.
func firstUniqChar(s string) int {
    arr1, arr2 := make([]int, 123), make([]int, 123)            // count, first index
    for i := 97; i < 123; i++ {
        arr2[i] = -1
    }
    for i, v := range s {
        arr1[v]++
        if arr2[v] == -1 {
            arr2[v] = i
        }
    }
    ans := len(s)
    for i := 97; i < 123; i++ {
        if arr1[i] == 1 {
            if arr2[i] < ans {
                ans = arr2[i]
            }
        }
    }
    if ans == len(s) {
        return -1
    }
    return ans
}


func main_LT0387_20220124() {
// func main() {

    fmt.Println("ans:")


}

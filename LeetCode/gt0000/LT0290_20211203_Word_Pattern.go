// package sdq
package main

import (
    "fmt"
    "strings"
)


// if (map.containsValue(arr[i]))
// "map.containsValue(arr[i])" is O(n).


// for (Integer i=0; i<words.length; ++i)
// if (index.put(pattern.charAt(i), i) != index.put(words[i], i))
//     return false;


// istringstream in(str);
// int i = 0, n = pattern.size();
// for (string word; in >> word; ++i) {
//     if (i == n || p2i[pattern[i]] != w2i[word])
//         return false;
//     p2i[pattern[i]] = w2i[word] = i + 1;
// }



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Pattern.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Word Pattern.
func wordPattern(pattern string, s string) bool {
    arr := strings.Fields(s)
    if len(pattern) != len(arr) { return false }
    map2 := map[byte]string{}
    map3 := map[string]byte{}
    for i := 0; i < len(arr); i++ {
        t2 := pattern[i]
        t3 := arr[i]
        if _, ok := map2[t2]; ok {
            if map2[t2] != t3 {
                return false
            }
        }
        if _, ok := map3[t3]; ok {
            if map3[t3] != t2 { return false }
        }
        map2[t2] = t3
        map3[t3] = t2
    }
    return true
}


func main_LT0290_20211203() {
// func main() {

    fmt.Println("ans:")


}

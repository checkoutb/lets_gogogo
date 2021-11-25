// package sdq
package main

import (
    "fmt"
)



// func getAnagramKey(s string) string {
//     anagramKeyMap := [26]int{}
//     for _, char := range s {
//         anagramKeyMap[int(char)-97] = anagramKeyMap[int(char)-97] + 1
//     }
//     return strings.Join(strings.Fields(fmt.Sprint(anagramKeyMap)), "")
// }


// bool isAnagram(string s, string t) { 
//     sort(s.begin(), s.end());
//     sort(t.begin(), t.end());
//     return s == t; 
// }


// int[] alphabet = new int[26];
// for (int i = 0; i < s.length(); i++) alphabet[s.charAt(i) - 'a']++;
// for (int i = 0; i < t.length(); i++) {
//   alphabet[t.charAt(i) - 'a']--;
//   if(alphabet[t.charAt(i) - 'a'] < 0) return false;
// }
// for (int i : alphabet) if (i != 0) return false;
// return true;


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
// Memory Usage: 2.9 MB, less than 50.27% of Go online submissions for Valid Anagram.
func isAnagram(s string, t string) bool {
    sz1, sz2 := len(s), len(t)
    if sz1 != sz2 { return false }
    arr := [26]int{}
    for i := 0; i < sz1; i++ {
        arr[s[i] - 'a']++
        arr[t[i] - 'a']--
    }
    for _, v := range arr {
        if v != 0 {
            return false
        }
    }
    return true
}

func main_LT0242_20211125() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
    "fmt"
)


// while (i < s.length() && s[i] != ']') {
//     if (!isdigit(s[i])) {
//         res += s[i++];
//     } else {
//         int n=0;
//         while (i < s.length() && isdigit(s[i])) {
//             int num = s[i] - '0';
//             n = n*10+num;
//             i++;
//         }
//         i++;
//         string cur = decodeString(s, i);
//         i++;
//         while (n-- > 0) {
//             res += cur;
//         }
//     }
// }


// buf := bytes.NewBuffer([]byte{})
// buf.WriteByte(s[*pos])
// buf.WriteString(latter)
// return buf.String()


// public String decodeString(String s) {
//     Deque<Character> queue = new LinkedList<>();
//     for (char c : s.toCharArray()) queue.offer(c);
//     return helper(queue);
// }
//
// public String helper(Deque<Character> queue) {
//     StringBuilder sb = new StringBuilder();
//     int num = 0;
//     while (!queue.isEmpty()) {
//         char c= queue.poll();
//         if (Character.isDigit(c)) {
//             num = num * 10 + c - '0';
//         } else if (c == '[') {
//             String sub = helper(queue);
//             for (int i = 0; i < num; i++) sb.append(sub);   
//             num = 0;
//         } else if (c == ']') {
//             break;
//         } else {
//             sb.append(c);
//         }
//     }
//     return sb.toString();
// }

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Decode String.
// Memory Usage: 2.3 MB, less than 32.01% of Go online submissions for Decode String.
func decodeString(s string) string {
    idx := 0
    // s = "1[" + s + "]"
    ans := dfsa394(s, &idx)
    return ans
}

func dfsa394(s string, idx *int) string {
    if *idx >= len(s) || s[*idx] == ']' {
        return ""
    }
    cnt := 1
    ans := ""
    if s[*idx] >= '0' && s[*idx] <= '9' {
        cnt = 0
        for s[*idx] >= '0' && s[*idx] <= '9' {
            cnt *= 10
            cnt += int(s[*idx] - '0')
            *idx++
        }
        // 有数字 后面必然是 [
        *idx++
        t2 := dfsa394(s, idx)
        for i := 0; i < cnt; i++ {
            ans += t2
        }
        *idx++          // 跳过 ]
        // return ans
        ans = ans + dfsa394(s, idx)
    } else {
        arr := []byte{}
        for *idx < len(s) && s[*idx] >= 'a' && s[*idx] <= 'z' {
            // ans += s[*idx]
            arr = append(arr, s[*idx])
            *idx++
        }
        // *idx++      // 跳过 ]
        ans = string(arr)
        // return ans
        ans = ans + dfsa394(s, idx)
    }
    // fmt.Println(s[*idx], *idx)
    return ans
}


func main_LT0394_20220125() {
// func main() {

    fmt.Println("ans:")

    // s := "3[a]2[bc]"
    // s := "3[a2[c]]"
    s := "2[abc]3[cd]ef"

    fmt.Println(decodeString(s))

}

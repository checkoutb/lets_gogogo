package sdq
// package main

import (
    "fmt"
    "sort"
)






// Runtime42 ms
// Beats
// 38.89%
// Memory6.9 MB
// Beats
// 72.22%

// Runtime32 ms
// Beats
// 44.44%
// Memory7.2 MB
// Beats
// 44.44%
// .. bu jili a.

func sortVowels(s string) string {
    // fmt.Println("--", []byte(s))
    // fmt.Println("--", []rune(s))
    // vc := []byte(s);
    // fmt.Println("..", string(vc))
    // // vc := []rune(s);     // cannot case []rune -> []byte
    // // fmt.Println(".,", string(vc));
    // vr := []rune(s);
    // fmt.Println("..", string(vr));

    vowels := map[byte]bool {
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
        'A': true,
        'E': true,
        'I': true,
        'O': true,
        'U': true,
    }

    vc := []byte(s);
    v2 := []byte{};
    for i, c := range vc {
        if _, okay := vowels[c]; okay == true {
            v2 = append(v2, c);
            vc[i] = 0
        }
    }

    // sort.Bytes(v2);
    // sort.Sort(v2);
    sort.Slice(v2, func(i, j int) bool { return v2[i] < v2[j] });

    idx := 0
    for i, c := range vc {
        if c == 0 {
            vc[i] = v2[idx];
            idx += 1;
        }
    }

    return string(vc);
}



func main_LT2785_20231010() {
// func main() {

    // var s = "lEetcOde";
    var s = "lYmpH";

    fmt.Println("ans:", sortVowels(s))


    fmt.Println(" test \n")

    // fmt.Println(" ", s[0].(type));

    // switch tp := s[0].(type) {      // s[0] is not a interface
    // case int:
    //     fmt.Println("1")
    // case byte:
    //     fmt.Println("2");
    // default: 
    //     fmt.Println("??");
    // }

}

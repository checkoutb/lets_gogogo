// package sdq
package main

import (
    "fmt"
    "sort"
    "strconv"
)

// D D

// m := make(map[string][]int)
// for _, t := range access_times {
//     time, _ := strconv.Atoi(t[1])
//     minute := 60*(time/100) + time%100
//     m[t[0]] = append(m[t[0]], minute)
// }



// Runtime31 ms
// Beats
// 17.80%
// Memory6.9 MB
// Beats
// 85.59%
func findHighAccessEmployees(access_times [][]string) []string {
    var map2 map[string][]string = map[string][]string{}

    for _, vs := range access_times {
        map2[vs[0]] = append(map2[vs[0]], vs[1])
    }

    var ans []string = []string{}

    sameHour := func(st, en string) bool {
        t2, _ := strconv.Atoi(st);
        t3, _ := strconv.Atoi(en);

        t2 = t2 / 100 * 60 + t2 % 100
        t3 = t3 / 100 * 60 + t3 % 100

        return t3 - t2 < 60;
    }

    for k, vs := range map2 {
        sort.Strings(vs)
        for i := 2; i < len(vs); i++ {
            if sameHour(vs[i - 2], vs[i]) {
                ans = append(ans, k)
                break
            }
        }
    }

    return ans;
}

func main_LT2933_20231205() {
// func main() {

    fmt.Println("ans:")


}

// package sdq
package main

import (
    "fmt"
    "math"
)



// Runtime706 ms
// Beats
// 100%
// Sorry, there are not enough accepted submissions to show data
// Memory11.9 MB
// Beats
// 100%

// /2 ,/3 ... /300  graph connected?  uf
func canTraverseAllPairs(nums []int) bool {
    
    prm := prime1(nums)

    sz1 := len(nums)

    map2 := make(map[int][]int)

    for idx, n := range nums {
        for i := 0; i < len(prm) && n > prm[i]; i++ {
            if n % prm[i] == 0 {
                map2[prm[i]] = append(map2[prm[i]], idx)
            }
            for n % prm[i] == 0 {
                n /= prm[i]
            }
        }
        if n != 1 {             // kan le biji cai faxian
            map2[n] = append(map2[n], idx)
        }
    }

    uf := make([]int, sz1)
    for i := range uf {
        uf[i] = -1
    }
    
    merge1 := func(i, j int) {
        if ufa(i, uf) != ufa(j, uf) {               // [30, 30] 因为使用-1 作为标记的，所以这里必须判断是否相同，不能直接赋予
            uf[ufa(i, uf)] = ufa(j, uf)
        }
    }

    for _, v := range map2 {
        for i := 1; i < len(v); i++ {
            merge1(v[i], v[i - 1])
        }
    }

    t2 := ufa(0, uf)
    for i := range uf {
        if t2 != ufa(i, uf) {
            return false
        }
    }
    return true
}

//ufa :=        // cannot recursive ?
func ufa(i int, uf []int) int {
    if (uf[i] == -1) {
        return i
    }

    fmt.Println(uf)

    uf[i] = ufa(uf[i], uf)
    return uf[i]
}

func prime1(nums []int) []int {
    mx := 1
    for _, n := range nums {
        if (n > mx) {
            mx = n
        }
    }
    sqt := int(math.Sqrt(float64(mx)))
    vb := make([]bool, sqt + 1)     // be divided ?
    for i := range vb {
        vb[i] = false
    }
    arr := []int{}
    for i := 2; i <= sqt; i++ {
        if !vb[i] {
            arr = append(arr, i)
            for a := i + i; a <=sqt; a += i {
                vb[a] = true
            }
        }
    }
    return arr
}



func main_LT2709_20230528() {
// func main() {

    // arr := []int{2,3,6}
    arr := []int{30,30}

    fmt.Println("ans:", canTraverseAllPairs(arr))


}

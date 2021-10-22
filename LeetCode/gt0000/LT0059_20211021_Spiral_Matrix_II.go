// package sdq
package main

import (
    "fmt"
    // "sync"
)


// left++
// right--
// top++
// bottom--


// A, lo = [], n*n+1
// while lo > 1:
//     lo, hi = lo - len(A), lo
//     A = [range(lo, hi)] + zip(*A[::-1])
// return A



// sync.WaitGroup 会死锁。。 之前有错误，就是 没有 等待 go完成，而是直接返回了。 所以 大部分时候 会缺 最后一个数字 。 n=2的时候大部分缺， 1和3好像不缺。。
// Runtime: 4 ms, faster than 8.93% of Go online submissions for Spiral Matrix II.
// Memory Usage: 2.3 MB, less than 41.96% of Go online submissions for Spiral Matrix II.
func generateMatrix(n int) [][]int {
    ans := make([][]int, n)
    for i := 0; i < n; i++ {
        ans[i] = make([]int, n)
    }
    // var wg sync.WaitGroup
    // ch := <-
    i, j, v := 0, -1, 1
    ch := make(chan int)        // no buffer
    ch2 := make(chan int)
    fun2 := func() {
        for v <= n * n {
            for j + 1 < n && ans[i][j + 1] == 0 {
                j++
                ans[i][j] = <- ch
                v++
            }
            for i + 1 < n && ans[i + 1][j] == 0 {
                i++
                ans[i][j] = <- ch
                v++
            }
            // fmt.Println(i, ", ", j)
            for j - 1 >= 0 && ans[i][j - 1] == 0 {
                j--
                fmt.Println("zzz, ", i, ", ", j)
                t2 := <- ch
                ans[i][j] = t2 
                // ans[i][j] = <- ch
                fmt.Println(",,,, ", ans[i][j])
                v++
            }
            for i - 1 >= 0 && ans[i - 1][j] == 0 {
                i--
                ans[i][j] = <- ch
                v++
            }
        }
        // fmt.Println("go end")
        ch2 <- 1}
    
    // wg.Add(1)
    go fun2()
    // defer fun2()     // no buffer 死锁。
    for a := 1; a <= n * n; a++ {
        ch <- a
        fmt.Println(a)
    }
    // close(ch)           // 等待数据用完 还是 直接close。  close后， <- ch 会报错吧。
    // wg.Wait()
    <- ch2
    return ans
}


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Spiral Matrix II.
// Memory Usage: 2.5 MB, less than 6.25% of Go online submissions for Spiral Matrix II.
func generateMatrix2(n int) [][]int {
    ans := make([][]int, n)
    for i := 0; i < n; i++ {
        ans[i] = make([]int, n)
    }
    i, j, v := 0, -1, 1
    for v <= n * n {
        for j + 1 < n && ans[i][j + 1] == 0 {
            j++
            ans[i][j] = v
            v++
        }
        for i + 1 < n && ans[i + 1][j] == 0 {
            i++
            ans[i][j] = v
            v++
        }
        for j - 1 >= 0 && ans[i][j - 1] == 0 {
            j--
            ans[i][j] = v
            v++
        }
        for i - 1 > 0 && ans[i - 1][j] == 0 {
            i--
            ans[i][j] = v
            v++
        }
    }
    return ans
}


func main_LT0059_20211021() {
// func main() {

    fmt.Println("ans:")

    // n := 3
    // n := 1
    n := 2

    ans := generateMatrix(n)

    fmt.Println(ans)

}

// package sdq
package main

import (
    "fmt"
)

// D 

// Fix one and count pairs

// 第一个人，拿 [0, min(n, limit)]
// 剩下的糖果就是 后2个人分： 
//   可能的分法： 
//      第二个人 拿到limit个 (当然 得保证第三个人不能是 负数) -> 第二个人拿0个(得保证第三个人不能 超过limit)
//      就是计算出 第二个 最多多少个， 最少多少个， 相减 就是 分配的方法 个数

// 3sum， 一段固定，计算 pair





// g


/*

nCm * (n-m)C(q) * (n-m-q)Cp

n! / (m! * (n-m)!) * (n-m)! / (q! * (n-m-q)!) * (n-m-q)! / (p! * (n-m-q-p)!)

n! / m! * 1/q! * 1/(p! * (n-m-q-p)!)

n=m+q+p

n! / m! * 1/q! * 1/p!

n! / (m! * q! * p!)

n! / (m! * q! * (n-m-q)!)

no no no
p = n-m-q   so  (n-m-q)Cp = 1

nCm * (n-m)Cq == n! / (m! * q! * (n-m-q)!)

*/


// n <= limit,  3^n,  no, candy are same.
// n > limit*3,  0
// limit < n <= limit*3
func distributeCandies(n int, limit int) int64 {

}


// tle
// n, limit < 10^6
func distributeCandies1(n int, limit int) int64 {
	var vi []int64 = make([]int64, n+1)
	vi[n] = 1
	var v2 []int64
	for i := 0; i < 3; i++ { // 3 children
		v2 = make([]int64, n+1)
		for j := 0; j <= n; j++ { // candy
			mxk := min(limit, j)
			for k := 0; k <= mxk; k++ {
				v2[j-k] += vi[j]
			}
		}
		vi = v2
		// fmt.Println(vi)
	}
	return vi[0]
}


func main_LT2929_20231205() {
// func main() {

    a,b := 5,2

    fmt.Println("ans:", distributeCandies(a, b))


}

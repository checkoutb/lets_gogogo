// package sdq
package main

import (
    "fmt"
)



// int mid = l + (r-l)/2;
// go.int == go.int64


// (1..n).bsearch(&method(:is_bad_version))
// (1..n).bsearch { |i| is_bad_version(i) }
// return bisect.bisect(type('', (), {'__getitem__': lambda self, i: isBadVersion(i)})(), False, 0, n)
// rb, py 内置二分。。。


// Runtime: 0 ms, faster than 100.00% of Go online submissions for First Bad Version.
// Memory Usage: 1.9 MB, less than 48.27% of Go online submissions for First Bad Version.
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
//
func firstBadVersion(n int) int {
    l, r := 1, n
    for l < r {
        mid := (l + r) / 2
        if isBadVersion(mid) {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}


func main_LT0278_20211130() {
// func main() {

    fmt.Println("ans:")


}

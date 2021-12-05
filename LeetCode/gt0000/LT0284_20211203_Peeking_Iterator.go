// package sdq
package main

import (
    "fmt"
)


// nextVal *int


// func (this *PeekingIterator) hasNext() bool {
// 	return this.num > 0
// }
// func (this *PeekingIterator) next() int {
// 	num := this.num
// 	this.num = this.iter.next()
// 	return num
// }
// func (this *PeekingIterator) peek() int {
// 	return this.num
// }



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Peeking Iterator.
// Memory Usage: 2.5 MB, less than 44.90% of Go online submissions for Peeking Iterator.

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
//
type PeekingIterator struct {
    Val int
    Ext bool
    Iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{ 0, false, iter }
}

func (this *PeekingIterator) hasNext() bool {
    return this.Iter.hasNext() || this.Ext
}

func (this *PeekingIterator) next() int {
    if this.Ext {
        this.Ext = false
        return this.Val
    } else {
        return this.Iter.next()
    }
}

func (this *PeekingIterator) peek() int {
    this.Val, this.Ext = this.next(), true
    return this.Val
}

func main_LT0284_20211203() {
// func main() {

    fmt.Println("ans:")


}

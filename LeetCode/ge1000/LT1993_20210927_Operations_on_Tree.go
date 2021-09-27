
package main

import (
    "fmt"
)

// D D

// struct Node {
//     vector<Node*> ch;
//     int lock = -1;
// };
// vector<Node> tree;
// vector<int> parent;
// LockingTree(vector<int>& parent) : parent(parent) {
//     tree = vector<Node>(parent.size());
//     for (int i = 1; i < parent.size(); ++i)
//         tree[parent[i]].ch.push_back(&tree[i]);
// }
// bool updateLock(Node &node, int oldVal, int newVal) {
//     if (node.lock != oldVal)
//         return false;
//     node.lock = newVal;
//     return true;
// }
// bool lock(int num, int user) { return updateLock(tree[num], -1, user); }
// bool unlock(int num, int user) { return updateLock(tree[num], user, -1); }
// bool upgrade(int num, int user) {
//     for (int i = num; i != -1; i = parent[i])
//         if (tree[i].lock != -1)
//             return false;
//     if (!anyLocked(&tree[num]))
//         return false;
//     unlock(&tree[num]);
//     return updateLock(tree[num], -1, user);
// }
// bool anyLocked(const Node *n) {
//     return n->lock != -1 || any_of(begin(n->ch), end(n->ch), [&](const auto *n){ return anyLocked(n); });
// }
// void unlock(Node *n) {
//     n->lock = -1;
//     for(auto ch : n->ch)
//         unlock(ch);
// }



// type LockingTree struct {
//     n int;
//     gr [][]int;
//     parent []int;
//     locked []int;
// }
// func Constructor(parent []int) LockingTree {
//     var this LockingTree;
//     this.parent = parent;
//     this.n = len(parent);
//     this.gr = make([][]int, this.n);
//     this.locked = make([]int, this.n);
//     for i := 0; i < this.n; i++ {
//         this.gr[i] = make([]int, 0);
//     }
//     for i, p := range parent {
//         this.locked[i] = -1;
//         if p == -1 {
//             continue;
//         }
//         this.gr[p] = append(this.gr[p], i);
//     }
//     return this;
// }




// Runtime: 356 ms, faster than 90.91% of Go online submissions for Operations on Tree.
// Memory Usage: 8.5 MB, less than 72.73% of Go online submissions for Operations on Tree.
type LockingTree struct {
    map2 map[int][]int
    parent []int
    locks []int
}

// ..Constructor，这个是普通方法吧？
func Constructor(parent []int) LockingTree {
    var lt LockingTree
    lt.map2 = make(map[int][]int)
    lt.parent = parent
    lt.locks = make([]int, len(parent));
    for idx, val := range parent {
        lt.map2[val] = append(lt.map2[val], idx)            // map2[qweasadw] == null ?  估计是 该类型的 0值
    }
    return lt
}


func (this *LockingTree) Lock(num int, user int) bool {
    if this.locks[num] != 0 {
        return false
    }
    this.locks[num] = user
    return true
}


func (this *LockingTree) Unlock(num int, user int) bool {
    if (this.locks[num] == user) {
        this.locks[num] = 0
        return true
    }
    return false
}


func (this *LockingTree) Upgrade(num int, user int) bool {
    if (this.locks[num] != 0) {
        return false
    }
    if !this.childLocked(num) {
        return false
    }
    if this.parentLocked(num) {
        return false
    }
    this.unlockChild(num)
    this.locks[num] = user          // 看example 中 应该是需要的
    return true
}

func (this *LockingTree) unlockChild(num int) {
    if this.locks[num] != 0 {
        this.locks[num] = 0
    }
    for _, child := range this.map2[num] {
        this.unlockChild(child)
    }
}

// all parent should be unlocked, so find if there is any parent is locked, no meaning all parent are unlocked
func (this *LockingTree) parentLocked(num int) bool {
    for num != -1 {
        if (this.locks[num] != 0) {
            return true
        }
        num = this.parent[num]
    }
    return false;

    // if num == -1 {
    //     return false;
    // }
    // if (this.locks[num] != 0) {
    //     return true
    // }
    // // it don't need revursion....
}

func (this *LockingTree) childLocked(num int) bool {
    if this.locks[num] != 0 {
        return true
    }
    for _, child := range this.map2[num] {          // nil ?        0值
        if (this.childLocked(child)) {
            return true
        }
    }
    return false
}


//func main_LT1993_20210927() {
func main() {
    
    arr := []int{-1, 0, 0, 1, 1, 2, 2}

    fmt.Println("ans:")

    lt := Constructor(arr)
    fmt.Println(lt.Lock(2, 2))
    fmt.Println(lt.Unlock(2, 3))
    fmt.Println(lt.Unlock(2, 2))
    fmt.Println(lt.Lock(4, 5))
    fmt.Println(lt.Upgrade(0, 1))
    fmt.Println(lt.Lock(0, 1))

}

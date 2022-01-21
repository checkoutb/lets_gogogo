// package sdq
package main

import (
    "fmt"
)



// pri_q, 每个人 取一个， 拿出来以后，这个tweet的owner的 下一个 tweet 入队。  (tweet 指向了 前一个tweet)



// class Twitter {
//     private:
//         vector<pair<int,int>>posts;
//         unordered_map<int, unordered_map<int, int>>follows;
//     public:
//         Twitter() {}
        
//         void postTweet(int userId, int tweetId) {
//             posts.push_back(make_pair(userId, tweetId));
//         }
        
//         vector<int> getNewsFeed(int userId) {
//             vector<int>feed;
//             int count = 0;
//             for(int i = posts.size() - 1; i >= 0 && count < 10; i--)
//                 if(posts[i].first == userId || follows[userId][posts[i].first])
//                     feed.push_back(posts[i].second), count++;
//             return feed;
//         }
        
//         void follow(int followerId, int followeeId) {
//             follows[followerId][followeeId] = 1;
//         }
        
//         void unfollow(int followerId, int followeeId) {
//             follows[followerId][followeeId] = 0;
//         }
//     };
// 不过 人多了以后，直接崩。。。


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Design Twitter.
// Memory Usage: 2.2 MB, less than 64.00% of Go online submissions for Design Twitter.
type Twitter struct {
    map2 map[int][]Tweet        // new is last
    map3 map[int]map[int]bool              // key/follower follow value/followee
    ord int
}

type Tweet struct {
    ord int
    id int
}

func Constructor() Twitter {
    return Twitter{ map[int][]Tweet{}, map[int]map[int]bool{}, 1 }
}

func (this *Twitter) PostTweet(userId int, tweetId int)  {
    twt := Tweet{ this.ord, tweetId }
    this.ord++
    if len(this.map2[userId]) >=10 {
        this.map2[userId] = this.map2[userId][1 : ]
    }
    this.map2[userId] = append(this.map2[userId], twt)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
    arr := [][]Tweet{}
    iarr := []int{}
    arr = append(arr, this.map2[userId])
    iarr = append(iarr, len(arr[len(arr) - 1]) - 1)
    for k, _ := range this.map3[userId] {
        arr = append(arr, this.map2[k])
        iarr = append(iarr, len(arr[len(arr) - 1]) - 1)
    }
    ans := []int{}

    for len(ans) < 10 {
        idx := -1
        for i := 0; i < len(iarr); i++ {
            if iarr[i] >= 0 {
                if idx == -1 {
                    idx = i
                } else {
                    if arr[idx][iarr[idx]].ord < arr[i][iarr[i]].ord {
                        idx = i
                    }
                }
            }
        }
        if idx == -1 {
            break
        } else {
            ans = append(ans, arr[idx][iarr[idx]].id)
            iarr[idx]--
        }
    }
    return ans
}

func (this *Twitter) Follow(followerId int, followeeId int)  {
    if _, ok := this.map3[followerId]; !ok {
        this.map3[followerId] = map[int]bool{}
    }
    this.map3[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    delete(this.map3[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */


// func main_LT0355_20220119() {
func main() {

    fmt.Println("ans:")

    twitter := Constructor()
    twitter.PostTweet(1, 5); // User 1 posts a new tweet (id = 5).
    fmt.Println(twitter.GetNewsFeed(1));  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
    twitter.Follow(1, 2);    // User 1 follows user 2.
    twitter.PostTweet(2, 6); // User 2 posts a new tweet (id = 6).
    fmt.Println(twitter.GetNewsFeed(1));  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
    twitter.Unfollow(1, 2);  // User 1 unfollows user 2.
    fmt.Println(twitter.GetNewsFeed(1));  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.


}

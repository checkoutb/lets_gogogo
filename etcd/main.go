package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"go.etcd.io/etcd/clientv3"
)

var gClient *clientv3.Client

func main() {

// ---------- conn st
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"172.17.0.2:2333"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("------conn fail, err: %v", err)
		return
	}
	defer client.Close()
	fmt.Println("conn ok")
// ----------- conn en

	var mykey string = "my key";
	setKV(client, mykey, "my value")

// ----------
	getV(client, mykey, "")
	getV(client, "a key", "")
	getV(client, "not exist key", "")
	getV(client, "not exist key", "")


}

func setKV(client *clientv3.Client, key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	resp, err := client.Put(ctx, key, value)
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	fmt.Println("---------------setKV response\n", resp);
}

// 6, no overload
// func getV(client *clientv3.Client, key string) (string, error) {
// 	return getV(client, key, "")
// }
func getV(client *clientv3.Client, key, defaultValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	resp, err := client.Get(ctx, key);
	cancel()
	fmt.Println("----------------getV response\n", resp)
	var value string = defaultValue;
	if resp.Kvs == nil {
		fmt.Println("not exist")
	} else {
		fmt.Println(string((*(resp.Kvs[0])).Key))
		fmt.Println(string(resp.Kvs[0].Value))		// auto de-ref, don't need *
		// fmt.Println(mvccpb.String(resp.Kvs[0]))
		fmt.Println(resp.Kvs[0].String())
		// fmt.Println(resp.Kvs[0].XXX_Unmarshal(resp.Kvs[0].Value))
		value = string((*resp.Kvs[0]).Value)
	}
	return value, err
}

func getCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second * 5)
}

func CAS () {

}

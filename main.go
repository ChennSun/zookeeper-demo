package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func main() {
	conn, eventChan, err := zk.Connect([]string{"127.0.0.1:2181", "127.0.0.1:2182", "127.0.0.1:2183"}, time.Second * 5)
	if(err != nil){
		fmt.Println(err)
	}
	go func() {
		for  {
			select {
			case event := <-eventChan:
				fmt.Printf("%s %s %s \n", event.Type, event.State, event.Server)
			}
		}
	}()
	res, err := conn.Create("/test", []byte("我的节点"), 1, []zk.ACL{{
		Perms:  31,
		Scheme: "world",
		ID:     "anyone",
		}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	time.Sleep(time.Second * 25)
	//conn, _, err := zk.Connect([]string{"127.0.0.1:2181", "127.0.0.1:2182", "127.0.0.1:2183"}, time.Second * 15)
	//if(err != nil){
	//	fmt.Println(err)
	//}
	//res2, _, event, err := conn.GetW("/test")
	//if err != nil{
	//	fmt.Print(err)
	//}
	//fmt.Println(res2)
	//go func() {
	//	for  {
	//		select {
	//		case msg := <-event:
	//			fmt.Println(msg)
	//		}
	//	}
	//}()
	//select {
	//
	//}
}
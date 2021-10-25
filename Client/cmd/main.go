package main

import (
	"fmt"
	"go-demo/Client/kernel"
	"go-demo/Services/proto"
)

func main() {
	kernel.Init("localhost:8900")
	CallRPC()
	CallFeed1()
	CallFeed2()
}

func CallRPC() {
	args := proto.SelectCompaniesByNameArgs{
		Name: "黑猫",
	}
	reply := proto.SelectCompaniesByNameReply{

	}
	kernel.GetEdcProxy().SelectCompaniesByName(&args, &reply)
	for _, v := range reply.Id {
		fmt.Println(v)
	}
}

func CallFeed1(){
	args := proto.IndexFeedArgs{
		Type: 1,
	}
	reply := proto.IndexFeedReply{}
	err := kernel.GetEdcProxy().Feed(&args, &reply)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range reply.Title {
		fmt.Println(v)
	}
	fmt.Println(len(reply.Title))
}

func CallFeed2(){
	args := proto.IndexFeedArgs{
		Type: 1,
	}
	reply := proto.IndexFeedReply{}
	err := kernel.GetEdcProxy().Feed(&args, &reply)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range reply.Title {
		fmt.Println(v)
	}
	fmt.Println(len(reply.Title))
}

package main

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	"go-demo/Server/db"
	"go-demo/Server/rpc"
)

var (
	dataSource = "root:Maple510809772@@tcp(127.0.0.1:3306)/guoyutong?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true"
	//dataSource = "root:root007@tcp(172.16.252.22:3306)/tousu?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true"
	addr       = "localhost:8900"
)

var s *server.Server

type AppFramework struct {
}

func (receiver *AppFramework) Run() {
	receiver.initEverything()
	receiver.StartRPCX()
}

func (receiver *AppFramework) initEverything(){
	err := receiver.initDB()
	if err != nil {
		fmt.Println("err in initDB", err)
		return
	}
	err = receiver.initRPCX()
	if err != nil {
		fmt.Println("err in initRPCX", err)
		return
	}
}

func (receiver *AppFramework) initDB() error{
	return db.InitDefaultDbEngine(dataSource)
}

func (receiver *AppFramework) initRPCX() error {
	s = server.NewServer()
	return s.Register(new(rpc.ComplaintHandler), "")
}

func (receiver *AppFramework) initRedis() error {
	return nil
}

func (receiver *AppFramework) initMemCache() error {
	return nil
}

func (receiver *AppFramework) initRPC() error {
	return nil
}

func (receiver *AppFramework) StartRPCX() {
	err := s.Serve("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
}


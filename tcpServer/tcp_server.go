package main

import (
	"entryTask/common/log"
	"entryTask/protocal/entry_task/pb"
	"entryTask/tcpServer/Dao"
	"entryTask/tcpServer/config"
	"entryTask/tcpServer/logic"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := log.Init()
	if err != nil {
		return
	}
	err = config.Init()
	if err != nil {
		log.Log.Errorf("err:%s", err)
		return
	}

	Dao.InitRedisPool()

	err = logic.InitCache()
	if err != nil {
		log.Log.Errorf("init cache err:%s", err)
		return
	}

	err = Dao.InitDB()
	if err != nil {
		log.Log.Errorf("err:%s", err)
		return
	}

	log.Log.Infof("server is run...")
	lis, err := net.Listen("tcp", config.Config.ServerAddr)
	if err != nil {
		log.Log.Errorf("listen err:%s", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterEntryTaskServer(s, &logic.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Log.Errorf("failed to serve: %v", err)
		return
	}

	select {}
}

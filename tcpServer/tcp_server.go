package main

import (
	"encoding/gob"
	"entryTask/common/log"
	"entryTask/common/zrpc"
	"entryTask/protocal"
	"entryTask/tcpServer/Dao"
	"entryTask/tcpServer/config"
	"entryTask/tcpServer/logic"
)

// RegisterRpcStruct 注册RPC结构
func RegisterRpcStruct() {
	gob.Register(protocal.SetSessionInfoRequest{})
	gob.Register(protocal.SetSessionInfoReply{})
	gob.Register(protocal.GetUserInfoRequest{})
	gob.Register(protocal.GetUserInfoReply{})
	gob.Register(protocal.UpdateUserInfoRequest{})
	gob.Register(protocal.UpdateUserInfoReplay{})
	gob.Register(protocal.GetSessionInfoRequest{})
	gob.Register(protocal.GetSessionInfoReply{})
	gob.Register(protocal.RefreshSessionRequest{})
	gob.Register(protocal.RefreshSessionReply{})
}

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

	RegisterRpcStruct()

	srv := zrpc.NewServer(config.Config.ServerAddr)
	srv.Register("SetSessionInfo", logic.SetSessionInfo)
	srv.Register("GetUserInfo", logic.GetUserInfo)
	srv.Register("UpdateUserInfo", logic.UpdateUserInfo)
	srv.Register("GetSessionInfo", logic.GetSessionInfo)
	srv.Register("RefreshSession", logic.RefreshSession)

	log.Log.Infof("server is run...")

	go srv.Run()

	select {}
}

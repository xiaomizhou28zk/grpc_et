package handleLogic

import (
	"encoding/gob"
	"encoding/json"
	"entryTask/common/log"
	"entryTask/common/zrpc"
	"entryTask/httpServer/common"
	"entryTask/protocal"
	"fmt"
	"net"
)

func init() {
	gob.Register(protocal.GetUserInfoRequest{})
	gob.Register(protocal.GetUserInfoReply{})
	gob.Register(protocal.UpdateUserInfoRequest{})
	gob.Register(protocal.UpdateUserInfoReplay{})
	gob.Register(protocal.GetSessionInfoRequest{})
	gob.Register(protocal.GetSessionInfoReply{})
	gob.Register(protocal.SetSessionInfoRequest{})
	gob.Register(protocal.SetSessionInfoReply{})
	gob.Register(protocal.RefreshSessionRequest{})
	gob.Register(protocal.RefreshSessionReply{})
}

// getUserInfoRpc 获取用户信息
func getUserInfoRpc(uid string) (userInfo common.UserInfo, err error) {

	conn, err := common.MyPool.Get()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return
	}
	cli := zrpc.NewClient(conn.(net.Conn))

	cli.Call("GetUserInfo", &protocal.GetUserInfo)
	req := protocal.GetUserInfoRequest{
		Uid: uid,
	}

	u, err := protocal.GetUserInfo(req)
	if err != nil {
		log.Log.Errorf("GetUserInfo err:%s", err)
		fmt.Println("GetUserInfo err:", err)
		return userInfo, err
	}

	userInfo = common.UserInfo{
		UID:      uid,
		Nick:     u.Nick,
		Picture:  u.Pic,
		PassWord: u.Pwd,
	}

	_ = common.MyPool.Put(conn)
	return userInfo, nil
}

// updateUserInfoRpc 更新用户信息
func updateUserInfoRpc(info common.UserInfo) error {
	conn, err := common.MyPool.Get()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return err
	}
	cli := zrpc.NewClient(conn.(net.Conn))

	cli.Call("UpdateUserInfo", &protocal.UpdateUserInfo)
	req := protocal.UpdateUserInfoRequest{
		Uid:  info.UID,
		Nick: info.Nick,
		Pic:  info.Picture,
	}

	r, err := protocal.UpdateUserInfo(req)
	if err != nil {
		log.Log.Errorf("err:%s", err)
		return err
	}

	if r.Ret != 0 {
		log.Log.Errorf("err ret:%d", r.Ret)
		return fmt.Errorf("ret:%d", r.Ret)
	}

	_ = common.MyPool.Put(conn)

	return nil
}

// getSessionInfo 获取会话信息
func getSessionInfo(sessionID string) (sessionInfo common.SessionInfo, err error) {
	sessionInfo = common.SessionInfo{}

	conn, err := common.MyPool.Get()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		fmt.Println("session get con err", err)
		return
	}

	cli := zrpc.NewClient(conn.(net.Conn))

	cli.Call("GetSessionInfo", &protocal.GetSessionInfo)
	req := protocal.GetSessionInfoRequest{
		SessionID: sessionID,
	}

	r, err := protocal.GetSessionInfo(req)
	if err != nil {
		log.Log.Errorf("err:%s", err)
		return
	}

	_ = json.Unmarshal([]byte(r.SessionInfo), &sessionInfo)
	sessionInfo.SessionID = sessionID

	_ = common.MyPool.Put(conn)

	return sessionInfo, nil
}

// setSessionRpc 设置会话
func setSessionRpc(sessionID string, info common.SessionInfo) {
	conn, err := common.MyPool.Get()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		fmt.Println("session get con err", err)
		return
	}

	cli := zrpc.NewClient(conn.(net.Conn))

	cli.Call("SetSessionInfo", &protocal.SetSessionInfo)
	infoStr, _ := json.Marshal(info)
	req := protocal.SetSessionInfoRequest{
		SessionID:   sessionID,
		SessionInfo: string(infoStr),
	}
	r, err := protocal.SetSessionInfo(req)
	if err != nil {
		fmt.Printf("session err:%s   uid:%s\n", err, info.UID)
		log.Log.Errorf("err:%s", err)
		return
	}
	if r.Ret != 0 {
		log.Log.Errorf("err ret:%d", r.Ret)
		return
	}
	log.Log.Debugf("set session ok")

	err = common.MyPool.Put(conn)
	if err != nil {
		fmt.Println("session put err:", err)
	}
}

// refreshSessionRpc 刷新会话
func refreshSessionRpc(sessionID string) {
	conn, err := common.MyPool.Get()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return
	}
	cli := zrpc.NewClient(conn.(net.Conn))

	cli.Call("RefreshSession", &protocal.RefreshSession)

	req := protocal.RefreshSessionRequest{
		SessionID: sessionID,
	}
	r, err := protocal.RefreshSession(req)
	if err != nil {
		log.Log.Errorf("err:%s", err)
		return
	}
	if r.Ret != 0 {
		log.Log.Errorf("err ret:%d", r.Ret)
		return
	}
	log.Log.Debugf("refresh session ok")
	_ = common.MyPool.Put(conn)
}

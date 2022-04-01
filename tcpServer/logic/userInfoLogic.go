package logic

import (
	"context"
	"entryTask/common/log"
	"entryTask/protocal"
	"entryTask/tcpServer/Dao"
	_ "github.com/xiaomizhou28zk/grpc_et/protocal/pb"
)

// GetUserInfo 获取用户信息
func GetUserInfo(in protocal.GetUserInfoRequest) (protocal.GetUserInfoReply, error) {

	rsp := protocal.GetUserInfoReply{
		Ret: 0,
	}

	userInfo := Dao.UserInfo{}
	var err error
	v, ok := UserCache.Get(in.Uid)
	if ok {
		userInfo = v.(Dao.UserInfo)
	} else {
		userInfo, err = Dao.QueryUserInfo(in.Uid)
		if err != nil {
			rsp.Ret = -1
			return rsp, nil
		}
		UserCache.Add(in.Uid, userInfo)
	}

	rsp.Uid = userInfo.Uid
	rsp.Nick = userInfo.Nick
	rsp.Pic = userInfo.Picture
	rsp.Pwd = userInfo.Password

	return rsp, nil

}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(in protocal.UpdateUserInfoRequest) (protocal.UpdateUserInfoReplay, error) {
	log.Log.Debugf("Received uid: %v", in.Uid)
	rsp := protocal.UpdateUserInfoReplay{
		Ret: 0,
	}
	if in.Nick == "" && in.Pic == "" {
		rsp.Ret = -1
		return rsp, nil
	}
	err := Dao.UpdateUserInfo(in.Uid, in.Nick, in.Pic)
	if err != nil {
		rsp.Ret = -1
		return rsp, nil
	}

	v, ok := UserCache.Get(in.Uid)
	if ok {
		u := v.(Dao.UserInfo)
		u.Nick = in.Nick
		u.Picture = in.Pic
		UserCache.Add(in.Uid, u)
	}

	return rsp, nil
}

// GetUserInfo 获取用户信息
func (s *Server)GetUserInfo(ctx context.Context, in *pb.HelloRequest) (protocal.GetUserInfoReply, error) {

	rsp := protocal.GetUserInfoReply{
		Ret: 0,
	}

	userInfo := Dao.UserInfo{}
	var err error
	v, ok := UserCache.Get(in.Uid)
	if ok {
		userInfo = v.(Dao.UserInfo)
	} else {
		userInfo, err = Dao.QueryUserInfo(in.Uid)
		if err != nil {
			rsp.Ret = -1
			return rsp, nil
		}
		UserCache.Add(in.Uid, userInfo)
	}

	rsp.Uid = userInfo.Uid
	rsp.Nick = userInfo.Nick
	rsp.Pic = userInfo.Picture
	rsp.Pwd = userInfo.Password

	return rsp, nil

}
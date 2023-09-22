package logic

import (
	"entryTask/common/log"
	"entryTask/protocal/entry_task/pb"
	"entryTask/tcpServer/Dao"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
)

// GetUserInfo 获取用户信息
func (s *Server) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {

	rsp := &pb.GetUserInfoResponse{}

	userInfo, err := getUserInfoByUid(req.GetUid())
	if err != nil {
		rsp.Ret = proto.Int32(-1)
		return rsp, nil
	}

	rsp.Uid = proto.String(userInfo.Uid)
	rsp.Nick = proto.String(userInfo.Nick)
	rsp.Pic = proto.String(userInfo.Picture)
	rsp.Pwd = proto.String(userInfo.Password)

	return rsp, nil

}

func getUserInfoByUid(uid string) (*Dao.UserInfo, error) {
	userInfo := Dao.UserInfo{}
	var err error
	v, ok := UserCache.Get(uid)
	if ok {
		userInfo = v.(Dao.UserInfo)
	} else {
		userInfo, err = Dao.QueryUserInfo(uid)
		if err != nil {
			return nil, err
		}
		UserCache.Add(uid, userInfo)
	}
	return &userInfo, nil
}

// UpdateUserInfo 更新用户信息
func (s *Server) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	log.Log.Debugf("Received uid: %v, nick:%s, pic:%s", req.GetUid(), req.GetNick(), req.GetPic())
	rsp := &pb.UpdateUserInfoResponse{}
	if req.GetNick() == "" && req.GetPic() == "" {
		rsp.Ret = proto.Int32(-1)
		return rsp, nil
	}
	err := Dao.UpdateUserInfo(req.GetUid(), req.GetNick(), req.GetPic())
	if err != nil {
		rsp.Ret = proto.Int32(-1)
		return rsp, nil
	}

	v, ok := UserCache.Get(req.GetUid())
	if ok {
		u := v.(Dao.UserInfo)
		u.Nick = req.GetNick()
		u.Picture = req.GetPic()
		UserCache.Add(req.GetUid(), u)
	}

	return rsp, nil
}

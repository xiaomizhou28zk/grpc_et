package logic

import (
	"entryTask/common/log"
	"entryTask/protocal/entry_task/pb"
	"entryTask/tcpServer/Dao"
	"entryTask/tcpServer/config"

	"github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
)

// GetSessionInfo 获取会话信息
func (s *Server) GetSessionInfo(ctx context.Context, req *pb.GetSessionInfoRequest) (*pb.GetSessionInfoResponse, error) {

	rsp := &pb.GetSessionInfoResponse{}
	c := Dao.RedisClient.Get()
	defer c.Close()

	ret, err := redis.String(c.Do("GET", req.GetSessionId()))
	if err != nil {
		log.Log.Errorf("err:%s", err)
		rsp.Ret = proto.Int32(-1)
		return rsp, err
	}
	rsp.SessionInfo = proto.String(ret)

	return rsp, nil
}

// RefreshSession 刷新会话信息
func (s *Server) RefreshSession(ctx context.Context, req *pb.RefreshSessionRequest) (*pb.RefreshSessionResponse, error) {

	rsp := &pb.RefreshSessionResponse{}

	c := Dao.RedisClient.Get()
	defer c.Close()

	_, err := c.Do("EXPIRE", req.GetSessionId(), config.Config.SessionExpireTime)
	if err != nil {
		rsp.Ret = proto.Int32(-1)
		return rsp, err
	}

	return rsp, nil
}

// SetSessionInfo 设置会话信息
func (s *Server) SetSessionInfo(ctx context.Context, req *pb.SetSessionInfoRequest) (*pb.SetSessionInfoResponse, error) {

	rsp := &pb.SetSessionInfoResponse{}

	c := Dao.RedisClient.Get()

	defer c.Close()

	_, err := c.Do("SETEX", req.GetSessionId(), config.Config.SessionExpireTime, req.GetSessionInfo())
	if err != nil {
		rsp.Ret = proto.Int32(-1)
		return rsp, err
	}

	return rsp, nil
}

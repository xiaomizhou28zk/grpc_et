package logic

import (
	"entryTask/common/log"
	"entryTask/protocal"
	"entryTask/tcpServer/Dao"
	"entryTask/tcpServer/config"

	"github.com/garyburd/redigo/redis"
)

// GetSessionInfo 获取会话信息
func GetSessionInfo(in protocal.GetSessionInfoRequest) (protocal.GetSessionInfoReply, error) {

	rsp := protocal.GetSessionInfoReply{}
	c := Dao.RedisClient.Get()
	defer c.Close()

	ret, err := redis.String(c.Do("GET", in.SessionID))
	if err != nil {
		log.Log.Errorf("err:%s", err)
		rsp.Ret = -1
		return rsp, err
	}
	rsp.SessionInfo = ret

	return rsp, nil
}

// RefreshSession 刷新会话信息
func RefreshSession(in protocal.RefreshSessionRequest) (protocal.RefreshSessionReply, error) {

	rsp := protocal.RefreshSessionReply{}

	c := Dao.RedisClient.Get()
	defer c.Close()

	_, err := c.Do("EXPIRE", in.SessionID, config.Config.SessionExpireTime)
	if err != nil {
		rsp.Ret = -1
		return rsp, err
	}

	return rsp, nil
}

// SetSessionInfo 设置会话信息
func SetSessionInfo(in protocal.SetSessionInfoRequest) (protocal.SetSessionInfoReply, error) {

	rsp := protocal.SetSessionInfoReply{}

	c := Dao.RedisClient.Get()

	defer c.Close()

	_, err := c.Do("SETEX", in.SessionID, config.Config.SessionExpireTime, in.SessionInfo)
	if err != nil {
		rsp.Ret = -1
		return rsp, err
	}

	return rsp, nil
}

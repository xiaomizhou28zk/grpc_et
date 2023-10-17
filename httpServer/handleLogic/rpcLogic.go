package handleLogic

import (
	"context"
	"encoding/json"
	"entryTask/common/log"
	"entryTask/httpServer/common"
	"entryTask/protocal/entry_task/pb"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func getClient() (pb.EntryTaskClient, *common.ConnRes, error) {
	conn, err := common.MyPool.Get()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return nil, nil, err
	}
	cli := pb.NewEntryTaskClient(conn.(grpc.ClientConnInterface))
	return cli, &conn, nil
}

// getUserInfoRpc 获取用户信息
func getUserInfoRpc(uid string) (userInfo common.UserInfo, err error) {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return
	}
	req := &pb.GetUserInfoRequest{
		Uid: proto.String(uid),
	}

	resp, err := cli.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Log.Errorf("GetUserInfo err:%s", err)
		return userInfo, err
	}

	userInfo = common.UserInfo{
		UID:      uid,
		Nick:     resp.GetNick(),
		Picture:  resp.GetPic(),
		PassWord: resp.GetPwd(),
	}

	_ = common.MyPool.Put(*conn)
	return userInfo, nil
}

// updateUserInfoRpc 更新用户信息
func updateUserInfoRpc(info common.UserInfo) error {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return err
	}

	req := &pb.UpdateUserInfoRequest{
		Uid:  proto.String(info.UID),
		Nick: proto.String(info.Nick),
		Pic:  proto.String(info.Picture),
	}

	resp, err := cli.UpdateUserInfo(context.Background(), req)
	if err != nil {
		log.Log.Errorf("err:%s", err)
		return err
	}

	if resp.GetRet() != 0 {
		log.Log.Errorf("err ret:%d", resp.GetRet())
		return fmt.Errorf("ret:%d", resp.GetRet())
	}

	_ = common.MyPool.Put(*conn)

	return nil
}

// getSessionInfo 获取会话信息
func getSessionInfo(sessionID string) (sessionInfo common.SessionInfo, err error) {
	sessionInfo = common.SessionInfo{}
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return sessionInfo, err
	}
	req := &pb.GetSessionInfoRequest{
		SessionId: proto.String(sessionID),
	}
	resp, err := cli.GetSessionInfo(context.Background(), req)
	if err != nil {
		log.Log.Errorf("err:%s", err)
		return
	}
	if resp.GetRet() != 0 {
		log.Log.Errorf("err ret:%d", resp.GetRet())
		return sessionInfo, fmt.Errorf("ret:%d", resp.GetRet())
	}

	_ = json.Unmarshal([]byte(resp.GetSessionInfo()), &sessionInfo)
	sessionInfo.SessionID = sessionID

	_ = common.MyPool.Put(*conn)

	return sessionInfo, nil
}

// setSessionRpc 设置会话
func setSessionRpc(sessionID string, info common.SessionInfo) {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return
	}
	infoStr, _ := json.Marshal(info)
	req := &pb.SetSessionInfoRequest{
		SessionId:   proto.String(sessionID),
		SessionInfo: proto.String(string(infoStr)),
	}
	resp, err := cli.SetSessionInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("session err:%s   uid:%s\n", err, info.UID)
		log.Log.Errorf("err:%s", err)
		return
	}
	if resp.GetRet() != 0 {
		log.Log.Errorf("err ret:%d", resp.GetRet())
		return
	}
	log.Log.Debugf("set session ok")

	err = common.MyPool.Put(*conn)
	if err != nil {
		fmt.Println("session put err:", err)
	}
}

// refreshSessionRpc 刷新会话
func refreshSessionRpc(sessionID string) {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return
	}
	req := &pb.RefreshSessionRequest{
		SessionId: proto.String(sessionID),
	}
	resp, err := cli.RefreshSession(context.Background(), req)
	if err != nil {
		log.Log.Errorf("err:%s", err)
		return
	}
	if resp.GetRet() != 0 {
		log.Log.Errorf("err ret:%d", resp.GetRet())
		return
	}
	log.Log.Debugf("refresh session ok")
	_ = common.MyPool.Put(*conn)
}

func getMessageListRpc(uid string, page, pageSize int32) (msgList []*pb.MessageInfo, count int32, err error) {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return
	}
	req := &pb.GetMessageListRequest{
		Uid:      proto.String(uid),
		PageSize: proto.Int32(pageSize),
		Page:     proto.Int32(page),
	}

	resp, err := cli.GetMessageList(context.Background(), req)
	if err != nil {
		log.Log.Errorf("GetMessageList err:%s", err)
		return nil, 0, err
	}
	_ = common.MyPool.Put(*conn)
	return resp.GetList(), resp.GetTotal(), nil
}

func publishMessage(uid, userName, msg string) error {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return err
	}
	req := &pb.PublishMessageRequest{
		Uid:      proto.String(uid),
		UserName: proto.String(userName),
		Message:  proto.String(msg),
	}

	resp, err := cli.PublishMessage(context.Background(), req)
	if err != nil {
		log.Log.Errorf("PublishMessage err:%s", err)
		return err
	}
	if resp.GetRet() != 0 {
		return errors.New("PublishMessage error")
	}
	_ = common.MyPool.Put(*conn)
	return nil
}

func deleteMessageRpc(uid string, msgId uint64) error {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return err
	}
	req := &pb.DeleteMessageRequest{
		Uid:   proto.String(uid),
		MsgId: proto.Uint64(msgId),
	}

	resp, err := cli.DeleteMessage(context.Background(), req)
	if err != nil {
		log.Log.Errorf("DeleteMessage err:%s", err)
		return err
	}
	if resp.GetRet() != 0 {
		return errors.New("DeleteMessage error")
	}
	_ = common.MyPool.Put(*conn)
	return nil
}

func deleteSessionRpc(sessionId string) error {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return err
	}
	req := &pb.DeleteSessionRequest{SessionId: proto.String(sessionId)}

	resp, err := cli.DeleteSession(context.Background(), req)
	if err != nil {
		log.Log.Errorf("DeleteSession err:%s", err)
		return err
	}
	if resp.GetRet() != 0 {
		return errors.New("DeleteSession error")
	}
	_ = common.MyPool.Put(*conn)
	return nil
}

func getCommentsByMessageIdsRpc(messageIds []uint64) (*pb.GetCommentsByMessageIdsResponse, error) {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return nil, err
	}

	req := &pb.GetCommentsByMessageIdsRequest{MessageIds: messageIds}

	resp, err := cli.GetCommentsByMessageIds(context.Background(), req)
	if err != nil {
		log.Log.Errorf("GetCommentsByMessageIds err:%s", err)
		return nil, err
	}
	if resp.GetRet() != 0 {
		return nil, errors.New("GetCommentsByMessageIds error")
	}
	_ = common.MyPool.Put(*conn)
	return resp, nil
}

func getReplyByCommentIdsRpc(commentIds []uint64) (*pb.GetReplyByCommentIdsResponse, error) {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return nil, err
	}

	req := &pb.GetReplyByCommentIdsRequest{CommentIds: commentIds}

	resp, err := cli.GetReplyByCommentIds(context.Background(), req)
	if err != nil {
		log.Log.Errorf("GetReplyByCommentIds err:%s", err)
		return nil, err
	}
	fmt.Println("GetReplyByCommentIds rpc:", len(resp.GetList()), resp.GetRet())
	if resp.GetRet() != 0 {
		return nil, errors.New("GetReplyByCommentIds error")
	}
	_ = common.MyPool.Put(*conn)
	return resp, nil
}

func setCommentRpc(uid, comment string) error {
	cli, conn, err := getClient()
	if err != nil {
		log.Log.Errorf("get conn err:%s", err)
		return err
	}

	req := &pb.SetCommentRequest{
		Comment: proto.String(comment),
		Uid:     proto.String(uid),
	}

	resp, err := cli.SetComment(context.Background(), req)
	if err != nil {
		log.Log.Errorf("SetComment err:%s", err)
		return err
	}
	fmt.Println("setCommentRpc rpc:", resp.GetRet())
	if resp.GetRet() != 0 {
		return errors.New("setCommentRpc error")
	}
	_ = common.MyPool.Put(*conn)
	return nil
}

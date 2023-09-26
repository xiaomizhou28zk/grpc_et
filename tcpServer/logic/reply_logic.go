package logic

import (
	"context"
	"entryTask/common/log"
	"entryTask/protocal/entry_task/pb"
	"entryTask/tcpServer/Dao"
)

func (s *Server) GetReplyByCommentIds(ctx context.Context, req *pb.GetReplyByCommentIdsRequest) (*pb.GetReplyByCommentIdsResponse, error) {
	rsp := &pb.GetReplyByCommentIdsResponse{}

	list, err := Dao.GetReplyListByCommentIds(req.GetCommentIds(), 0, 0)
	if err != nil {
		log.Log.Errorf("GetReplyByCommentIds err:%s", err)
		return rsp, err
	}
	for _, elem := range list {
		userInfo, err := getUserInfoByUid(elem.Uid)
		if err != nil {
			log.Log.Errorf("GetReplyByCommentIds getUserInfoByUid err:%s", err)
			continue
		}
		toUserInfo, err := getUserInfoByUid(elem.ToUid)
		if err != nil {
			log.Log.Errorf("GetReplyByCommentIds getUserInfoByUid err:%s", err)
			continue
		}
		rsp.List = append(rsp.List, &pb.CommentReplyItem{
			CommentId:  elem.CommentId,
			Reply:      elem.Reply,
			ReplyId:    elem.ID,
			Ctime:      elem.Ctime,
			Uid:        elem.Uid,
			ToUid:      elem.ToUid,
			UserName:   userInfo.Nick,
			ToUserName: toUserInfo.Nick,
		})
	}
	return rsp, nil
}

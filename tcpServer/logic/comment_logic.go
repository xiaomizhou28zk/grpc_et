package logic

import (
	"context"
	"entryTask/common/log"
	"entryTask/protocal/entry_task/pb"
	"entryTask/tcpServer/Dao"
)

func (s *Server) GetCommentsByMessageIds(ctx context.Context, req *pb.GetCommentsByMessageIdsRequest) (*pb.GetCommentsByMessageIdsResponse, error) {
	rsp := &pb.GetCommentsByMessageIdsResponse{}

	list, err := Dao.GetCommentListByMessageIds(req.GetMessageIds(), 0, 0)
	if err != nil {
		log.Log.Errorf("GetCommentsByMessageIds err:%s", err)
		return rsp, err
	}
	for _, elem := range list {
		userInfo, err := getUserInfoByUid(elem.Uid)
		if err != nil {
			log.Log.Errorf("GetReplyByCommentIds getUserInfoByUid err:%s", err)
			continue
		}
		rsp.List = append(rsp.List, &pb.MessageCommentsItem{
			MessageId: elem.MessageId,
			Comment:   elem.Comment,
			CommentId: elem.ID,
			Ctime:     elem.Ctime,
			Uid:       elem.Uid,
			UserName:  userInfo.Nick,
		})
	}
	return rsp, nil
}

func (s *Server) SetComment(ctx context.Context, req *pb.SetCommentRequest) (*pb.SetCommentResponse, error) {
	rsp := &pb.SetCommentResponse{}
	err := Dao.AddComment(req.GetUid(), req.GetComment(), req.GetMessageId())
	if err != nil {
		log.Log.Errorf("SetComment err:%s", err)
		return rsp, err
	}
	return rsp, nil
}

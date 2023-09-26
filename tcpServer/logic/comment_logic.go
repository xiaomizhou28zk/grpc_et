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
		rsp.List = append(rsp.List, &pb.MessageCommentsItem{
			MessageId: elem.MessageId,
			Comment:   elem.Comment,
			CommentId: elem.ID,
			Ctime:     elem.Ctime,
		})
	}
	return rsp, nil
}

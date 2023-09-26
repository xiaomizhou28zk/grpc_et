package logic

import (
	"context"
	"entryTask/common/log"
	"entryTask/protocal/entry_task/pb"
	"entryTask/tcpServer/Dao"
	"fmt"
)

func (s *Server) GetReplyByCommentIds(ctx context.Context, req *pb.GetReplyByCommentIdsRequest) (*pb.GetReplyByCommentIdsResponse, error) {
	rsp := &pb.GetReplyByCommentIdsResponse{}

	list, err := Dao.GetReplyListByCommentIds(req.GetCommentIds(), 0, 0)
	if err != nil {
		fmt.Println("GetReplyByCommentIds err:", err)
		log.Log.Errorf("GetReplyByCommentIds err:%s", err)
		return rsp, err
	}
	fmt.Println("GetReplyByCommentIds:", len(list))
	for _, elem := range list {
		fmt.Println("haha:", elem.Uid, elem.ToUid)
		userInfo, err := getUserInfoByUid(elem.Uid)
		if err != nil {
			fmt.Println("GetReplyByCommentIds: 11111", err, elem.Uid)
			log.Log.Errorf("GetReplyByCommentIds getUserInfoByUid err:%s", err)
			continue
		}
		toUserInfo, err := getUserInfoByUid(elem.ToUid)
		if err != nil {
			fmt.Println("GetReplyByCommentIds: 2222", err, elem.ToUid)
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
	fmt.Println("GetReplyByCommentIds222:", len(rsp.GetList()))
	return rsp, nil
}

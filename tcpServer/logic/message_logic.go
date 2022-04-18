package logic

import (
	"context"
	"entryTask/common/log"
	"entryTask/protocal/entry_task/pb"
	"entryTask/tcpServer/Dao"

	"google.golang.org/protobuf/proto"
)

func (s *Server) GetMessageList(ctx context.Context, req *pb.GetMessageListRequest) (*pb.GetMessageListResponse, error) {

	rsp := &pb.GetMessageListResponse{}

	msgList, err := Dao.GetMessageList(req.GetUid())
	if err != nil {
		log.Log.Errorf("GetMessageList err:%s", err)
	}
	for _, elem := range msgList {
		rsp.List = append(rsp.List, &pb.MessageInfo{
			Id:    proto.Uint64(elem.ID),
			Msg:   proto.String(elem.Message),
			Image: proto.String(elem.Image),
			Owner: proto.String(elem.Owner),
			Ctime: proto.Uint64(elem.Ctime),
			Mtime: proto.Uint64(elem.Mtime),
			Uid:   proto.String(elem.Uid),
		})
	}

	log.Log.Debugf("GetMessageList len:%d", len(rsp.GetList()))

	return rsp, nil

}

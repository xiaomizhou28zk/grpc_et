package handleLogic

import (
	"encoding/base64"
	"encoding/json"
	"entryTask/common/log"
	"entryTask/httpServer/common"
	"entryTask/httpServer/config"
	"net/http"
	"fmt"
)

type MessageInfo struct {
	UID         string     `json:"uid"` //用户ID
	Msg         string     `json:"msg"`
	Image       string     `json:"image"`
	ID          uint64     `json:"id"`
	Owner       string     `json:"owner"`
	CTime       string     `json:"c_time"`
	MTime       string     `json:"m_time"`
	CommentList []*Comment `json:"comment_list"`
}

type Reply struct {
	ID         uint64 `json:"id"`
	Uid        string `json:"uid"`
	UserName   string `json:"user_name"`
	ToUid      string `json:"to_uid"`
	ToUserName string `json:"to_user_name"`
	Reply      string `json:"reply"`
	CommentId  uint64 `json:"comment_id"`
}

type Comment struct {
	ID        uint64   `json:"id"`
	Comment   string   `json:"comment"`
	CTime     string   `json:"CTime"`
	MessageId uint64   `json:"message_id"`
	ReplyList []*Reply `json:"reply_list"`
	Uid       string   `json:"uid"`
	UserName  string   `json:"user_name"`
}

type getMsgListRsp struct {
	Ret      int32          `json:"ret"` //业务返回码
	List     []*MessageInfo `json:"list"`
	Count    int32          `json:"count"`
	Page     int32          `json:"page"`
	PageSize int32          `json:"page_size"`
}

type getMsgListRequest struct {
	Page       int32 `json:"page"`
	PageSize   int32 `json:"pageSize"`
	AllMessage bool  `json:"all_message"`
}

func GetMessageList(w http.ResponseWriter, r *http.Request) {
	rsp := getMsgListRsp{
		Ret: common.SucCode,
	}

	sessionInfo, status := checkSession(r)
	if !status {
		log.Log.Errorf("get session err")
		rsp.Ret = common.InvalidSession
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	// 解析 JSON 参数
	var req getMsgListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("-------e-------")
		rsp.Ret = common.MissingParams
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}
	if req.Page < 1 {
		req.PageSize = 1
	}
	if req.PageSize != 10 {
		req.PageSize = 10
	}
	rsp.Page = req.Page
	rsp.PageSize = req.PageSize

	uid := sessionInfo.UID
	if req.AllMessage {
		uid = ""
	}

	fmt.Println("getMessageListRpc +++++1")
	msgList, count, err := getMessageListRpc(uid, req.Page, req.PageSize)
	if err != nil {
		fmt.Println("getMessageListRpc err")
		log.Log.Errorf("getMessageListRpc err")
		rsp.Ret = common.ServerErrCode
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
	}
	fmt.Println("getMessageListRpc +++++2")

	messageIds := make([]uint64, 0)
	for _, elem := range msgList {
		messageIds = append(messageIds, elem.GetId())
	}
	msg2CommentMap := getCommentAndReply(messageIds)

	for _, elem := range msgList {
		// Base64 解码
		decoded, err := base64.StdEncoding.DecodeString(elem.GetMsg())
		if err != nil {
			log.Log.Errorf("getMessageListRpc err:%s", err)
			return
		}
		rsp.List = append(rsp.List, &MessageInfo{
			UID:         elem.GetUid(),
			ID:          elem.GetId(),
			Msg:         string(decoded),
			Image:       elem.GetImage(),
			Owner:       elem.GetOwner(),
			CTime:       common.GetTimeFromTimestamp(elem.GetCtime()),
			MTime:       common.GetTimeFromTimestamp(elem.GetMtime()),
			CommentList: msg2CommentMap[elem.GetId()],
		})
	}
	rsp.Count = count

	msg, _ := json.Marshal(rsp)
	_, _ = w.Write(msg)
}

func getCommentAndReply(messageIds []uint64) map[uint64][]*Comment {
	commentRsp, err := getCommentsByMessageIdsRpc(messageIds)
	if err != nil {
		return nil
	}
	commentIds := make([]uint64, 0)
	for _, elem := range commentRsp.List {
		commentIds = append(commentIds, elem.CommentId)
	}
	fmt.Println("comment list", commentIds)
	reply, err := getReplyByCommentIdsRpc(commentIds)
	if err != nil {
		fmt.Println("comment err:",err)
		return nil
	}
	fmt.Println("comment=---------- list", len(reply.GetList()))
	commentReplyMap := make(map[uint64][]*Reply)
	for _, elem := range reply.GetList() {
		fmt.Println("reply:", elem.GetReply())
		commentReplyMap[elem.CommentId] = append(commentReplyMap[elem.CommentId], &Reply{
			ID:         elem.ReplyId,
			Reply:      elem.Reply,
			CommentId:  elem.CommentId,
			Uid:        elem.GetUid(),
			UserName:   elem.GetUserName(),
			ToUid:      elem.GetToUid(),
			ToUserName: elem.GetToUserName(),
		})
	}
	ret := make(map[uint64][]*Comment)
	for _, elem := range commentRsp.List {
		ret[elem.MessageId] = append(ret[elem.MessageId], &Comment{
			ID:        elem.CommentId,
			Comment:   elem.Comment,
			CTime:     common.GetTimeFromTimestamp(elem.GetCtime()),
			MessageId: elem.MessageId,
			ReplyList: commentReplyMap[elem.CommentId],
			Uid:       elem.GetUid(),
			UserName:  elem.GetUserName(),
		})
	}
	return ret
}

type PublishRequest struct {
	Message string `json:"message"`
}
type PublishResponse struct {
	Ret int32  `json:"ret"`
	Msg string `json:"msg"`
	Url string `json:"url"`
}

func PublishMessage(w http.ResponseWriter, r *http.Request) {
	rsp := PublishResponse{}

	// 解析 JSON 参数
	var params PublishRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		rsp.Ret = common.MissingParams
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	sessionInfo, status := checkSession(r)
	if !status {
		log.Log.Errorf("get session err")
		rsp.Ret = common.InvalidSession
		rsp.Url = config.Config.LoginPage
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	u, err := getUserInfoRpc(sessionInfo.UID)
	if err != nil {
		log.Log.Errorf("getUserInfoRpc err")
		rsp.Ret = common.ServerErrCode
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
	}

	message := base64.StdEncoding.EncodeToString([]byte(params.Message))
	err = publishMessage(u.UID, u.Nick, message)
	if err != nil {
		log.Log.Errorf("publishMessage err")
		rsp.Ret = common.ServerErrCode
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
	}

	msg, _ := json.Marshal(rsp)
	_, _ = w.Write(msg)
}

type DeleteRequest struct {
	MessageId uint64 `json:"message_id"`
}
type DeleteResponse struct {
	Ret int32  `json:"ret"`
	Msg string `json:"msg"`
	Url string `json:"url"`
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	rsp := DeleteResponse{}

	// 解析 JSON 参数
	var params DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		rsp.Ret = common.MissingParams
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	sessionInfo, status := checkSession(r)
	if !status {
		log.Log.Errorf("get session err")
		rsp.Ret = common.InvalidSession
		rsp.Url = config.Config.LoginPage
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	err := deleteMessageRpc(sessionInfo.UID, params.MessageId)
	if err != nil {
		log.Log.Errorf("deleteMessageRpc err")
		rsp.Ret = common.ServerErrCode
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
	}

	msg, _ := json.Marshal(rsp)
	_, _ = w.Write(msg)
}

package handleLogic

import (
	"encoding/json"
	"entryTask/common/log"
	"entryTask/httpServer/common"
	"entryTask/httpServer/config"
	"net/http"
)

type MessageInfo struct {
	UID   string `json:"uid"` //用户ID
	Msg   string `json:"msg"`
	Image string `json:"image"`
	ID    uint64 `json:"id"`
	Owner string `json:"owner"`
	CTime uint64 `json:"c_time"`
	MTime uint64 `json:"m_time"`
}

type getMsgListRsp struct {
	Ret      int32          `json:"ret"` //业务返回码
	List     []*MessageInfo `json:"list"`
	Count    int32          `json:count`
	Page     int32          `json:"page"`
	PageSize int32          `json:"page_size"`
}

type getMsgListRequest struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"pageSize"`
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

	msgList, count, err := getMessageListRpc(sessionInfo.UID, req.Page, req.PageSize)
	if err != nil {
		log.Log.Errorf("getMessageListRpc err")
		rsp.Ret = common.ServerErrCode
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
	}

	for _, elem := range msgList {
		rsp.List = append(rsp.List, &MessageInfo{
			UID:   elem.GetUid(),
			ID:    elem.GetId(),
			Msg:   elem.GetMsg(),
			Image: elem.GetImage(),
			Owner: elem.GetOwner(),
			CTime: elem.GetCtime(),
			MTime: elem.GetMtime(),
		})
	}
	rsp.Count = count

	msg, _ := json.Marshal(rsp)
	_, _ = w.Write(msg)
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

	err = publishMessage(u.UID, u.Nick, params.Message)
	if err != nil {
		log.Log.Errorf("publishMessage err")
		rsp.Ret = common.ServerErrCode
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
	}

	msg, _ := json.Marshal(rsp)
	_, _ = w.Write(msg)
}

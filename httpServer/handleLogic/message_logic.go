package handleLogic

import (
	"encoding/json"
	"entryTask/common/log"
	"entryTask/httpServer/common"
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
	Ret  int32          `json:"ret"` //业务返回码
	List []*MessageInfo `json:"list"`
}

func GetMessageList(w http.ResponseWriter, r *http.Request) {
	rsp := getMsgListRsp{Ret: common.SucCode}

	_, status := checkSession(r)
	if !status {
		log.Log.Errorf("get session err")
		rsp.Ret = common.InvalidSession
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	msgList, err := getMessageListRpc("")
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

	msg, _ := json.Marshal(rsp)
	_, _ = w.Write(msg)
}

package handleLogic

import (
	"encoding/json"
	"entryTask/common/log"
	"entryTask/httpServer/common"
	"entryTask/httpServer/config"
	"fmt"
	"html/template"
	"net/http"
)

// userLoginRsp 用户登入返回包
type userLoginRsp struct {
	Ret     int32  `json:"ret"`     //业务返回码
	Uid     string `json:"uid"`     //用户ID
	Nick    string `json:"nick"`    //用户昵称
	Picture string `json:"picture"` //用户头像
	Url     string `json:"url"`     //跳转链接
}

// getUserInfoRsp 获取用户信息返回包
type getUserInfoRsp struct {
	Ret  int32  `json:"ret"`  //业务返回码
	UID  string `json:"uid"`  //用户ID
	Nick string `json:"nick"` //昵称
	Pic  string `json:"pic"`  //头像
	Url  string `json:"url"`  //跳转链接
}

// updateUserInfoRsp 更新用户信息返回包
type updateUserInfoRsp struct {
	Ret int32  `json:"ret"` //业务返回码
	Url string `json:"url"` //跳转链接
}

// LoginPage 登入界面
func LoginPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(config.Config.IndexPage)
	if err != nil {
		log.Log.Errorf("err:%s", err)
	}
	_ = t.Execute(w, nil)
}

// UserLogin 用户登入
func UserLogin(w http.ResponseWriter, r *http.Request) {
	rsp := userLoginRsp{
		Ret: common.SucCode,
	}

	_ = r.ParseForm()
	if !checkParams("uid", r) || !checkParams("pwd", r) {
		rsp.Ret = common.MissingParams
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	pwd := r.Form["pwd"][0]
	uid := r.Form["uid"][0]

	userInfo, err := getUserInfoRpc(uid)
	if err != nil {
		rsp.Ret = common.ServerErrCode
		fmt.Printf("get err:%s   uid:%s\n", err, uid)
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	md5Pwd := encryptedByMD5(uid + pwd)
	if userInfo.PassWord != md5Pwd {
		rsp.Ret = common.WrongAccountInfoCode
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	sessionID := createSessionID(uid)
	cookie := http.Cookie{Name: "sessionID", Value: sessionID, Path: "/", HttpOnly: false, MaxAge: 3600}
	http.SetCookie(w, &cookie)

	//离线处理session
	setSession(sessionID, r, userInfo)

	rsp.Uid = uid
	rsp.Nick = userInfo.Nick
	rsp.Picture = userInfo.Picture
	rsp.Url = config.Config.UserInfoPage
	msg, _ := json.Marshal(rsp)
	_, _ = w.Write(msg)

}

// UserLogout 用户登出
func UserLogout(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "user logout!")
}

// GetUserInfo 获取用户信息
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	rsp := getUserInfoRsp{Ret: common.SucCode}

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
		log.Log.Errorf("get session err")
		rsp.Ret = common.ServerErrCode
		rsp.Url = config.Config.LoginPage
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
	}

	rsp.UID = u.UID
	rsp.Nick = u.Nick
	rsp.Pic = u.Picture

	msg, _ := json.Marshal(rsp)
	_, _ = w.Write(msg)
}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	rsp := updateUserInfoRsp{Ret: common.SucCode}
	sessionInfo, status := checkSession(r)
	if !status {
		log.Log.Errorf("get session err")
		rsp.Ret = common.InvalidSession
		rsp.Url = config.Config.LoginPage
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}
	_ = r.ParseForm()
	userInfo := common.UserInfo{UID: sessionInfo.UID}
	if checkParams("nick", r) {
		userInfo.Nick = r.Form["nick"][0]
	}

	if checkParams("pic", r) {
		userInfo.Picture = r.Form["pic"][0]
	}

	err := updateUserInfoRpc(userInfo)
	if err != nil {
		log.Log.Errorf("update user info err:%s", err)
		rsp.Ret = common.ServerErrCode
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}
	msg, _ := json.Marshal(rsp)
	_, _ = w.Write(msg)

}

// checkParams 检查参数
func checkParams(key string, r *http.Request) bool {
	param, ok := r.Form[key]
	if !ok || len(param) == 0 || len(param[0]) == 0 {
		return false
	}
	return true
}

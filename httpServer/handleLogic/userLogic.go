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
	List []string `json:"list"` //消息列表
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

type LoginParams struct {
	UID string `json:"uid"`
	PWD string `json:"pwd"`
}

// UserLogin 用户登入
func UserLogin(w http.ResponseWriter, r *http.Request) {
	rsp := userLoginRsp{
		Ret: common.SucCode,
	}

	// 解析 JSON 参数
	var params LoginParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		rsp.Ret = common.MissingParams
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	//_ = r.ParseForm()
	pwd := params.PWD
	uid := params.UID
	if pwd == "" || uid == "" {
		rsp.Ret = common.MissingParams
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	userInfo, err := getUserInfoRpc(uid)
	if err != nil {
		rsp.Ret = common.ServerErrCode
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
	http.SetCookie(w, &http.Cookie{Name: "uid", Value: uid, Path: "/", HttpOnly: false, MaxAge: 3600})
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
		fmt.Println("111111")
		rsp.Ret = common.InvalidSession
		rsp.Url = config.Config.LoginPage
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
		return
	}

	u, err := getUserInfoRpc(sessionInfo.UID)
	if err != nil {
		fmt.Println("2222222")
		log.Log.Errorf("get session err")
		rsp.Ret = common.ServerErrCode
		rsp.Url = config.Config.LoginPage
		msg, _ := json.Marshal(rsp)
		_, _ = w.Write(msg)
	}

	rsp.UID = u.UID
	rsp.Nick = u.Nick
	rsp.Pic = u.Picture
	rsp.List = []string{
		"国家金融监督管理总局发布风险提示:防范虚假网络投资理财类诈骗",
		"开辟绿色通道!河南公布“三秋”农机作业便民服务措施",
		"近年来,山东省人大常委会坚持以习近平新时代中国特色社会主义思想为指导,认真贯彻落实习近平法治思想、习近平总书记关于坚持和完善人民代表大会制度的重要思想,坚决落实党对立法工作的全面领导,充分发挥人大主导作用,深入推进科学立法、民主立法、...",
		"月底竣工 8月19日,在山西长铁绿能物流园项目一期工程施工现场,施工人员紧张施工,全速推进铁路敞顶箱装车平台及堆箱场地硬化工程建设进程。据悉,该项目一期工程将于8月底完成。届时,这里将成为自山东港口到长治地区铁矿石铁路运输的“中转站...",
		"中新网郑州9月19日电 (记者 刘鹏)18日下午,国家地热能中心河南分中心与国际地热协会(以下简称IGA)召开座谈会,签署《河南省地热领域战略合作备忘录》",
		"今年,从中央到地方一系列支持民营经济发展的举措鼓点更密、措施更实,民营企业发展信心更强,底气更足。山东是一片民营经济发展的沃土,拿出时不我待、只争朝夕的紧迫感继续聚焦民营经济出实招、下实功,全省民营经济就能稳住好势头、巩固...",
	}
	if u.UID == "9999910" {
		rsp.List[0] = "杨大村罪行：1.胡乱扔袜子。 2.经常不洗头。 3.无恶不作"
		rsp.List[1] = "杨大村语录：1.我比你强. 2.我这不单一. 3.咱俩今天开始要好好吃饭了"
	}

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
		fmt.Printf("update user info err:%s", err)
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

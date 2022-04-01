package handleLogic

import (
	"crypto/md5"
	"encoding/hex"
	"entryTask/common/log"
	"entryTask/httpServer/common"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// encryptedByMD5 MD5加密
func encryptedByMD5(s string) string {
	h := md5.New()
	//uid不会发生变化
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// createSessionID 创建一个会话ID，毫秒时间戳+随机数+用户名的MD5值
func createSessionID(uid string) string {
	timestamp := time.Now().UnixMilli()
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)
	s := fmt.Sprintf("%d.%d.%s", timestamp, randomNum, uid)
	sessionID := encryptedByMD5(s)
	return sessionID
}

// checkSession 检查会话
func checkSession(r *http.Request) (sessionInfo common.SessionInfo, ok bool) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Log.Errorf("get cookie err:%s", err)
		return
	}
	sessionID := cookie.Value
	sessionInfo, err = getSessionInfo(sessionID)
	if err != nil {
		return sessionInfo, false
	}

	arr := strings.Split(r.RemoteAddr, ":")
	if len(arr) != 2 || arr[0] == "" {
		return sessionInfo, false
	}

	if sessionInfo.IP != arr[0] {
		return sessionInfo, false
	}
	// check session表明有用户操作，即刷新session过期时间
	go refreshSessionRpc(sessionID)

	return sessionInfo, true
}

// setSession 设置会话
func setSession(sessionID string, r *http.Request, info common.UserInfo) {
	var ip string
	addr := r.RemoteAddr
	arr := strings.Split(addr, ":")
	if len(arr) == 2 {
		ip = arr[0]
	}
	sessionInfo := common.SessionInfo{
		UID: info.UID,
		//Nick:    info.Nick,
		//Picture: info.Picture,
		IP: ip,
	}
	setSessionRpc(sessionID, sessionInfo)
}

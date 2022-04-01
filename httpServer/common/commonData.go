package common

// UserInfo 用户信息
type UserInfo struct {
	UID      string //用户ID
	Nick     string //昵称
	Picture  string //头像
	PassWord string //密码
}

// SessionInfo 会话信息
type SessionInfo struct {
	UID       string `json:"uid"`       //用户ID
	IP        string `json:"ip"`        //用户IP地址
	SessionID string `json:"sessionID"` //会话ID
}

const (
	SucCode              int32 = 0 //成功码
	MissingParams        int32 = 1 //缺失参数
	WrongAccountInfoCode int32 = 2 //账户信息错误
	ServerErrCode        int32 = 3 //服务内部错误
	InvalidSession       int32 = 4 //无效会话
)

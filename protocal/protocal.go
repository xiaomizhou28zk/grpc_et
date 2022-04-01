package protocal

//获取用户信息

type GetUserInfoRequest struct {
	Uid string
}

type GetUserInfoReply struct {
	Ret  int32
	Uid  string
	Nick string
	Pic  string
	Pwd  string
}

//更新用户信息

type UpdateUserInfoRequest struct {
	Uid  string
	Nick string
	Pic  string
}

type UpdateUserInfoReplay struct {
	Ret int32
}

//获取会话信息

type GetSessionInfoRequest struct {
	SessionID string
}

type GetSessionInfoReply struct {
	Ret         int32
	SessionInfo string
}

//设置会话信息

type SetSessionInfoRequest struct {
	SessionID   string
	SessionInfo string
}

type SetSessionInfoReply struct {
	Ret int32
}

//刷新会话

type RefreshSessionRequest struct {
	SessionID string
}

type RefreshSessionReply struct {
	Ret int32
}

//rpc接口

var GetUserInfo func(in GetUserInfoRequest) (GetUserInfoReply, error)
var UpdateUserInfo func(in UpdateUserInfoRequest) (UpdateUserInfoReplay, error)
var GetSessionInfo func(in GetSessionInfoRequest) (GetSessionInfoReply, error)
var RefreshSession func(in RefreshSessionRequest) (RefreshSessionReply, error)
var SetSessionInfo func(in SetSessionInfoRequest) (SetSessionInfoReply, error)

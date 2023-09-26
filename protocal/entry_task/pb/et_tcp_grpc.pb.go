// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: et_tcp.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EntryTaskClient is the client API for EntryTask service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntryTaskClient interface {
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	GetSessionInfo(ctx context.Context, in *GetSessionInfoRequest, opts ...grpc.CallOption) (*GetSessionInfoResponse, error)
	RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error)
	SetSessionInfo(ctx context.Context, in *SetSessionInfoRequest, opts ...grpc.CallOption) (*SetSessionInfoResponse, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error)
	GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error)
	PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*PublishMessageResponse, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error)
	GetCommentsByMessageIds(ctx context.Context, in *GetCommentsByMessageIdsRequest, opts ...grpc.CallOption) (*GetCommentsByMessageIdsResponse, error)
	GetReplyByCommentIds(ctx context.Context, in *GetReplyByCommentIdsRequest, opts ...grpc.CallOption) (*GetReplyByCommentIdsResponse, error)
}

type entryTaskClient struct {
	cc grpc.ClientConnInterface
}

func NewEntryTaskClient(cc grpc.ClientConnInterface) EntryTaskClient {
	return &entryTaskClient{cc}
}

func (c *entryTaskClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) GetSessionInfo(ctx context.Context, in *GetSessionInfoRequest, opts ...grpc.CallOption) (*GetSessionInfoResponse, error) {
	out := new(GetSessionInfoResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/GetSessionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error) {
	out := new(RefreshSessionResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/RefreshSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) SetSessionInfo(ctx context.Context, in *SetSessionInfoRequest, opts ...grpc.CallOption) (*SetSessionInfoResponse, error) {
	out := new(SetSessionInfoResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/SetSessionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error) {
	out := new(DeleteSessionResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error) {
	out := new(GetMessageListResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/GetMessageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*PublishMessageResponse, error) {
	out := new(PublishMessageResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/PublishMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error) {
	out := new(DeleteMessageResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/DeleteMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) GetCommentsByMessageIds(ctx context.Context, in *GetCommentsByMessageIdsRequest, opts ...grpc.CallOption) (*GetCommentsByMessageIdsResponse, error) {
	out := new(GetCommentsByMessageIdsResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/GetCommentsByMessageIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryTaskClient) GetReplyByCommentIds(ctx context.Context, in *GetReplyByCommentIdsRequest, opts ...grpc.CallOption) (*GetReplyByCommentIdsResponse, error) {
	out := new(GetReplyByCommentIdsResponse)
	err := c.cc.Invoke(ctx, "/entry_task.EntryTask/GetReplyByCommentIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntryTaskServer is the server API for EntryTask service.
// All implementations should embed UnimplementedEntryTaskServer
// for forward compatibility
type EntryTaskServer interface {
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	GetSessionInfo(context.Context, *GetSessionInfoRequest) (*GetSessionInfoResponse, error)
	RefreshSession(context.Context, *RefreshSessionRequest) (*RefreshSessionResponse, error)
	SetSessionInfo(context.Context, *SetSessionInfoRequest) (*SetSessionInfoResponse, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error)
	GetMessageList(context.Context, *GetMessageListRequest) (*GetMessageListResponse, error)
	PublishMessage(context.Context, *PublishMessageRequest) (*PublishMessageResponse, error)
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	GetCommentsByMessageIds(context.Context, *GetCommentsByMessageIdsRequest) (*GetCommentsByMessageIdsResponse, error)
	GetReplyByCommentIds(context.Context, *GetReplyByCommentIdsRequest) (*GetReplyByCommentIdsResponse, error)
}

// UnimplementedEntryTaskServer should be embedded to have forward compatible implementations.
type UnimplementedEntryTaskServer struct {
}

func (UnimplementedEntryTaskServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedEntryTaskServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedEntryTaskServer) GetSessionInfo(context.Context, *GetSessionInfoRequest) (*GetSessionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionInfo not implemented")
}
func (UnimplementedEntryTaskServer) RefreshSession(context.Context, *RefreshSessionRequest) (*RefreshSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshSession not implemented")
}
func (UnimplementedEntryTaskServer) SetSessionInfo(context.Context, *SetSessionInfoRequest) (*SetSessionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSessionInfo not implemented")
}
func (UnimplementedEntryTaskServer) DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedEntryTaskServer) GetMessageList(context.Context, *GetMessageListRequest) (*GetMessageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageList not implemented")
}
func (UnimplementedEntryTaskServer) PublishMessage(context.Context, *PublishMessageRequest) (*PublishMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishMessage not implemented")
}
func (UnimplementedEntryTaskServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedEntryTaskServer) GetCommentsByMessageIds(context.Context, *GetCommentsByMessageIdsRequest) (*GetCommentsByMessageIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByMessageIds not implemented")
}
func (UnimplementedEntryTaskServer) GetReplyByCommentIds(context.Context, *GetReplyByCommentIdsRequest) (*GetReplyByCommentIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplyByCommentIds not implemented")
}

// UnsafeEntryTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntryTaskServer will
// result in compilation errors.
type UnsafeEntryTaskServer interface {
	mustEmbedUnimplementedEntryTaskServer()
}

func RegisterEntryTaskServer(s grpc.ServiceRegistrar, srv EntryTaskServer) {
	s.RegisterService(&EntryTask_ServiceDesc, srv)
}

func _EntryTask_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_GetSessionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).GetSessionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/GetSessionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).GetSessionInfo(ctx, req.(*GetSessionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_RefreshSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).RefreshSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/RefreshSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).RefreshSession(ctx, req.(*RefreshSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_SetSessionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSessionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).SetSessionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/SetSessionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).SetSessionInfo(ctx, req.(*SetSessionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_GetMessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).GetMessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/GetMessageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).GetMessageList(ctx, req.(*GetMessageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_PublishMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).PublishMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/PublishMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).PublishMessage(ctx, req.(*PublishMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/DeleteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_GetCommentsByMessageIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsByMessageIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).GetCommentsByMessageIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/GetCommentsByMessageIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).GetCommentsByMessageIds(ctx, req.(*GetCommentsByMessageIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryTask_GetReplyByCommentIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReplyByCommentIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryTaskServer).GetReplyByCommentIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry_task.EntryTask/GetReplyByCommentIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryTaskServer).GetReplyByCommentIds(ctx, req.(*GetReplyByCommentIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntryTask_ServiceDesc is the grpc.ServiceDesc for EntryTask service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntryTask_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entry_task.EntryTask",
	HandlerType: (*EntryTaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _EntryTask_GetUserInfo_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _EntryTask_UpdateUserInfo_Handler,
		},
		{
			MethodName: "GetSessionInfo",
			Handler:    _EntryTask_GetSessionInfo_Handler,
		},
		{
			MethodName: "RefreshSession",
			Handler:    _EntryTask_RefreshSession_Handler,
		},
		{
			MethodName: "SetSessionInfo",
			Handler:    _EntryTask_SetSessionInfo_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _EntryTask_DeleteSession_Handler,
		},
		{
			MethodName: "GetMessageList",
			Handler:    _EntryTask_GetMessageList_Handler,
		},
		{
			MethodName: "PublishMessage",
			Handler:    _EntryTask_PublishMessage_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _EntryTask_DeleteMessage_Handler,
		},
		{
			MethodName: "GetCommentsByMessageIds",
			Handler:    _EntryTask_GetCommentsByMessageIds_Handler,
		},
		{
			MethodName: "GetReplyByCommentIds",
			Handler:    _EntryTask_GetReplyByCommentIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "et_tcp.proto",
}

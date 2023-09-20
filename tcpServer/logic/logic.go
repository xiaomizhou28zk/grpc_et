package logic

import (
	"entryTask/protocal/entry_task/pb"
)

var _ pb.EntryTaskServer = &Server{}

type Server struct{}

package zrpc

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

type Server struct {
	addr  string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

func (s *Server) Run() {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}

		go func() {
			srvTransport := NewTransport(conn)

			for {

				req, err := srvTransport.Receive()
				if err != nil {
					if err != io.EOF {
						fmt.Println("server   rcv  err:%", err)
					}
					return
				}

				f, ok := s.funcs[req.Name]
				if !ok {
					e := fmt.Sprintf("func %s does not exist", req.Name)
					log.Println(e)
					if err = srvTransport.Send(Data{Name: req.Name, Err: e}); err != nil {
						log.Printf("transport write err: %v\n", err)
					}
					continue
				}

				inArgs := make([]reflect.Value, len(req.Args))
				for i := range req.Args {
					inArgs[i] = reflect.ValueOf(req.Args[i])
				}

				out := f.Call(inArgs)

				outArgs := make([]interface{}, len(out)-1)
				for i := 0; i < len(out)-1; i++ {
					outArgs[i] = out[i].Interface()
				}

				var e string
				if _, ok := out[len(out)-1].Interface().(error); !ok {
					e = ""
				} else {
					e = out[len(out)-1].Interface().(error).Error()
				}

				err = srvTransport.Send(Data{Name: req.Name, Args: outArgs, Err: e})
				if err != nil {
					log.Printf("transport write err: %v\n", err)
				}
			}
		}()
	}
}

func (s *Server) Register(name string, f interface{}) {
	if _, ok := s.funcs[name]; ok {
		return
	}
	s.funcs[name] = reflect.ValueOf(f)
}

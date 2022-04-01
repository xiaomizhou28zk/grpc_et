package common

import (
	"entryTask/httpServer/config"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

type ConnRes interface {
	Close() error
}

type Factory func() (ConnRes, error)

type Conn struct {
	conn ConnRes
	//连接时间
	time time.Time
}

type ConnPool struct {
	mu          sync.Mutex
	conns       chan *Conn
	factory     Factory
	closed      bool
	connTimeOut time.Duration
}

func NewConnPool(factory Factory, cap int, connTimeOut time.Duration) (*ConnPool, error) {
	if cap <= 0 {
		return nil, errors.New("cap < 0")
	}
	if connTimeOut <= 0 {
		return nil, errors.New("connTimeOut < 0")
	}

	cp := &ConnPool{
		mu:          sync.Mutex{},
		conns:       make(chan *Conn, cap),
		factory:     factory,
		closed:      false,
		connTimeOut: connTimeOut,
	}
	for i := 0; i < cap; i++ {
		connRes, err := cp.factory()
		if err != nil {
			cp.Close()
			return nil, errors.New("factory err")
		}
		cp.conns <- &Conn{conn: connRes, time: time.Now()}
	}

	return cp, nil
}

func (cp *ConnPool) Get() (ConnRes, error) {
	if cp.closed {
		return nil, errors.New("pool closed")
	}

	for {
		select {
		//从通道中获取连接资源
		case connRes, ok := <-cp.conns:
			{
				if !ok {
					return nil, errors.New("pool closed")
				}

				//if time.Now().Sub(connRes.time) > cp.connTimeOut {
				//	_ = connRes.conn.Close()
				//	continue
				//}
				return connRes.conn, nil
			}
			/*
				default:
				{
					//如果无法从通道中获取资源，则重新创建一个资源返回
					connRes, err := cp.factory()
					if err != nil {
						return nil, err
					}
					return connRes, nil
				}
			*/
		}
	}
}

func (cp *ConnPool) Put(conn ConnRes) error {
	if cp.closed {
		return errors.New("连接池已关闭")
	}

	select {
	//向通道中加入连接资源
	case cp.conns <- &Conn{conn: conn, time: time.Now()}:
		{
			return nil
		}
	default:
		{
			//如果无法加入，则关闭连接
			_ = conn.Close()
			return errors.New("连接池已满")
		}
	}
}

func (cp *ConnPool) Close() {
	if cp.closed {
		return
	}
	cp.mu.Lock()
	cp.closed = true
	close(cp.conns)
	for conn := range cp.conns {
		_ = conn.conn.Close()
	}
	cp.mu.Unlock()
}

/*
func (cp *ConnPool) len() int {
	return len(cp.conns)
}
*/

var MyPool *ConnPool

func InitPool() {
	addr := fmt.Sprintf("%s:%d", config.Config.RpcServerIP, config.Config.RpcServerPort)
	MyPool, _ = NewConnPool(func() (ConnRes, error) {
		return net.Dial("tcp", addr)
	}, 4, time.Second*time.Duration(config.Config.ConnTimeOut))
}

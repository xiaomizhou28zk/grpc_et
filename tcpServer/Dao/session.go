package Dao

import (
	"entryTask/tcpServer/config"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	netWorkType = "tcp"
)

var RedisClient *redis.Pool

func InitRedisPool() {
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle: 6, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   6,                 //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			c, err := redis.Dial(netWorkType, config.Config.RedisAddr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

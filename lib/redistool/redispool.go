package redistool

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

//addr "host:port"
func NewPool(addr string, passwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if passwd != "" {
				if _, err := c.Do("AUTH", passwd); err != nil {
					c.Close()
					fmt.Printf("cant auth by passwd:%s", passwd)
					return nil, err
				}
			}
			return c, nil
		},
	}
}

func NewConn(addr string, passwd string) (redis.Conn, error) {
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	if passwd != "" {
		if _, err := c.Do("AUTH", passwd); err != nil {
			c.Close()
			fmt.Printf("cant auth by passwd:%s", passwd)
			return nil, err
		}
	}
	return c, nil
}

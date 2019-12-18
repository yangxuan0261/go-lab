package test_redis

import (
	"fmt"
	"reflect"
	"sync"
)

import (
	"github.com/gomodule/redigo/redis"
)

type subscribeFn func(string, []byte)

type Subscriber struct {
	client redis.PubSubConn
	cbMap  map[string][]subscribeFn
	cbMu   sync.Mutex
}

func (s *Subscriber) Connect(addr string) error {
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		return err
	}

	s.client = redis.PubSubConn{Conn: c}
	s.cbMap = make(map[string][]subscribeFn)
	s.listen()
	return nil
}

func (s *Subscriber) listen() {
	go func() {
		defer fmt.Printf("--- listen over\n")
		for {
			switch res := s.client.Receive().(type) { // 阻塞
			case redis.Message:
				for _, fn := range s.cbMap[res.Channel] {
					go fn(res.Channel, res.Data)
				}
			//case redis.Subscription:
			//	fmt.Printf("--- redis.Subscription, res:%+v\n", res)
			case error: // client close
				return
			}
		}
	}()
}

func (s *Subscriber) Close() error {
	return s.client.Close()
}

func (s *Subscriber) Sub(tp string, cb subscribeFn) error {
	err := s.client.Subscribe(tp)
	if err != nil {
		return err
	}

	s.cbMu.Lock()
	defer s.cbMu.Unlock()

	fnArr, _ := s.cbMap[tp]
	s.cbMap[tp] = append(fnArr, cb)
	return nil
}

func (s *Subscriber) Unsub(tp string, cb subscribeFn) {
	s.cbMu.Lock()
	defer s.cbMu.Unlock()

	if fnArr, ok := s.cbMap[tp]; ok {
		for k, v := range fnArr {
			if reflect.ValueOf(v).Pointer() == reflect.ValueOf(cb).Pointer() {
				s.cbMap[tp] = append(fnArr[0:k], fnArr[k+1:]...)
				break
			}
		}
	}

	//// debug dump
	//if fnArr, ok := s.cbMap[tp]; ok {
	//	fmt.Printf("--- tp:%s, len:%d\n", tp, len(fnArr))
	//}
}

package test_redis

import (
	"github.com/gomodule/redigo/redis"
	"reflect"
	"sync"
)

type subFn func(string, []byte)

type subscriber struct {
	client redis.PubSubConn
	cbMap  map[string][]subFn
	cbMu   sync.Mutex
}

func NewSubscriber(addr string) (*subscriber, error) {
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	ins := new(subscriber)
	ins.client = redis.PubSubConn{Conn: c}
	ins.cbMap = make(map[string][]subFn)
	ins.listen()
	return ins, nil
}

func (s *subscriber) listen() {
	go func() {
		for {
			switch res := s.client.Receive().(type) { // 阻塞
			case redis.Message:
				for _, fn := range s.cbMap[res.Channel] {
					go fn(res.Channel, res.Data)
				}
			//case redis.Subscription:
			//	fmt.Printf("--- redis.Subscription, res:%+v\n", res)
			case error: // client close
				//fmt.Printf("--- error:%+v\n", res)
				return
			}
		}
	}()
}

func (s *subscriber) Close() error {
	return s.client.Close()
}

func (s *subscriber) Sub(tp string, cb subFn) error {
	s.cbMu.Lock()
	defer s.cbMu.Unlock()

	fnArr, _ := s.cbMap[tp]
	if len(fnArr) == 0 {
		err := s.client.Subscribe(tp)
		if err != nil {
			return err
		}
	}
	s.cbMap[tp] = append(fnArr, cb)
	return nil
}

// 实测报错: error:ERR only (P)SUBSCRIBE / (P)UNSUBSCRIBE / PING / QUIT allowed in this context
// 也就是 PubSubConn 中只能有 订阅, 不能用于 发布
//func (s *subscriber) Pub(tp string, msg interface{}) (reply interface{}, err error) {
//	return s.client.Conn.Do("PUBLISH", tp, msg)
//}

func (s *subscriber) Unsub(tp string, cb subFn) {
	s.cbMu.Lock()
	defer s.cbMu.Unlock()

	if fnArr, ok := s.cbMap[tp]; ok {
		for k, v := range fnArr {
			if reflect.ValueOf(v).Pointer() == reflect.ValueOf(cb).Pointer() {
				newArr := append(fnArr[0:k], fnArr[k+1:]...)
				s.cbMap[tp] = newArr

				if len(newArr) == 0 {
					s.client.Unsubscribe(tp)
				}
				break
			}
		}
	}
}

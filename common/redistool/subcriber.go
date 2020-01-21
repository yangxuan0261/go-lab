package redistool

import (
	"github.com/gomodule/redigo/redis"
	"reflect"
	"sync"
)

type SubFn func(string, []byte)

type Subscriber struct {
	client redis.PubSubConn
	cbMap  map[string][]SubFn
	cbMu   sync.Mutex
}

func NewSubscriber(addr, pass string) (*Subscriber, error) {
	c, err := NewConn(addr, pass)
	if err != nil {
		return nil, err
	}

	ins := new(Subscriber)
	ins.client = redis.PubSubConn{Conn: c}
	ins.cbMap = make(map[string][]SubFn)
	ins.listen()
	return ins, nil
}

func (s *Subscriber) listen() {
	go func() {
		//defer fmt.Println("--- defer Subscriber listen close")

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

func (s *Subscriber) Close() error {
	return s.client.Close()
}

func (s *Subscriber) Sub(tp string, cb SubFn) error {
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
//func (s *Subscriber) Pub(tp string, msg interface{}) (reply interface{}, err error) {
//	return s.client.Conn.Do("PUBLISH", tp, msg)
//}

func (s *Subscriber) Unsub(tp string, cb SubFn) {
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

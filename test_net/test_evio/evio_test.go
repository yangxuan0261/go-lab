package test_evio

import (
	"fmt"
	"github.com/tidwall/evio"
	"log"
	"strings"
	"sync/atomic"
	"testing"
)

func Test_evioTcp(t *testing.T) {
	port := 5000

	var events evio.Events
	events.NumLoops = 2 // -1: cpu 实际线程数
	events.Serving = func(srv evio.Server) (action evio.Action) {
		log.Printf("--- echo server started on port %d (loops: %d)", port, srv.NumLoops)

		return
	}
	id := uint32(0)
	events.Opened = func(c evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
		log.Printf("--- Opened, addr:%s\n", c.RemoteAddr().String())
		c.SetContext(atomic.AddUint32(&id, 1))
		return
	}
	events.Closed = func(c evio.Conn, err error) (action evio.Action) {
		log.Printf("--- Closed, addr:%s, err:%v\n", c.RemoteAddr().String(), err)
		return
	}
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		log.Printf("--- id:%v, in:%s", c.Context(), strings.TrimSpace(string(in)))
		//time.Sleep(time.Second * 1000) // 如果 NumLoops < 这里 sleep 的数量, 将阻塞掉
		out = in
		return
	}

	log.Fatal(evio.Serve(events, fmt.Sprintf("tcp://:%d", port)))
}

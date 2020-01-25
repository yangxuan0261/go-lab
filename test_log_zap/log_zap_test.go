package log

import (
	syslog "go-lab/test_log_zap/log"
	"testing"

	"go.uber.org/zap"
)

/*
参考:
https://studygolang.com/articles/17394
https://www.jianshu.com/p/b0de3b46e63f
*/

func Test_main(t *testing.T) {
	defer syslog.Access.Sync() // flushing any buffered log entries, 把缓存写入到 output

	syslog.Init("./access.json", "./error.json", 0)
	syslog.Error.Errorf("connect error [SssConnectRes],may be conn exception, id:%d", 123)
	syslog.Access.Info("Msg-666", zap.Any("UID", 123123123), zap.String("CMDNAME", "hello"), zap.Uint32("KEY", 666))
}

/*
2019-10-20T14:11:21.795+0800    test_log_zap/test_log_zap.go:13 connect error [SssConnectRes],may be conn exception, id:123
{"L":"info","T":"2019-10-20T14:11:21.811+0800","C":"test_log_zap/test_log_zap.go:14","M":"Msg-666","UID":123123123,"CMDNAME":"hello","KEY":666}
*/

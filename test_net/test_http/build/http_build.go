package main

import (
	syslog "GoLab/test_log_zap/log"
	"io/ioutil"
	"net/http"
)

func SayWorld(w http.ResponseWriter, req *http.Request) {
	ck := req.Header.Get("ccc") // 获取 token 之类的数据
	syslog.Access.Sugar().Infof("--- Cookie ccc:%+v", ck)

	reqBytes, _ := ioutil.ReadAll(req.Body)
	syslog.Access.Sugar().Infof("req body:%s, path:%s", string(reqBytes), req.URL.Path)

	w.Write([]byte("Hello world"))
}

func main() {
	syslog.Init("./temp_access.json", "./temp_error.json", 1)

	http.HandleFunc("/world", SayWorld)
	http.ListenAndServe(":8001", nil)
}

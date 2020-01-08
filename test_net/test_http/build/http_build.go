package main

import (
	syslog "GoLab/test_log_zap/log"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func Test1ue4(ctx *fasthttp.RequestCtx) {
	fmt.Printf("--- Test1ue4\n")
	postBody := ctx.PostBody()
	fmt.Fprint(ctx, postBody)
}

func Test2ue4(ctx *fasthttp.RequestCtx) {
	fmt.Printf("--- Test1ue4\n")
	//fmt.Fprint(ctx, "--- post ret abc:"+string(postBody))
}

func main() {
	syslog.Init("./temp_access.json", "./temp_error.json", 1)

	router := fasthttprouter.New()
	router.GET("/test1ue4", Test1ue4)
	router.POST("/test2ue4", Test2ue4)
	log.Fatal(fasthttp.ListenAndServe(":8002", router.Handler))
}

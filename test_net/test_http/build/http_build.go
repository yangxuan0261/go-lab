package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	syslog "go-lab/test_log_zap/log"
	"log"
)

func hello(ctx *fasthttp.RequestCtx) {
	fmt.Printf("--- hello\n")
	fmt.Fprint(ctx, "world")
}

func Test2ue4(ctx *fasthttp.RequestCtx) {
	postBody := ctx.PostBody()
	fmt.Printf("--- Test2ue4, len:(%d)\n", len(postBody))
	//fmt.Fprint(ctx, postBody)
	ctx.Response.SetBody(postBody)
}

func main() {
	syslog.Init("./temp_access.json", "./temp_error.json", 1)

	router := fasthttprouter.New()
	router.GET("/hello", hello)
	router.POST("/test2ue4", Test2ue4)
	fmt.Printf("--- fasthttp ListenAndServe: 8002\n")
	log.Fatal(fasthttp.ListenAndServe(":8002", router.Handler))
}

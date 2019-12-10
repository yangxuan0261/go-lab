package test_http

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// 参考官方示例: https://github.com/buaazp/fasthttprouter/tree/master/examples/auth

// basicAuth returns the username and password provided in the request's
// Authorization header, if the request uses HTTP Basic Authentication.
// See RFC 2617, Section 2.
func basicAuth(ctx *fasthttp.RequestCtx) (username, password string, ok bool) {
	auth := ctx.Request.Header.Peek("Authorization")
	fmt.Printf("--- Authorization:%s\n", string(auth))

	if auth == nil {
		return
	}
	return parseBasicAuth(string(auth))
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}

	enStr := auth[len(prefix):]
	fmt.Printf("--- enStr:%v\n", enStr)
	c, err := base64.StdEncoding.DecodeString(enStr)
	if err != nil {
		return
	}

	// 解码后根据自己的规则解出对应的信息
	cs := string(c)
	fmt.Printf("--- cs:%v\n", cs)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// BasicAuth is the basic auth handler
func BasicAuth(h fasthttp.RequestHandler, requiredUser, requiredPassword string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := basicAuth(ctx)

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(ctx)
			return
		}

		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
		ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
	}
}

// Index is the index handler
func Index222(ctx *fasthttp.RequestCtx) {
	uri := ctx.URI().String()
	fmt.Printf("--- uri 111:%s\n", uri) // http://localhost:8001/asd, 全路径
	uri = string(ctx.RequestURI())
	fmt.Printf("--- uri 222:%s\n", uri) // /asd, 相对路径

	fmt.Fprintf(ctx, "Not protected!\n")
}

// Protected is the Protected handler
func Protected(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Auth ok, hello world!\n")
}

func Test_SrvFasthttpAuth(t *testing.T) {
	user := "Aladdin"
	pass := "open sesame"

	router := fasthttprouter.New()
	router.GET("/asd", Index222)
	router.POST("/protected", BasicAuth(Protected, user, pass)) // hook, 合法才执行 Protected

	log.Fatal(fasthttp.ListenAndServe(":8001", router.Handler))
}

func Test_parseBasicAuth(t *testing.T) {
	authStr := "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ=="
	user, password, hasAuth := parseBasicAuth(authStr)
	fmt.Println("--- result:", user, password, hasAuth) // Aladdin open sesame true
}

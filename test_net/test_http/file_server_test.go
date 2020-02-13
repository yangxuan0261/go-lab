package test_http

import (
	"net/http"
	"testing"
)

const kPath = "F:/"

func Test_fileSrv(t *testing.T) {
	// http://wiki.jikexueyuan.com/project/magical-go/advanced-application.html
	h := http.FileServer(http.Dir(kPath))
	http.ListenAndServe(":8888", h)
}

// 参考: https://segmentfault.com/a/1190000016086653
//func Test_fileSrvPrefix(t *testing.T) {
//	fs := http.FileServer(http.Dir(kPath))
//	http.Handle("/vm_share/", http.StripPrefix("/vm_share/", fs))
//	err := http.ListenAndServe(":8888", nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

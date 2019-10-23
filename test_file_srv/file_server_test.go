package test_file_srv

import "net/http"

func main() {
	// http://wiki.jikexueyuan.com/project/magical-go/advanced-application.html
	h := http.FileServer(http.Dir("F:/a_downimages"))
	http.ListenAndServe(":8888", h)
}

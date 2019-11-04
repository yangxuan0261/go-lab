package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// 参考: https://cloud.tencent.com/developer/article/1388556
// TODO: 未测试

var (
	server   *http.Server
	listener net.Listener
	child    = flag.Bool("child", false, "")
)

func init() {
	updatePidFile()
}

func updatePidFile() {
	sPid := fmt.Sprint(os.Getpid())
	tmpDir := os.TempDir()
	if err := procExsit(tmpDir); err != nil {
		fmt.Printf("pid file exists, update\n")
	} else {
		fmt.Printf("pid file NOT exists, create\n")
	}
	pidFile, _ := os.Create(tmpDir + "/gracefulRestart.pid")
	defer pidFile.Close()
	pidFile.WriteString(sPid)
}

// 判断进程是否启动
func procExsit(tmpDir string) (err error) {
	pidFile, err := os.Open(tmpDir + "/gracefulRestart.pid")
	defer pidFile.Close()
	if err != nil {
		return
	}

	filePid, err := ioutil.ReadAll(pidFile)
	if err != nil {
		return
	}
	pidStr := fmt.Sprintf("%s", filePid)
	pid, _ := strconv.Atoi(pidStr)
	if _, err := os.FindProcess(pid); err != nil {
		fmt.Printf("Failed to find process: %v\n", err)
		return
	}

	return
}

func main() {
	flag.Parse()

	// 启动监听
	http.HandleFunc("/hello", HelloHandler)
	server = &http.Server{Addr: ":8081"}

	var err error
	if *child {
		fmt.Println("In Child, Listening...")

		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		fmt.Println("In Father, Listening...")

		listener, err = net.Listen("tcp", server.Addr)
	}
	if err != nil {
		fmt.Printf("Listening failed: %v\n", err)
		return
	}

	// 单独go程启动server
	go func() {
		err = server.Serve(listener)
		if err != nil {
			fmt.Printf("server.Serve failed: %v\n", err)
		}
	}()

	//监听系统信号
	singalHandler()
	fmt.Printf("singalHandler end\n")

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(20 * time.Second)
	for i := 0; i < 20; i++ {
		log.Printf("working %v\n", i)
		time.Sleep(1 * time.Second)
	}
	w.Write([]byte("world233333!!!!"))
}

func singalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		sig := <-ch
		fmt.Printf("signal: %v\n", sig)

		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Printf("stop")
			signal.Stop(ch)
			server.Shutdown(ctx)
			fmt.Printf("graceful shutdown\n")
			return
		case syscall.SIGHUP:
			// reload
			log.Printf("restart")
			err := restart()
			if err != nil {
				fmt.Printf("graceful restart failed: %v\n", err)
			}
			//更新当前pidfile
			updatePidFile()
			server.Shutdown(ctx)
			fmt.Printf("graceful reload\n")
			return
		}
	}
}

func restart() error {
	tl, ok := listener.(*net.TCPListener)
	if !ok {
		return fmt.Errorf("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}

	args := []string{"-child"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{f}
	return cmd.Start()
}

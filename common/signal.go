package common

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal() os.Signal {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		sig := <-sig
		return sig
	}
}

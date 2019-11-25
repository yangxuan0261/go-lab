package test_sync

import (
	"fmt"
	"sync"
	"testing"
)

func Test_once(t *testing.T) {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

// 单例
type CDog struct {
}

var (
	instance *CDog
	once     sync.Once
)

func Instance() *CDog {
	once.Do(func() {
		instance = &CDog{}
	})
	return instance
}

func Test_singleton(t *testing.T) {
	ins := Instance()
	_ = ins
}

package test_sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// http://www.iikira.com/2017/12/08/golang-note-3-sync-atomic/

func Test_atomic(t *testing.T) {
	var (
		wg     sync.WaitGroup
		nA, nB int64
	)
	wg.Add(2000)
	for i := 0; i < 1000; i++ {
		go func() {
			nA++
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&nB, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(nA, nB)
}

package test_redis

// Redis 分布式锁, 参考: https://blog.didiyun.com/index.php/2019/01/14/redis-3/

import (
	"GoLab/common/redistool"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"sync"
	"testing"
	"time"
)

/*
Redis在 2.6.12 版本开始，为 SET 命令增加了一系列选项：
SET 命令的天然原子性完全可以取代 SETNX 和 EXPIRE 命令
*/
func tryLock(pool *redis.Pool, lockKey string, ex uint, retry int) error {
	conn := pool.Get()
	defer conn.Close()

	if retry <= 0 {
		retry = 10
	}

	ts := time.Now() // 值随用一个随机值
	for i := 1; i <= retry; i++ {
		if i > 1 { // 睡眠一下, 因为其他线程释放锁需要时间
			time.Sleep(time.Second)
		}
		v, err := conn.Do("SET", lockKey, ts, "EX", retry, "NX")
		if err == nil {
			if v == nil {
				fmt.Println("get lock failed, retry times:", i)
			} else {
				fmt.Println("get lock success")
				break
			}
		} else {
			fmt.Println("get lock failed with err:", err)
		}
		if i >= retry {
			err = errors.New("get lock failed with max retry times.")
			return err
		}
	}
	return nil
}

/*
 */
func unLock(pool *redis.Pool, lockKey string) error {
	conn := pool.Get()
	defer conn.Close()

	v, err := redis.Bool(conn.Do("DEL", lockKey))
	if err == nil {
		if v {
			fmt.Println("unLock success")
		} else {
			fmt.Println("unLock failed")
			return fmt.Errorf("unLock failed")
		}
	} else {
		fmt.Println("unLock failed, err:", err)
		return err
	}

	return nil
}

func Test_dist_lock(t *testing.T) {
	var wg sync.WaitGroup

	const RedisAddr = "192.168.2.233:7379"

	key := "lock:wolegequ"

	pool := redistool.NewPool(RedisAddr, "")

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) { // 并发
			defer wg.Done()
			time.Sleep(time.Second)

			// tryLock
			err := tryLock(pool, key, 10, 10)
			if err != nil {
				fmt.Printf("--- worker[%d] get lock failed:%v\n", id, err)
				return
			}
			fmt.Printf("--- worker[%d] get lock success\n", id)

			// sleep for random, 模拟业务逻辑, 比如: 读取数据库, 更新缓存 操作
			sec := 2
			time.Sleep(time.Duration(sec) * time.Second)
			fmt.Printf("--- worker[%d] hold lock for %ds\n", id, sec)

			// unLock
			err = unLock(pool, key)
			if err != nil {
				fmt.Printf("--- worker[%d] unlock failed:%v\n", id, err)
			}
			fmt.Printf("--- worker[%d] done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("------------ Test_dist_lock is done!")
}

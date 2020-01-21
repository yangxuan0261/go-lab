package test_redis

import (
	"fmt"
)

// Redis 缓存击穿 解决方案 伪代码

const KKeyPrefix = "lock:"

type UserInfo struct {
}

// 分布式锁
func tryLock01(key string) bool {
	uuid := "asd"
	_, err := conn.Do("SET", KKeyPrefix+key, uuid, "NX", "PX", 10000) // 60s 过期时间, 一定要设置过期时间, 否则线程挂掉就不能解锁了
	_ = err
	return false
}

func redisLock(key string) {
	fmt.Printf("--- redisLock\n")
}

func redisUnlock(key string) {
	fmt.Printf("--- redisUnlock\n")
	conn.Do("DEL", KKeyPrefix+key)
}

func redisGetUserInfo(uid string) *UserInfo {

	return nil
}

func redisUpdateUserInfo(info *UserInfo) error {

	return nil
}

func mysqlGetUserInfo(uid string) *UserInfo {

	return nil
}

// 业务层接口
func getUserInfo(uid string) *UserInfo {
	// 1. 查询 Redis
	info := redisGetUserInfo(uid)
	if info != nil {
		return info // 命中缓存, 直接返回
	}

	redisLock(uid)
	defer redisUnlock(uid)

	// 2. 再次查询 Redis, 因为其他线程可以阻塞到 lock 中, 但已经有一个线程已经更新了缓存, 所以再次尝试命中缓存
	info = redisGetUserInfo(uid)
	if info != nil {
		return info // 命中缓存, 直接返回
	}

	// 3. 查询 数据库
	info = mysqlGetUserInfo(uid)
	if info != nil {
		err := redisUpdateUserInfo(info) // 4. 更新缓存
		if err != nil {
			fmt.Printf("--- update redis err:%v\n", err)
		}
		return info
	}
	return nil
}

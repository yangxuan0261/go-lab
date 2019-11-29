package test_redis

import (
	syslog "GoLab/test_log_zap/log"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

// 参考:
// - https://blog.csdn.net/wangshubo1989/article/details/75050024
// - https://www.jianshu.com/p/89ca34b84101
// - https://blog.csdn.net/weixin_37696997/article/details/78634145

/*
常见命令
hset(key, field, value)：向名称为key的hash中添加元素field
hget(key, field)：返回名称为key的hash中field对应的value
hmget(key, (fields))：返回名称为key的hash中field i对应的value
hmset(key, (fields))：向名称为key的hash中添加元素field
hincrby(key, field, integer)：将名称为key的hash中field的value增加integer
hexists(key, field)：名称为key的hash中是否存在键为field的域
hdel(key, field)：删除名称为key的hash中键为field的域
hlen(key)：返回名称为key的hash中元素个数
hkeys(key)：返回名称为key的hash中所有键
hvals(key)：返回名称为key的hash中所有键对应的value
hgetall(key)：返回名称为key的hash中所有的键（field）及其对应的value
*/

var (
	conn redis.Conn
)

func init() {
	var err error
	conn, err = redis.Dial("tcp", "113.102.163.179:7379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	// defer conn.Close() // 用完要关掉
}

// 构造一个错误检查函数
func errCheck(tp string, err error) {
	if err != nil {
		fmt.Printf("--- [%s] sorry,has some error for %+v.\r\n", tp, err)
		os.Exit(-1)
	}
}

func Test_pool(t *testing.T) {
	addr := "113.102.163.179:7379"
	rp := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			//if passwd != "" { // 密码认证
			//	if _, err := c.Do("AUTH", passwd); err != nil {
			//		c.Close()
			//		fmt.Println("cant auth by passwd:%s", passwd)
			//		return nil, err
			//	}
			//}
			fmt.Printf("--- connect success, addr:%s\n", addr)
			return c, nil
		},
	}

	conn2 := rp.Get()
	defer conn2.Close() // 用完要关掉

	_, err := conn2.Do("PING") // 测试是否能 ping 通
	if err != nil {
		syslog.Error.Errorf("connect redis error:%s", err)
		return
	}

	_, err = conn2.Do("SET", "mykey2", "superWang2")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(conn2.Do("GET", "mykey2"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey2: %v \n", username)
	}
}
/*
--- connect success, addr:113.102.163.179:7379
Get mykey2: superWang2
*/

func Test_setget(t *testing.T) {
	_, err := conn.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(conn.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
// Get mykey: superWang

// hmset 和 hgetall 命令的使用：
func Test_hmsethgetall(t *testing.T) {
	//构造实际场景的hash结构体
	var p1, p2 struct {
		Description string `redis:"description"`
		Url         string `redis:"url"`
		Author      string `redis:"author"`
	}

	p1.Description = "my blog"
	p1.Url = "http://xxbandy.github.io"
	p1.Author = "bgbiao"

	_, hmsetErr := conn.Do("hmset", redis.Args{}.Add("hao123").AddFlat(&p1)...)
	errCheck("hmset", hmsetErr)

	m := map[string]string{
		"description": "oschina",
		"url":         "http://my.oschina.net/myblog",
		"author":      "xxbandy",
	}

	_, hmset1Err := conn.Do("hmset", redis.Args{}.Add("hao").AddFlat(m)...)
	errCheck("hmset1", hmset1Err)

	for _, key := range []string{"hao123", "hao"} {
		v, err := redis.Values(conn.Do("hgetall", key))
		errCheck("hmgetV", err)
		//等同于hgetall的输出类型，输出字符串为k/v类型
		//hashV,_ := redis.StringMap(c.Do("hgetall",key))
		//fmt.Println(hashV)
		//等同于hmget 的输出类型，输出字符串到一个字符串列表
		hashV2, _ := redis.Strings(conn.Do("hmget", key, "description", "url", "author"))
		for _, hashv := range hashV2 {
			fmt.Println(hashv)
		}
		if err := redis.ScanStruct(v, &p2); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("--- p2:%+v\n", p2)
	}
}
/*
my blog
http://xxbandy.github.io
bgbiao
--- p2:{Description:my blog Url:http://xxbandy.github.io Author:bgbiao}
oschina
http://my.oschina.net/myblog
xxbandy
--- p2:{Description:oschina Url:http://my.oschina.net/myblog Author:xxbandy}
*/

// hset 和 hget 的使用：
func Test_hsethget(t *testing.T) {
	// core functions
	_, err := conn.Do("hset", "books", "name", "golang")
	errCheck("Test_hsethget", err)

	if r, err := redis.String(conn.Do("hget", "books", "name")); err == nil {
		fmt.Println("book name:", r)
	}
}
// book name: golang



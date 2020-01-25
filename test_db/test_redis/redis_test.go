package test_redis

import (
	syslog "go_lab/test_log_zap/log"
	"encoding/json"
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
// - https://juejin.im/post/5d9ea1bcf265da5bab5bc6fb

var (
	conn redis.Conn
)

func init() {
	var err error
	passwd := ""
	conn, err = redis.Dial("tcp", "192.168.2.233:7379")
	if err != nil {
		panic(err)
	}

	if passwd != "" {
		if res, err := conn.Do("AUTH", passwd); err != nil {
			conn.Close()
			panic(err)
		} else {
			fmt.Printf("--- AUTH, res%v, err:%v\n", res, err)
		}
	}

	reply, rerr := conn.Do("PING")
	fmt.Printf("--- PING, reply%v, rerr:%v\n", reply, rerr)
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
	addr := "127.0.0.1:7379"
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

	_, err = conn2.Do("SET", "mykey3", "superWang2")
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

// 参考 key 命令: https://www.runoob.com/redis/redis-keys.html
/*
[key] = xxx
*/
func Test_setget(t *testing.T) {
	reply, err := conn.Do("SET", "mykey", "superWang")
	fmt.Printf("--- SET, reply:%v, err:%v\n", reply, err) // reply:OK, err:<nil>

	username, err := redis.String(conn.Do("GET", "mykey"))
	fmt.Printf("--- GET, username:%v, err:%v\n", username, err) // username:superWang, err:<nil>

	reply, err = conn.Do("SET", "mykey", "superWang", "NX") // reply:<nil>, err:<nil>, NX 是指只有不存在这个 key 的情况下, 才会操作成功, 所以现在是操作失败的, 返回 nil
	fmt.Printf("--- SET, reply:%v, err:%v\n", reply, err)

	reply, err = conn.Do("SET", "mykey", "superWangXX", "XX") // eply:OK, err:<nil>, XX 与 NX 相反, 是指只有存在这个 key 的情况下, 才会操作成功
	fmt.Printf("--- SET, reply:%v, err:%v\n", reply, err)
}

func printKey(key string) {
	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- GET, key:%s, val:%v\n", key, val)
}

func Test_setNX(t *testing.T) {
	key := "wilker:005"
	reply, err := redis.String(conn.Do("SET", key, "yangx", "NX"))
	fmt.Printf("--- SET, reply:%v, err:%v\n", reply, err)
	// 1. 已存在的情况: --- SET, reply:, err:redigo: nil returned (也就是 redis.ErrNil)
	// 2. 不存在的情况: --- SET, reply:OK, err:<nil>
	if err == redis.ErrNil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	printKey(key)
}

func Test_delKey01(t *testing.T) {
	key := "wilker:004"
	_, err := redis.Bool(conn.Do("DEL", key))
	if err != nil {
		fmt.Println("DEL failed, err:", err)
	}
}

func Test_execLua(t *testing.T) {
	key := "wilker:005"
	val := "yangx"
	var delScript = redis.NewScript(1, `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`)

	reply, err := delScript.Do(conn, key, val) // 删除失败返回 0
	fmt.Printf("--- reply:%v, err:%v\n", reply, err)
}

// Get mykey: superWang

// hmset 和 hgetall 命令的使用：
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
func Test_hmsethgetall(t *testing.T) {
	// 方式一, 结构体
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

	// 方式二, map 映射
	m := map[string]string{
		"description": "oschina",
		"url":         "http://my.oschina.net/myblog",
		"author":      "xxbandy",
	}

	_, hmset1Err := conn.Do("hmset", redis.Args{}.Add("hao").AddFlat(m)...)
	errCheck("hmset1", hmset1Err)

	for _, key := range []string{"hao123", "hao"} {
		println("key:", key)
		// 返回每个 key 对应的 value 的数组
		hashV2, _ := redis.Strings(conn.Do("hmget", key, "description", "url", "author"))
		for k, hashv := range hashV2 {
			fmt.Println("--- hashv, kv:", k, hashv)
		}

		// 映射到一个 struct 中
		v, err := redis.Values(conn.Do("hgetall", key))
		errCheck("hmgetV", err)
		if err := redis.ScanStruct(v, &p2); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("--- p2:%+v\n", p2)
		println()
	}
}

/*
key: hao123
--- hashv, kv: 0 my blog
--- hashv, kv: 1 http://xxbandy.github.io
--- hashv, kv: 2 bgbiao
--- p2:{Description:my blog Url:http://xxbandy.github.io Author:bgbiao}

key: hao
--- hashv, kv: 0 oschina
--- hashv, kv: 1 http://my.oschina.net/myblog
--- hashv, kv: 2 xxbandy
--- p2:{Description:oschina Url:http://my.oschina.net/myblog Author:xxbandy}

*/

// hset 和 hget 的使用：
/*
[key] = map[string]xxx
*/
func Test_hsethget(t *testing.T) {
	// core functions
	_, err := conn.Do("hset", "books", "name", "golang")
	errCheck("Test_hsethget", err)

	if r, err := redis.String(conn.Do("hget", "books", "name")); err == nil {
		fmt.Println("book name:", r)
	}
}

// book name: golang
/*
[key] = []byte
*/
func Test_json(t *testing.T) {
	var err error
	var reply interface{}
	// 写入数据
	imap := map[string]string{"name": "zhang", "sex": "男"}
	// 序列化json数据
	value, _ := json.Marshal(imap)
	reply, err = conn.Do("SET", "jsonkey", value)
	fmt.Printf("--- SETNX, reply:%v, err:%v\n", reply, err)

	// 读取数据
	var imapGet map[string]string
	valueGet, err := redis.Bytes(conn.Do("GET", "jsonkey"))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}

	fmt.Println("imapGet", imapGet)
	fmt.Println("imapGet", imapGet["name"])
	fmt.Println("imapGet", imapGet["sex"])
	/*
		imapGet map[name:zhang sex:男]
		imapGet zhang
		imapGet 男
	*/
}

func Test_exist(t *testing.T) {
	//b, err := redis.Bool(conn.Do("EXISTS", "runoobkey"))
	b, err := redis.Bool(conn.Do("EXISTS", "jsonkey"))
	if err != nil {
		fmt.Println("--- error:", err)
	} else {
		fmt.Printf("--- exists:%v \n", b)
	}
}

func Test_Del(t *testing.T) {
	reply, rerr := conn.Do("SET", "delkey", "hello")
	fmt.Printf("--- set reply:%v, rerr:%v\n", reply, rerr)

	b, berr := redis.Bool(conn.Do("EXISTS", "delkey"))
	fmt.Printf("--- b:%v, berr:%v\n", b, berr)

	// 删除key
	reply, rerr = conn.Do("DEL", "delkey")
	fmt.Printf("--- del reply:%v, rerr:%v\n", reply, rerr)

	b, berr = redis.Bool(conn.Do("EXISTS", "delkey"))
	fmt.Printf("--- b:%v, berr:%v\n", b, berr)
}

package test_encrypt

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var tokenRand = rand.New(rand.NewSource(time.Now().UnixNano()))

var strA []byte = []byte("abcdefghijklmnopqrstuvwxyz")
var strB []byte = []byte("abcdefghijklmnopqrstuvwxyz0123456789")

func GetRandString(haveNum bool, length int) string {
	var str []byte = strA
	if haveNum {
		str = strB
	}
	var result bytes.Buffer
	for i := 0; i < length; i++ {
		tmp := str[tokenRand.Intn(len(str))]
		result.WriteByte(tmp)
	}
	return result.String()
}

func GenTokenByKey(key string) string {
	var result []byte
	result = append(result, []byte(key)...)
	str := GetRandString(true, 5)
	result = append(result, []byte(str)...)

	t := time.Now().UnixNano()
	var buff = make([]byte, 8)
	binary.BigEndian.PutUint64(buff, uint64(t))
	result = append(result, buff...)

	h := sha256.New()
	h.Write(result)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func GenTokenByKey32(key string) string {
	var result []byte
	result = append(result, []byte(key)...)
	str := GetRandString(true, 5)
	result = append(result, []byte(str)...)

	t := time.Now().UnixNano()
	var buff = make([]byte, 8)
	binary.BigEndian.PutUint64(buff, uint64(t))
	result = append(result, buff...)

	h := md5.New()
	h.Write(result)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Test_token_01(t *testing.T) {
	t1 := GenTokenByKey("hello")
	fmt.Println("--- t1:", t1)
	t1 = GenTokenByKey("hello")
	fmt.Println("--- t1:", t1)

	t2 := GenTokenByKey32("hello")
	fmt.Println("--- t2:", t2)
	t2 = GenTokenByKey32("hello")
	fmt.Println("--- t2:", t2)
}

package test_encrypt

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

func Test_base64_01(t *testing.T) {
	input := []byte("hello golang base64 快乐编程http://www.01happy.com +~")

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Printf("--- StdEncoding.EncodeToString:%v\n", encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("--- StdEncoding.DecodeString:%v\n", string(decodeBytes))

	fmt.Println()

	// 如果要用在url中，需要使用URLEncoding
	uEnc := base64.URLEncoding.EncodeToString(input)
	fmt.Println(uEnc)
	fmt.Printf("--- URLEncoding.EncodeToString:%v\n", uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("--- URLEncoding.DecodeString:%v\n", string(uDec))

}

func Test_base64_02(t *testing.T) {

}

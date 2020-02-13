package test_encrypt

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var rander = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	NumStr     = "0123456789"
	CharUpStr  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharLowStr = "abcdefghijklmnopqrstuvwxyz"
	SpecStr    = "~!@#$%^&*()_+-=?"
)

var StrArr1 = [][]byte{[]byte(NumStr), []byte(CharUpStr), []byte(CharLowStr), []byte(SpecStr)}
var StrArr2 = [][]byte{[]byte(NumStr), []byte(CharUpStr), []byte(CharLowStr)}
var StrArr3 = [][]byte{[]byte(CharUpStr), []byte(CharLowStr)}
var StrArr4 = [][]byte{[]byte(NumStr), []byte(CharUpStr)}
var StrArr5 = [][]byte{[]byte(NumStr), []byte(CharLowStr)}

func genPass(num int) string {

	var dstBts [][]byte
	dstBts = StrArr1
	//dstBts = StrArr2
	//dstBts = StrArr3
	//dstBts = StrArr4
	//dstBts = StrArr5

	getByteFn := func(bts [][]byte) byte {
		dstByte := bts[rander.Intn(len(bts))]
		return dstByte[rander.Intn(len(dstByte))]
	}

	newPass := make([]byte, num)
	for i := 0; i < num; i++ {
		newPass[i] = getByteFn(dstBts)
	}

	return string(newPass)
}

func Test_genPass(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println("--- pass:", genPass(10))
	}
}

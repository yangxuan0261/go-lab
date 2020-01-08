package test_pkg

import (
	randr "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	randf "math/rand"
	"testing"
	"time"
)

func init() {
	randf.Seed(time.Now().Unix())
}

func Test_randfAB(t *testing.T) {
	randomf := func(min, max int) int {
		return randf.Intn(max-min) + min // 伪随机, 需要依赖 randf.Seed(time.Now().Unix())
	}

	myrand := randomf(5, 20)
	fmt.Printf("--- val 1:%d\n", myrand)
	myrand2 := randomf(5, 20)
	fmt.Printf("--- val 2:%d\n", myrand2)

}

func Test_randrAB(t *testing.T) {
	randomr := func(min, max int) int {
		result, _ := randr.Int(randr.Reader, big.NewInt(int64(max-min)))
		return int(result.Int64()) + min
	}

	myrand := randomr(5, 20)
	fmt.Printf("--- val 1:%d\n", myrand)
	myrand2 := randomr(5, 20)
	fmt.Printf("--- val 2:%d\n", myrand2)

	for i := 0; i < 3; i++ {
		result, _ := randr.Int(randr.Reader, big.NewInt(5)) //
		fmt.Printf("--- result 1:%d\n", result)
		fmt.Printf("--- result 2:%d\n", result.Uint64())
	}
}

// 真随机 - https://golangnote.com/topic/235.html
func Test_randReal1(t *testing.T) {
	var n int32
	binary.Read(randr.Reader, binary.LittleEndian, &n)
	fmt.Printf("--- result:%d\n", n)

}

func Test_randReal2(t *testing.T) {
	// 不需要依赖 randf.Seed(time.Now().Unix())
	var tokenRand = randf.New(randf.NewSource(time.Now().UnixNano()))

	randomr := func(min, max int) int {
		return tokenRand.Intn(max-min) + min
	}

	myrand3 := randomr(10, 20)
	myrand4 := randomr(10, 20)
	fmt.Printf("--- val 3:%d\n", myrand3)
	fmt.Printf("--- val 4:%d\n", myrand4)
}

package test_bloomfilter

import (
	"fmt"
	"github.com/steakknife/bloomfilter"
	"hash/fnv"
	"testing"
)

/*
布隆过滤器
参考:
- https://github.com/steakknife/bloomfilter
*/

func Test_bf01(t *testing.T) {
	const (
		maxElements = 100000
		probCollide = 0.0000001
	)

	bf, err := bloomfilter.NewOptimal(maxElements, probCollide)
	if err != nil {
		panic(err)
	}

	aa := fnv.New64()
	aa.Write([]byte("aa"))
	bf.Add(aa) // 加入到 布隆过滤器
	fmt.Printf("--- 111:%t\n", bf.Contains(aa))
	if bf.Contains(aa) { // true 是可能存在, false 一定不存在
		// whatever
	}

	bb := fnv.New64()
	bb.Write([]byte("bb"))
	fmt.Printf("--- 222:%t\n", bf.Contains(bb))
	if bf.Contains(bb) {
		panic("This should never happen")
	}

	println()
	file := "./temp_1.bf.gz"
	_, err = bf.WriteFile(file) // saves this BF to a file
	if err != nil {
		panic(err)
	}

	bf2, _, err := bloomfilter.ReadFile(file) // read the BF to another var
	if err != nil {
		panic(err)
	}
	cc := fnv.New64()
	cc.Write([]byte("aa"))
	fmt.Printf("--- 333:%t\n", bf2.Contains(cc))
}

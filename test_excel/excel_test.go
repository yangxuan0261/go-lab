package test_excel

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-lab/test_excel/gen"
	"io/ioutil"
	"testing"
)

// gen 里面的数据由 xls_deploy_tool.py 工具生成

func Test_excel(t *testing.T) {
	dataArr := &datacfg.SkinArray{}
	_ = dataArr

	file := "./gen/cfg_skin.bytes"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = proto.Unmarshal(data, dataArr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("--- success, %+v\n", dataArr)
	for key, value := range dataArr.Items {
		fmt.Printf("key:%v, val:%+v\n", key, value)
	}
}

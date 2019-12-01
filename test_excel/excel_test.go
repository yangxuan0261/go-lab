package test_excel

import (
	"GoLab/test_excel/gen"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"testing"
)

// gen 里面的数据由 xls_deploy_tool.py 工具生成

func Test_excel(t *testing.T) {
	dataArr := &dataconfig.SkinArray{}
	_ = dataArr

	file := "I:/workspace/go/GoWinEnv_new/src/GoLab/test_excel/gen/cfg_skin.data"
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

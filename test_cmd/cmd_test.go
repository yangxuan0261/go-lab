package test_cmd

import (
	"GoLab/common/convert"
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"
)

func Test_001(t *testing.T) {
	//cmd := exec.Command("/bin/bash", "-c")
	cmd := exec.Command("cmd.exe")

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}

	stdin.Write([]byte("ipconfig\n"))
	stdin.Write([]byte("dir\n"))
	stdin.Close()

	outBytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("--- Execute failed when Wait:" + err.Error())
		return
	}

	cmdRe := convert.Byte2String(outBytes, convert.GB18030)
	fmt.Println(cmdRe)

	fmt.Printf("\n\n--- Execute finished:\n%s\n", cmdRe)
}

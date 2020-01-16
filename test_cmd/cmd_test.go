package test_cmd

import (
	"GoLab/common/convert"
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"
)

func execCmd(command string) ([]byte, error) {
	//cmd := exec.Command("/bin/bash") // linux
	cmd := exec.Command("cmd") // windows

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	stdin.Write([]byte(fmt.Sprintf("%s\n", command)))
	stdin.Close()

	outBytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return outBytes, nil
}

func Test_001(t *testing.T) {
	outBytes, err := execCmd("dir")
	if err != nil {
		panic(err)
	}

	outStr := convert.Byte2String(outBytes, convert.GB18030)
	fmt.Printf("\n\n--- Execute finished:\n%s\n", outStr)
}

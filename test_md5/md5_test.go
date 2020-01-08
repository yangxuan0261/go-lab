package test_md5

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"testing"
)

func GetFileMD5(file string) (string, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return "", err
	}
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func GetStringMD51(str string) (string, error) {
	h := md5.New()
	if _, err := io.WriteString(h, str); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%x", h.Sum(nil)), nil
	}
}

func GetStringMD52(str string) string {
	data := []byte(str)
	h := md5.Sum(data)
	return fmt.Sprintf("%x", h)
}

func Test_file(t *testing.T) {
	path := "i:/Git_repo_assets/z_mywiki/config/idea/goland-settings.zip"
	if md5Str, err := GetFileMD5(path); err != nil {
		panic(fmt.Sprintf("--- err:%+v\n", err))
	} else {
		fmt.Println("--- md5Str:", md5Str) // cfa262e465a513025cc0338ebb67343d
	}
}

func Test_str(t *testing.T) {
	str := "b2af71c4de723783"
	if md5Str, err := GetStringMD51(str); err != nil {
		panic(fmt.Sprintf("--- err:%+v\n", err))
	} else {
		fmt.Println("--- md5Str:", md5Str) // cf7be73c856c99c0fe02a78a562375c5
	}

	fmt.Println("--- md5Str 222:", GetStringMD52(str)) // cf7be73c856c99c0fe02a78a562375c5
}

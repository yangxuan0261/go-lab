package test_mail

import (
	"crypto/tls"
	"fmt"
	"go-lab/lib/tool"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

type Sender struct {
	username string
	password string
}

func getToArr(path string) []string {
	arr, err := tool.Readline(path)
	if err != nil {
		panic(err)
	}
	return arr
}

func getFromArr(path string) []*Sender {
	arr, err := tool.Readline(path)
	if err != nil {
		panic(err)
	}

	var senderArr []*Sender
	for _, line := range arr {
		args := strings.Split(line, ",")
		if len(args) != 2 {
			panic(fmt.Sprintf("line error, line: %s", line))
		}
		senderArr = append(senderArr, &Sender{
			username: args[0],
			password: args[1],
		})
	}

	return senderArr
}

func getContent(path string) string {
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bts)
}

func sendAll(mailArr []string, content string, senderArr []*Sender) {
	slen := len(senderArr)
	for index, to := range mailArr {
		sIndex := index % slen
		from := senderArr[sIndex]
		fmt.Printf("--- index:%02d, from: %s, to: %s\n", index, from.username, to)
		sendMail(from.username, to, from.password, content)
		time.Sleep(time.Second * 1) // 不敢发的太频繁, 间隔 1s
	}
}

func sendMail(from string, to string, password, content string) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", from, "RMG Station")
	m.SetAddressHeader("To", to, "")
	m.SetHeader("Subject", "RMG Station")
	m.SetBody("text/html", content) // 内容换行必须要用 <br>

	d := gomail.NewDialer("smtp.mxhichina.com", 465, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Printf("--- Error, from: %s, to:%s, msg: %v\n", from, to, err)
	}
}

func Test_sendAll(t *testing.T) {
	toPath := "E:/its_rummy_info/mail/20200325_144011/temp_support@rmgstation.com_20200325-114223_871.txt"
	fromPath := "E:/its_rummy_info/mail/20200325_144011/temp_sender.txt"
	contentPath := "E:/its_rummy_info/mail/20200325_144011/content.txt"

	toArr := getToArr(toPath)
	senderArr := getFromArr(fromPath)
	content := getContent(contentPath)
	fmt.Println("--- toArr len:", len(toArr))
	fmt.Println("--- senderArr len:", len(senderArr))
	//fmt.Println("--- content :", content)

	sendAll(toArr, content, senderArr)

	fmt.Println("--- all send success")
}

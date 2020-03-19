package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"testing"
	"time"
)

/*
参考:
- https://github.com/geemu/accept-mail/blob/master/test/main/acceptAllMail.go
- https://github.com/emersion/go-imap/wiki/Fetching-messages
*/

type Literal interface {
	io.Reader
	Len() int
}

type ResponseMessage struct {
	From    string
	Subject string
	Content string
}

type Account struct {
	Username string
	Password string
}

func AcceptAllMail(addr, user, pass string) ([]Literal, error) {
	client, err := client.DialTLS(addr, nil)
	if err != nil {
		return nil, err
	}
	defer client.Logout()
	if err := client.Login(user, pass); err != nil {
		log.Fatal(err)
	}
	// 收件箱
	mbox, err := client.Select("INBOX", true)
	if err != nil {
		return nil, err
	}

	if mbox.Messages == 0 {
		return nil, nil
	}

	seqset := new(imap.SeqSet)
	seqset.AddRange(uint32(1), mbox.Messages)

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- client.Fetch(seqset, []imap.FetchItem{"BODY[]"}, messages)
	}()
	// 返回体
	var response []Literal
	// 收件箱的所有邮件
	//		raw:    "BODY[]",
	section := &imap.BodySectionName{BodyPartName: imap.BodyPartName{}}

	for msg := range messages {
		//r := msg.GetBody("BODY[]")
		r := msg.GetBody(section)
		if r == nil {
			return nil, fmt.Errorf("没有邮件内容")
		}
		response = append(response, r)
	}
	return response, nil
}

func main() {
	Test_fetch01(nil)
}

func Test_fetch01(t *testing.T) {
	err := doFetch()
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		debug.PrintStack()
	}
}

func doFetch() error {
	fromMap := make(map[string]bool)

	accPath := "./temp_account.json"
	ignorePath := "./temp_ignore.txt"
	outputPath := "./temp_%s.txt"

	acc, err := readAccount(accPath)
	if err != nil {
		return err
	}

	fmt.Println("Username:", acc.Username)
	fmt.Println("Password:", acc.Password)

	println("正在处理中......")

	request, err := AcceptAllMail("imap.mxhichina.com:993", acc.Username, acc.Password)
	//request, err := AcceptAllMail("imap.mxhichina.com:993", "aaa@bbb.com", "ccc")
	//request, err := AcceptAllMail("imap.qq.com:993", "ccc@qq.com", "asdasd")
	if err != nil {
		log.Fatal(err)
	}

	var response []*ResponseMessage
	for _, r := range request {
		mr, err := mail.CreateReader(r)
		if err != nil {
			log.Fatal(err)
		}

		var existEntity = new(ResponseMessage)
		header := mr.Header
		if from, err := header.AddressList("From"); err == nil {
			for _, value := range from {
				existEntity.From = value.Address
			}
		}

		if subject, err := header.Subject(); err == nil {
			existEntity.Subject = subject
		}

		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			//mt := reflect.TypeOf(p.Header)
			//msgID := mt.Elem().Name()
			//println("msgID:", msgID)

			//mail.InlineHeader
			//mail.AttachmentHeader

			switch p.Header.(type) {
			//case mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			//b, err := ioutil.ReadAll(p.Body)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//existEntity.Content = string(b)
			}
			response = append(response, existEntity)
		}

		// 测试输出
		//fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>> 111 - 111")
		for _, value := range response {
			//fmt.Println("发件人：", value.From)
			//fmt.Println("主题：", value.Subject)
			//fmt.Println("发件内容：", value.Content)

			fromMap[value.From] = true // 去重收集 发件人

		}
		//fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>> 111 - 222")
		//fmt.Printf("--- total 222:%d\n", len(response))
	}

	//path := "C:/Users/wolegequ/Desktop/ali_mail_from.txt"

	ignoreMap, err := readIgnore(ignorePath)
	if err != nil {
		return err
	}

	removeIgnore(fromMap, ignoreMap)

	outputPath = fmt.Sprintf(outputPath, fmt.Sprintf("%s_%s_%d", acc.Username, time.Now().Format("20060102-150405"), len(fromMap)))
	fmt.Printf("output: %s\n", outputPath)
	writeFile(outputPath, fromMap)
	println("Success ^_^")
	time.Sleep(time.Hour * 100)
	return nil
}

func writeFile(path string, fromMap map[string]bool) {
	totalLog := fmt.Sprintf("------- total num:%d\n", len(fromMap))
	fmt.Printf(totalLog)

	//fl.Write([]byte(totalLog))
	var buff []byte
	for from, _ := range fromMap { // 遍历
		buff = append(buff, []byte(from+"\n")...)
		//_, err := fl.Write([]byte(from + "\n"))
		//if err != nil {
		//	panic(err)
		//}
	}

	ioutil.WriteFile(path, buff, os.ModePerm)
}

func readAccount(path string) (*Account, error) {
	if !ExistsFile(path) {
		return nil, fmt.Errorf("--- no account path found")
	}

	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	acc := &Account{}
	json.Unmarshal(bts, acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func readIgnore(path string) (map[string]bool, error) {
	if !ExistsFile(path) {
		return nil, nil
	}

	fi, err := os.Open(path)
	defer fi.Close()
	if err != nil {
		return nil, err
	}

	ignoreMap := make(map[string]bool)
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		line := strings.TrimSpace(string(a))
		ignoreMap[line] = true
	}
	return ignoreMap, nil
}

func removeIgnore(srcMap map[string]bool, ignoreMap map[string]bool) {
	if ignoreMap == nil {
		return
	}

	for acc, _ := range srcMap {
		for line, _ := range ignoreMap {
			if strings.Contains(acc, line) {
				delete(srcMap, acc)
			}
		}
	}
}

func ExistsFile(path string) bool {
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return true
	}
	return false
}

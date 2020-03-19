package test_mail

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"log"
	"os"
	"testing"
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

func Test_fetch01(t *testing.T) {
	fromMap := make(map[string]bool)

	request, err := AcceptAllMail("imap.mxhichina.com:993", "ccc@ddd.com", "eee")
	//request, err := AcceptAllMail("imap.qq.com:993", "ccc@qq.com", "asdasd")
	if err != nil {
		log.Fatal(err)
	}

	// 返回体
	var response []*ResponseMessage
	for _, r := range request {
		// Create a new mail reader
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

			//textproto.Header{}
			switch p.Header.(type) {

			//case mail.TextHeader:
			//	// This is the message's text (can be plain-text or HTML)
			//	b, err := ioutil.ReadAll(p.Body)
			//	if err != nil {
			//		fmt.Println(err)
			//	}
			//	existEntity.Content = string(b)
			//}
			}
			response = append(response, existEntity)
		}

		// 测试输出
		//fmt.Println(">>>>>>>>>> 111 - 111")
		for _, value := range response {
			//fmt.Println("发件人：", value.From)
			//fmt.Println("主题：", value.Subject)
			//fmt.Println("发件内容：", value.Content)
			//fmt.Println(">>>>>>>>>> 111 - 222")

			if _, ok := fromMap[value.From]; !ok { // 去重收集 发件人
				fromMap[value.From] = true
			}
		}
		//fmt.Println("--- total 222:", len(response))
	}

	path := "C:/Users/wolegequ/Desktop/ali_mail_from.txt"
	writeFile(path, fromMap)
	//fmt.Println("--- total 111:", len(response))
}

func writeFile(path string, fromMap map[string]bool) {
	fmt.Printf("--------- total:%d\n", len(fromMap))
	fl, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
	defer fl.Close()

	if err != nil {
		panic(err)
	}

	for from, _ := range fromMap { // 遍历
		_, err := fl.Write([]byte(from + "\n"))
		if err != nil {
			panic(err)
		}
	}
}

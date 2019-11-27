package test_mail

import (
	"fmt"
	"net/smtp"
	"strings"
	"testing"
)

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func Test_qq(t *testing.T) {
	user := "553972977@qq.com"
	password := "asdasdasd" // qq 邮箱的 授权码, 不是 qq 密码
	host := "smtp.qq.com:25"
	to := "364105996@qq.com" // 364105996@qq.com;364105996@qq.com, 用 ; 分割

	subject := "Test send email by golang"

	body := `
    <html>
    <body>
    <h3>
    "这是GO语言写的测试邮件。"
    </h3>
    </body>
    </html>
    `
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
}

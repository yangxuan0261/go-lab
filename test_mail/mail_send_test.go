package test_mail

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"testing"
)

/*
阿里云邮箱
企业云邮箱各个服务器地址及端口信息如下：

收件服务器地址：
POP 服务器地址：pop3.mxhichina.com 端口110，SSL 加密端口995
或
IMAP 服务器地址：imap.mxhichina.com 端口143，SSL 加密端口993

发件服务器地址：
SMTP 服务器地址：smtp.mxhichina.com 端口25， SSL 加密端口465
*/

/*
qq 邮箱
如何设置IMAP服务的SSL加密方式？
使用SSL的通用配置如下：
接收邮件服务器：imap.qq.com，使用SSL，端口号993
发送邮件服务器：smtp.qq.com，使用SSL，端口号465或587
*/
func Test_aliyun(t *testing.T) {
	from := "rmg10@rmgstation.com"
	to := "364105996@qq.com"
	content := "Thank you for being one of our royal members."

	m := gomail.NewMessage()
	m.SetAddressHeader("From", from, "RMG")
	m.SetAddressHeader("To", to, "")
	m.SetHeader("Subject", "Welcome to RMG Station")
	m.SetBody("text/html", content)

	d := gomail.NewDialer("smtp.mxhichina.com", 465, from, "Password")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("--- Test_aliyun send success")
	}
}

func Test_qq02(t *testing.T) {
	from := "553972977@qq.com"
	to := "364105996@qq.com"
	content := "Thank you for being one of our royal members."

	m := gomail.NewMessage()
	m.SetAddressHeader("From", from, "RMG")
	m.SetAddressHeader("To", to, "")
	m.SetHeader("Subject", "Welcome to RMG Station")
	m.SetBody("text/html", content)

	d := gomail.NewDialer("smtp.qq.com", 25, from, "Password")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("--- Test_qq02 send success")
	}
}

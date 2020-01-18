package test_mail

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"testing"
)

func Test_aliyun(t *testing.T) {
	from := "rmg10@rmgstation.com"
	to := "364105996@qq.com"
	content := "Thank you for being one of our royal members."

	m := gomail.NewMessage()
	m.SetAddressHeader("From", from, "RMG")
	m.SetAddressHeader("To", to, "")
	m.SetHeader("Subject", "Welcome to RMG Station")
	m.SetBody("text/html", content)

	d := gomail.NewDialer("smtp.mxhichina.com", 465, from, "password")
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

	d := gomail.NewDialer("smtp.qq.com", 25, from, "password")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("--- Test_qq02 send success")
	}
}

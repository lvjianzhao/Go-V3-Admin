package utils

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"server/global"
)

func SendEmail(subject, msg string, recipients []string) (err error) {
	smtp := global.TD27_CONFIG.SMTP
	emailBody := fmt.Sprintf("<h3>%v</h3>", msg)
	emailFrom := fmt.Sprintf("Orchsym Drawer <%v>", smtp.Username)

	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", recipients...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer(smtp.Server, smtp.Port, smtp.Username, smtp.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 发送邮件
	return d.DialAndSend(m)
}

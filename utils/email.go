package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/smtp"
	"strings"
)

type EmailType struct{}

var Email = EmailType{}

type EmailConfig struct {
	Enable   bool
	Username string
	Address  string
	Host     string
	Port     string
	Password string
	SSL      bool
}

// SendTemplate 发送HTML模板，从配置文件中读取邮箱配置
func SendTemplate(emailConfig *EmailConfig, target, htmlTemplate string, placeholders map[string]string) error {
	for placeholder, value := range placeholders {
		htmlTemplate = strings.ReplaceAll(htmlTemplate, placeholder, value)
	}
	err := SendEmail(emailConfig, target, htmlTemplate, true)
	if err != nil {
		return err
	}
	return nil
}

// SendEmail 发送邮件
func SendEmail(emailConfig *EmailConfig, target, content string, isHTML bool) error {
	if !emailConfig.Enable {
		return nil
	}
	auth := smtp.PlainAuth("", emailConfig.Username, emailConfig.Password, emailConfig.Host)
	var conn net.Conn
	var err error
	if emailConfig.SSL {
		conn, err = tls.Dial("tcp", net.JoinHostPort(emailConfig.Host, emailConfig.Port), &tls.Config{ServerName: emailConfig.Host})
	} else {
		conn, err = net.Dial("tcp", net.JoinHostPort(emailConfig.Host, emailConfig.Port))
	}
	if err != nil {
		return err
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client, err := smtp.NewClient(conn, emailConfig.Host)
	if err != nil {
		return err
	}
	defer func(client *smtp.Client) {
		err := client.Quit()
		if err != nil {

		}
	}(client)

	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(emailConfig.Address); err != nil {
		return err
	}

	if err = client.Rcpt(target); err != nil {
		return err
	}

	writer, err := client.Data()
	if err != nil {
		return err
	}
	defer func(writer io.WriteCloser) {
		err := writer.Close()
		if err != nil {

		}
	}(writer)
	var message string
	if isHTML {
		message = fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Notification\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s", emailConfig.Address, target, content)
	} else {
		message = fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Notification\r\n\r\n%s", emailConfig.Address, target, content)
	}
	_, err = writer.Write([]byte(message))
	return err
}

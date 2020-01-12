package provider

import (
	"strconv"
	"html/template"

	"github.com/yinxulai/goutils/config"
	"gopkg.in/gomail.v2"
)

// NewEmail NewEmail
func NewEmail() *Email {
	email := new(Email)
	email.host = config.MustGet("email_host")
	email.account = config.MustGet("email_account")
	email.templates = map[string]template.Template{}
	email.password = config.MustGet("email_password")
	email.port, _ = strconv.Atoi(config.MustGet("email_port"))
	email.templates["VerifyCode"], _ = template.ParseFiles(config.MustGet("email_verify_code_template"))
	return email
}

// Email 服务
type Email struct {
	port     int
	host     string
	account  string
	password string
	templates map[string]template.Template
}

// Weights 权重
func (srv Email) Weights() int {
	return 1
}

// SendVerifyCode 发送验证哪
func (srv Email) SendVerifyCode(email, action, code string) error {
	body := srv.templates["VerifyCode"].  action + code
	msg := gomail.NewMessage()
	msg.SetHeader("To", email)      //发送给多个用户
	msg.SetBody("text/html", body)  //设置邮件正文
	msg.SetHeader("Subject", "验证码") //设置邮件主题
	dial := gomail.NewDialer(srv.host, srv.port, srv.account, srv.password)
	return dial.DialAndSend(msg)
}

package service

import (
	"net/smtp"
	"bytes"
)
var(

)

type MailService struct {
	smtpServer string
	smtpPort string
	from string
	plainAuth smtp.Auth

}

func NewMailService (smtpServer, smtpPort, from, pass string) *MailService{


	return &MailService{smtpServer, smtpPort, from, plainAuth}
}

func (self *MailService) SendMail (to string, subject string, body []byte) error{



	return smtp.SendMail(self.smtpServer + self.smtpPort,
		self.plainAuth,  self.from, []string{to}, msg.Bytes())
}
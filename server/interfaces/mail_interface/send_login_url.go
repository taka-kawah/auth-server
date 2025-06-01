package mailinterface

import (
	"fmt"
	"net/smtp"
	"os"
)

type MailSender interface {
	SendPasswordFromEmail(recipientAddress string, id int64) error
}

type mailSenderImpl struct {
	senderAddress string
	auth          smtp.Auth
}

func NewMailSender() *mailSenderImpl {
	var (
		senderAddress = os.Getenv("senderAddress")
		password      = os.Getenv("password")
	)
	return &mailSenderImpl{
		senderAddress: senderAddress,
		auth:          smtp.CRAMMD5Auth(senderAddress, password),
	}
}

func (m *mailSenderImpl) SendPasswordFromEmail(recipientAddress string, id int64) error {
	var (
		recipient = []string{recipientAddress}
		hostname  = os.Getenv("hostname")
		port      = os.Getenv("port")
		subject   = "パスワードを入力して下さい"
	)
	body := fmt.Sprintf("アカウント登録のお申し込みありがとうございます。\n以下のウェブページからパスワードを入力し、アカウント登録を完了させてください。\nURLの有効時間は10分です。\n\n%s", m.provideUrl(id))
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", recipientAddress, subject, body))
	if err := smtp.SendMail(fmt.Sprintf("%s:%s", hostname, port), m.auth, m.senderAddress, recipient, msg); err != nil {
		return err
	}
	return nil
}

func (m *mailSenderImpl) provideUrl(id int64) string {
	return fmt.Sprintf("%s/signup?id=%v", os.Getenv("domain"), id)
}

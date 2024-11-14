package infrastructure

import (
	"fmt"
	"net/smtp"
)

type EmailService interface {
	SendVerificationCode(email, code string) error
}

type emailService struct {
	hostEmail    string
	hostPassword string
	hostAddress  string
	post         string
}

func NewEmailService(hostEmail, hostPassword, hostAddress, port string) EmailService {
	return &emailService{
		hostEmail:    hostEmail,
		hostPassword: hostPassword,
		hostAddress:  hostAddress,
		post:         port,
	}
}

func (e *emailService) SendVerificationCode(email, code string) error {

	auth := smtp.PlainAuth("", e.hostEmail, e.hostPassword, e.hostAddress)
	err := smtp.SendMail(e.hostAddress+":"+e.post, auth, e.hostEmail, []string{email}, []byte("Subject: Verification Code\n\n"+code))

	if err != nil {
		fmt.Println("err", err)
		return fmt.Errorf("failed to send verification code: %v", err)
	}

	return nil
}

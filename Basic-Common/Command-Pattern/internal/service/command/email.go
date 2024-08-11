package command

import (
	"fmt"
	"net/smtp"
)

const CommandTypeEmail CommandType = "email"

// TODO: Load values from ENV
const (
	smtpUserName = ""
	smtpPassword = ""
	smtpHost     = ""
	smtpPort     = ""
)

type emailCommand struct {
	addr  string
	email EmailInfo
}

type EmailInfo struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func UnmarshalEmailCommand(body string) (Command, error) {
	/*
		TODO: Unmarshal JSON string into EmailInfo object
	*/
	return NewEmailCommand(EmailInfo{From: "", To: []string{""}, Subject: "", Body: ""}), nil
}

func NewEmailCommand(email EmailInfo) Command {
	return &emailCommand{
		email: email,
	}
}

func (e *emailCommand) Serialize() (string, error) {
	/*
		TODO: Marshal EmailInfo into CommandPayload
	*/
	return "", nil
}

func (e *emailCommand) Execute() error {
	smtpAuth := smtp.PlainAuth("", smtpUserName, smtpPassword, smtpHost)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := fmt.Sprintf("Subject: %s\n", e.email.Subject)
	msg := []byte(subject + mime + "\n" + e.email.Body)

	return smtp.SendMail(e.addr, smtpAuth, e.email.From, e.email.To, msg)
}

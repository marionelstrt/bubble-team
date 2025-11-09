package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/marionelstrt/bubble-team/models"
)

// https://stackoverflow.com/questions/11065913/send-email-through-unencrypted-connection
type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func SendVerificationEmail(user models.User) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	tmpl, err := template.ParseFiles("templates/account_verification.html")
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var body bytes.Buffer

	to := []string{user.Email}
	subject := "Subject: Account Verification\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body.WriteString(subject)
	body.WriteString(mime)

	err = tmpl.Execute(&body, user)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	auth := unencryptedAuth{smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, to, body.Bytes())
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

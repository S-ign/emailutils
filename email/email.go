package email

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/s-ign/emailutils/config"
	"gopkg.in/gomail.v2"
)

// Message email message information
type Message struct {
	subject      string
	templatePath string
	message      string
	to           []string
}

// NewMessage Creates a message object, if templatePath is supplied it will
// apply message to template. Template must contain {{.Body}}, message
// variable will replaces this placeholder
func NewMessage(subject, message, templatePath string, to []string) (*Message, error) {

	// embeds message into template, updating message string to its template html
	// format
	if templatePath != "" {
		t, _ := template.ParseFiles(templatePath)
		var body bytes.Buffer

		t.Execute(&body, struct {
			Body string
		}{
			Body: message,
		})
		message = body.String()
	}

	return &Message{
		subject: subject,
		message: message,
		to:      to,
	}, nil
}

// SendEmail wraper method around gomail.NewMessage, makes applying html
// templates to messages simpler
func SendEmail(auth *config.Auth, smtp *config.SMTPConfig, message *Message) error {
	gm := gomail.NewMessage()
	gm.SetHeader("From", auth.From)
	gm.SetHeader("To", message.to...)
	gm.SetHeader("Subject", message.subject)
	gm.SetBody("text/html", message.message)

	d := gomail.NewDialer(smtp.SMTPHost, smtp.SMTPPort, auth.From, auth.Password)

	if err := d.DialAndSend(gm); err != nil {
		return fmt.Errorf("SendEmail: unable to send email: %v", err)
	}

	return nil
}

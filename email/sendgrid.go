package email

import (
	"fmt"
	"log"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	From    string
	To      string
	Content string
	Subject string
	Api     string
}

func SendEmail(e Email) {
	from := mail.NewEmail(strings.Split(e.From, "@")[0], e.From)
	subject := e.Subject
	to := mail.NewEmail(strings.Split(e.To, "@")[0], e.To)
	plainTextContent := e.Content
	htmlContent := "<strong>" + plainTextContent + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(e.Api)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
	}
}

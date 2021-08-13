package main
import (
	"bytes"
	"net/smtp"
	"strconv"
	"text/template"   // Uses text templates to send plain text email
)

//data structure for an email
type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

type EmailCredentials struct {
	Username, Password, Server string
	Port                       int
}
//The email template as s string
const emailTemplate = `From:{{.From}}
To:{{.To}}
Subject {{.Subject}}
{{.Body}}
`

var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(emailTemplate)
}
func main() {
// populates a dataset with the  email for the template and mail client
	message := &EmailMessage{
		From:    "me@example.com",
		To:      []string{"you@example.com"},
		Subject: "A test",
		Body:    "Just saying hi",
	}
	// populate a Buffer with the rendered message text from the template
	var body bytes.Buffer
	t.Execute(&body, message)

	//Sets up the SMTP mail Client
	authCreds := &EmailCredentials{
		Username: "myUsername",
		Password: "myPass",
		Server:   "smt.example.com",
		Port:     8080,
	}
	auth := smtp.PlainAuth("",
		authCreds.Username,
		authCreds.Password,
		authCreds.Server,
	)

	// sends the email
	smtp.SendMail(authCreds.Server + ":"+strconv.Itoa(authCreds.Port),
	auth,
	message.From,
	message.To,
	body.Bytes())
}






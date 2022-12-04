package utils

import (
	"log"
	"os"

	mail "github.com/xhit/go-simple-mail/v2"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, World</title>
</head>
<body>
   <p>
	 	The ISS is right now over you.
	 </p>
	 <p>
		You are receiving this message because you subscribed to get notification for ISS location.
	</p>
</body>
`

func SendEmail(to string) {

	server := mail.NewSMTPClient()
	server.Host = os.Getenv("MAIL_HOST")
	server.Port = 587
	server.Username = os.Getenv("MAIL_DOMAIN_ADDRESS")
	server.Password = os.Getenv("MAIL_PASSWORD")
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewMSG()
	email.SetFrom(os.Getenv("MAIL_DOMAIN_ADDRESS"))
	email.AddTo(to)
	email.SetSubject("ISS is on your timezone")

	email.SetBody(mail.TextHTML, htmlBody)

	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
}

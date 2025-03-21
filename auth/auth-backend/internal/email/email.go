package email

import (
	"crypto/rand"
	"errors"
	"os"
	"strconv"

	"github.com/wneessen/go-mail"
)

const charSet = "ABCDEFGHJKLMNOPQRSTUVWXYZ123456789"

var ErrCreateMail = errors.New("Unable to create email body")

func GenerateVerificationCode(length int) (string, error) {
	code := make([]byte, length)
	_, err := rand.Read(code)
	if err != nil {
		return "", err
	}
	for i := range code {
		code[i] = charSet[code[i]%byte(len(charSet))]
	}
	return string(code), nil
}

func SendVerficationCode(recipient string, code string) error {
	/* Create message */
	message := mail.NewMsg()
	if err := message.From(os.Getenv("SMTP_FROM")); err != nil {
		return errors.Join(ErrCreateMail, err)
	}
	if err := message.To(recipient); err != nil {
		return errors.Join(ErrCreateMail, err)
	}
	message.Subject("DeGuzman Cloud Verification Code")
	message.SetBodyString(mail.TypeTextPlain, "Your verification code is "+code)

	/* Send email */
	host, port, username, password := os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"), os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	client, err := mail.NewClient(
		host,
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithPort(portInt),
		mail.WithUsername(username),
		mail.WithPassword(password),
		mail.WithTLSPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return err
	}
	if err := client.DialAndSend(message); err != nil {
		return err
	}
	return nil
}

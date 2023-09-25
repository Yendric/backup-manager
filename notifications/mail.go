package notifications

import (
	"errors"
	"io"

	"github.com/go-mail/mail"
	"github.com/yendric/backup-manager/config"
)

func sendMail(subject string, content string, logs string) error {
	mailEnabled := config.Configuration.Notification.Email.Enabled
	if !mailEnabled {
		return errors.New("mail is not enabled")
	}

	message := mail.NewMessage()
	message.SetAddressHeader("From", config.Configuration.Notification.Email.FromAddress, "Backup Manager")
	message.SetHeader("To", config.Configuration.Notification.Email.To...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", content)
	message.Attach(
		"log.txt",
		mail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write([]byte(logs))
			return err
		}),
	)

	dialer := mail.NewDialer(
		config.Configuration.Notification.Email.SmtpHost,
		config.Configuration.Notification.Email.SmtpPort,
		config.Configuration.Notification.Email.SmtpUsername,
		config.Configuration.Notification.Email.SmtpPassword,
	)

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}

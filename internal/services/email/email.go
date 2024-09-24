package email

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/Magic-Kot/store-notification-service/internal/config"

	"github.com/rs/zerolog"
)

type Email struct {
	username string
	password string
}

func NewEmailService(cfg *config.MailService) *Email {
	return &Email{username: cfg.Username, password: cfg.Password}
}

// SendEmail -sending an e-mail message to the user
func (e *Email) SendEmail(ctx context.Context, email string, subject string, body string) error {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("starting the 'SendEmail' service")

	auth := smtp.PlainAuth("", e.username, e.password, "smtp.gmail.com")

	msg := fmt.Sprintf("Subject: %s\n%s", subject, body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, e.username, []string{email}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

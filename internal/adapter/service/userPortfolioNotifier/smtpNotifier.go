package userportfolionotifier

import "github.com/stasdashkevitch/crypto_info/internal/config"

type SMTPNotifier struct {
	SMTPHost  string
	SMTPPort  string
	Username  string
	Password  string
	Recipient string
}

func NewSMTPNotifier(cfg *config.Config, recipient string) *SMTPNotifier {
	return &SMTPNotifier{
		SMTPHost:  cfg.SMTP.SMTPHost,
		SMTPPort:  cfg.SMTP.SMTPPort,
		Username:  cfg.SMTP.Username,
		Password:  cfg.SMTP.Password,
		Recipient: recipient,
	}
}

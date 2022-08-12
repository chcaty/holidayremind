package smtp

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"
)

var pool *email.Pool

func SendEmail(message EmailMessage, config Config) error {
	var err error
	if pool == nil {
		pool, err = email.NewPool(config.Host+":"+config.Port, config.MaxClient, smtp.PlainAuth("", config.UserName, config.Password, config.Host))
		if err != nil {
			return fmt.Errorf("create new email pool fail. error: %w", err)
		}
	}
	e := &email.Email{
		From:    config.UserName,
		To:      message.Receiver,
		Subject: message.Subject,
		HTML:    []byte(message.Html),
	}
	if len(message.Attachment) > 0 {
		for i := 0; i < len(message.Attachment); i++ {
			if len(message.Attachment[i]) > 0 {
				_, err = e.AttachFile(message.Attachment[i])
				if err != nil {
					return fmt.Errorf("email attach file fail. error: %w", err)
				}
				time.Sleep(time.Second * 10)
			}
		}
	}

	err = pool.Send(e, 30*time.Second)
	if err != nil {
		return fmt.Errorf("send email fail. error: %w", err)
	}
	return nil
}

package smtp

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"time"
)

var pool *email.Pool

func SendEmail(subject, body string, attachment []string) error {
	var err error
	if pool == nil {
		pool, err = email.NewPool(Smtp.Host+":"+Smtp.Port, Smtp.MaxClient, smtp.PlainAuth("", Smtp.UserName, Smtp.Password, Smtp.Host))
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	e := &email.Email{
		From:    Smtp.UserName,
		To:      Receiver,
		Subject: subject,
		HTML:    []byte(body),
	}
	if len(attachment) > 0 {
		for i := 0; i < len(attachment); i++ {
			if len(attachment[i]) > 0 {
				_, err = e.AttachFile(attachment[i])
				if err != nil {
					return err
				}
				time.Sleep(time.Second * 1)
			}
		}
	}

	err = pool.Send(e, 10*time.Second)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

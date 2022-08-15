package configs

import (
	"holidayRemind/internal/holidayremind/smtp"
)

const DingTalkToken = "afc3c084e0a0a7936196b6a686f9bd382dcb5859609ee58b7c234ff6d94ad929"

var (
	SmtpConfig = smtp.Config{
		Host:      "smtp.163.com",
		Port:      "465",
		UserName:  "chcaty@163.com",
		Password:  "LVAULLJARBXIKAAC",
		MaxClient: 5,
	}

	Receiver = []string{
		//"chenzuo@hotmail.com",
		"1120873075@qq.com",
	}
)

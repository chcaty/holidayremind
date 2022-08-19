package smtp

type Config struct {
	Host      string
	Port      string
	UserName  string
	Password  string
	MaxClient int
}

type SimpleEmail struct {
	Subject    string
	Html       string
	Attachment []string
	Receiver   []string
}

func SetSimpleEmail(email *SimpleEmail, subject string, html string, attachment []string, receiver []string) {
	email.Html = html
	email.Subject = subject
	email.Attachment = attachment
	email.Receiver = receiver
}

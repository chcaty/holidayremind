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

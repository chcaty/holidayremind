package template

type Type int

const (
	MarkDown Type = iota
	Email
)

var templateFileName = "messagetemplate"

package net

type ContentType string

const (
	Json ContentType = "application/json"
	Xml  ContentType = "application/xml"
	Form ContentType = "multipart/form-data"
)

var DefaultHeader = map[string]string{
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
}

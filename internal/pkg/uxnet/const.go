package uxnet

type contentType string

const (
	Json contentType = "application/json"
	Xml  contentType = "application/xml"
	Form contentType = "multipart/form-data"
)

var DefaultHeader = map[string]string{
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
}

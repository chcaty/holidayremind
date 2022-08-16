package net

type ContentType string

const (
	Json ContentType = "application/json"
	Xml  ContentType = "application/xml"
	Form ContentType = "multipart/form-data"
)

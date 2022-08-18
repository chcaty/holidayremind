package template

import (
	"holidayRemind/internal/holidayremind/template"
	"testing"
)

func TestGetTemplate(t *testing.T) {
	templateValue := ""
	err := template.GetTemplate("RssBody", template.MarkDown, &templateValue)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(templateValue)
}

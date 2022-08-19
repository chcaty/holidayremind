package template

import (
	"errors"
	"holidayRemind/internal/pkg"
)

func GetTemplate(template *string, name string, templateType Type) error {
	config := Config{}
	err := pkg.GetConfigByJson(&config, "messagetemplate")
	if err != nil {
		return err
	}
	err2 := getTemplateByType(template, name, templateType, config)
	if err2 != nil {
		return err2
	}
	return nil
}

func GetTemplateList(templates *map[string]string, names []string, templateType Type) error {
	config := Config{}
	err := pkg.GetConfigByJson(&config, "messagetemplate")
	if err != nil {
		return err
	}
	template := ""
	for _, name := range names {
		err := getTemplateByType(&template, name, templateType, config)
		if err != nil {
			return err
		}
		(*templates)[name] = template
	}
	return nil
}

func getTemplateByType(template *string, name string, templateType Type, config Config) error {
	if templateType == MarkDown {
		err := getMapTemplate(template, name, config.MarkDownTemplate)
		if err != nil {
			return err
		}
	} else {
		err := getMapTemplate(template, name, config.EmailTemplate)
		if err != nil {
			return err
		}
	}
	return nil
}

func getMapTemplate(template *string, name string, config map[string]string) error {
	if value, ok := config[name]; ok {
		*template = value
		return nil
	}
	return errors.New("calendar not contain date")
}

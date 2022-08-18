package template

import (
	"errors"
	"holidayRemind/internal/pkg"
)

func GetTemplate(name string, templateType Type, template *string) error {
	config := Config{}
	err := pkg.GetConfigByJson("messagetemplate", &config)
	if err != nil {
		return err
	}
	err2 := getTemplateByType(name, templateType, template, config)
	if err2 != nil {
		return err2
	}
	return nil
}

func GetTemplateList(names []string, templateType Type, templates *map[string]string) error {
	config := Config{}
	err := pkg.GetConfigByJson("messagetemplate", &config)
	if err != nil {
		return err
	}
	template := ""
	for _, s := range names {
		err := getTemplateByType(s, templateType, &template, config)
		if err != nil {
			return err
		}
		(*templates)[s] = template
	}
	return nil
}

func getTemplateByType(name string, templateType Type, template *string, config Config) error {
	if templateType == MarkDown {
		err := getMapTemplate(name, config.MarkDownTemplate, template)
		if err != nil {
			return err
		}
	} else {
		err := getMapTemplate(name, config.EmailTemplate, template)
		if err != nil {
			return err
		}
	}
	return nil
}

func getMapTemplate(name string, config map[string]string, template *string) error {
	if value, ok := config[name]; ok {
		*template = value
		return nil
	}
	return errors.New("calendar not contain date")
}

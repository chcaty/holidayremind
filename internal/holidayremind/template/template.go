package template

import (
	"holidayRemind/internal/pkg"
)

func GetTemplate(template *string, name string, templateType Type) error {
	var err error
	config := Config{}
	err = pkg.GetConfigByJson(&config, templateFileName)
	if err != nil {
		return err
	}
	err = getTemplateByType(template, name, templateType, config)
	if err != nil {
		return err
	}
	return nil
}

func GetTemplateList(templates *map[string]string, names []string, templateType Type) error {
	var err error
	config := Config{}
	err = pkg.GetConfigByJson(&config, templateFileName)
	if err != nil {
		return err
	}
	template := ""
	for _, name := range names {
		err = getTemplateByType(&template, name, templateType, config)
		if err != nil {
			return err
		}
		(*templates)[name] = template
	}
	return nil
}

func getTemplateByType(template *string, name string, templateType Type, config Config) error {
	if templateType == MarkDown {
		err := pkg.GetMapValue(template, name, config.MarkDownTemplate)
		if err != nil {
			return err
		}
	} else {
		err := pkg.GetMapValue(template, name, config.EmailTemplate)
		if err != nil {
			return err
		}
	}
	return nil
}

package template

import (
	"holidayRemind/internal/pkg"
)

func GetTemplate(template *string, name string, templateType Type) error {
	var err error
	config := Config{}
	err = pkg.GetConfigByJson(templateFileName, pkg.Json, pkg.Path, &config)
	if err != nil {
		return err
	}
	err = getTemplateByType(name, templateType, config, template)
	if err != nil {
		return err
	}
	return nil
}

func GetTemplateList(templates *map[string]string, names []string, templateType Type) error {
	var err error
	config := Config{}
	err = pkg.GetConfigByJson(templateFileName, pkg.Json, pkg.Path, &config)
	if err != nil {
		return err
	}
	template := ""
	for _, name := range names {
		err = getTemplateByType(name, templateType, config, &template)
		if err != nil {
			return err
		}
		(*templates)[name] = template
	}
	return nil
}

func getTemplateByType(name string, templateType Type, config Config, template *string) error {
	var err error
	switch templateType {
	case MarkDown:
		err = pkg.GetMapValue(config.MarkDownTemplate, name, template)
	case Email:
		err = pkg.GetMapValue(config.EmailTemplate, name, template)
	}
	if err != nil {
		return err
	}
	return nil
}

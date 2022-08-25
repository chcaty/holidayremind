package template

import (
	"holidayRemind/internal/pkg/uxconfig"
	"holidayRemind/internal/pkg/uxmap"
)

func GetTemplate(template *string, name string, templateType Type) error {
	var err error
	config := Config{}
	err = uxconfig.GetValue(templateFileName, uxconfig.Json, uxconfig.Path, &config)
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
	err = uxconfig.GetValue(templateFileName, uxconfig.Json, uxconfig.Path, &config)
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
		err = uxmap.GetMapValue(config.MarkDownTemplate, name, template)
	case Email:
		err = uxmap.GetMapValue(config.EmailTemplate, name, template)
	}
	if err != nil {
		return err
	}
	return nil
}
